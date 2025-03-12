/*
 * @Descripttion:
 * @version:
 * @Author: ran.ding
 * @Date: 2019-09-02 18:29:55
 * @LastEditors: ran.ding
 * @LastEditTime: 2019-09-10 15:39:14
 */

// The MIT License (MIT)

// Copyright (c) 2014 RongCloud Rong Cloud

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

/*
 * RongCloud Server API Go Client
 * Created by RongCloud
 * Creation date: 2018-11-28
 * Version: v3
 */

package sdk

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/google/uuid"
)

const (
	// RONGCLOUDSMSURI Default SMS API URL for RongCloud
	RONGCLOUDSMSURI = "http://api.sms.ronghub.com"
	// RONGCLOUDURI Default API URL for RongCloud
	RONGCLOUDURI = "http://api.rong-api.com"
	// RONGCLOUDURI2 Backup API URL for RongCloud
	RONGCLOUDURI2 = "http://api-b.rong-api.com"
	// ReqType Body type
	ReqType = "json"
	// USERAGENT SDK name
	USERAGENT = "rc-go-sdk/3.2.23"
	// DEFAULTTIMEOUT Default timeout, 10 seconds
	DEFAULTTIMEOUT = 10
	// DEFAULT_KEEPALIVE Default HTTP keepalive time, 30 seconds
	DEFAULT_KEEPALIVE = 30
	// DEFAULT_MAXIDLECONNSPERHOST Default maximum idle connections per host, 100
	DEFAULT_MAXIDLECONNSPERHOST = 100
	// DEFAULT_CHANGE_URI_DURATION Interval for automatic API URL switching, in seconds
	DEFAULT_CHANGE_URI_DURATION = 30
)

var (
	defaultExtra = rongCloudExtra{
		rongCloudURI:        RONGCLOUDURI,
		rongCloudSMSURI:     RONGCLOUDSMSURI,
		timeout:             DEFAULTTIMEOUT,
		keepAlive:           DEFAULT_KEEPALIVE,
		maxIdleConnsPerHost: DEFAULT_MAXIDLECONNSPERHOST,
		count:               0,
		changeUriDuration:   DEFAULT_CHANGE_URI_DURATION,
		lastChageUriTime:    0,
	}
	rc   *RongCloud
	once sync.Once
)

// RongCloud appKey appSecret extra
type RongCloud struct {
	appKey    string
	appSecret string
	*rongCloudExtra
	uriLock         sync.Mutex
	globalTransport http.RoundTripper
}

// rongCloudExtra extends RongCloud with custom RongCloud server address and request timeout
type rongCloudExtra struct {
	rongCloudURI        string
	rongCloudSMSURI     string
	timeout             time.Duration
	keepAlive           time.Duration
	maxIdleConnsPerHost int
	count               uint
	changeUriDuration   int64
	lastChageUriTime    int64
}

// getSignature generates a local signature
// Signature calculation method: Concatenate the App Secret, Nonce (random number),
// and Timestamp (Unix timestamp) in order, then compute the SHA1 hash. If the signature verification fails, the API call will return HTTP status code 401.
func (rc RongCloud) getSignature() (nonce, timestamp, signature string) {
	nonceInt := rand.Int()
	nonce = strconv.Itoa(nonceInt)
	timeInt64 := time.Now().Unix()
	timestamp = strconv.FormatInt(timeInt64, 10)
	h := sha1.New()
	_, _ = io.WriteString(h, rc.appSecret+nonce+timestamp)
	signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

// fillHeader adds API signature to the Http Header
func (rc RongCloud) fillHeader(req *httplib.BeegoHTTPRequest) {
	nonce, timestamp, signature := rc.getSignature()
	req.Header("App-Key", rc.appKey)
	req.Header("Nonce", nonce)
	req.Header("Timestamp", timestamp)
	req.Header("Signature", signature)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Header("User-Agent", USERAGENT)
}

// v2 sdk header
func (rc RongCloud) fillHeaderV2(req *httplib.BeegoHTTPRequest) string {
	requestId := uuid.New().String()
	nonce, timestamp, signature := rc.getSignature()
	req.Header("RC-App-Key", rc.appKey)
	req.Header("RC-Timestamp", timestamp)
	req.Header("RC-Nonce", nonce)
	req.Header("RC-Signature", signature)
	req.Header("Content-Type", "application/json")
	req.Header("User-Agent", USERAGENT)
	req.Header("RC-Request-Id", requestId)
	return requestId
}

// fillJSONHeader sets the Http Header Content-Type to JSON format
func fillJSONHeader(req *httplib.BeegoHTTPRequest) {
	req.Header("Content-Type", "application/json")
}

// NewRongCloud creates a RongCloud object
func NewRongCloud(appKey, appSecret string, options ...rongCloudOption) *RongCloud {
	once.Do(func() {
		// Default extended configuration
		defaultRongCloud := defaultExtra
		defaultRongCloud.lastChageUriTime = 0
		rc = &RongCloud{
			appKey:         appKey,
			appSecret:      appSecret,
			rongCloudExtra: &defaultRongCloud,
		}

		for _, option := range options {
			option(rc)
		}

		if rc.globalTransport == nil {
			rc.globalTransport = &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   rc.timeout * time.Second,
					KeepAlive: rc.keepAlive * time.Second,
				}).DialContext,
				MaxIdleConnsPerHost: rc.maxIdleConnsPerHost,
			}
		}
	})

	return rc
}

// GetRongCloud retrieves the RongCloud object
func GetRongCloud() *RongCloud {
	return rc
}

// Customizes HTTP parameters
func (rc *RongCloud) SetHttpTransport(httpTransport http.RoundTripper) {
	rc.globalTransport = httpTransport
}

func (rc *RongCloud) GetHttpTransport() http.RoundTripper {
	return rc.globalTransport
}

// changeURI automatically switches the API server address
// It toggles between api and api2. Cannot switch to other domains. Use PrivateURI for other domain settings.
func (rc *RongCloud) ChangeURI() {
	nowUnix := time.Now().Unix()
	// Check the time interval since the last URI change
	rc.uriLock.Lock()
	if (nowUnix - rc.lastChageUriTime) >= rc.changeUriDuration {
		switch rc.rongCloudURI {
		case RONGCLOUDURI:
			rc.rongCloudURI = RONGCLOUDURI2
		case RONGCLOUDURI2:
			rc.rongCloudURI = RONGCLOUDURI
		default:
		}
		rc.lastChageUriTime = nowUnix
	}
	rc.uriLock.Unlock()
}

// PrivateURI sets the API address for private cloud
func (rc *RongCloud) PrivateURI(uri, sms string) {
	rc.rongCloudURI = uri
	rc.rongCloudSMSURI = sms
}

// urlError checks if the error is a url.Error
func (rc *RongCloud) urlError(err error) {
	// This method is deprecated
}

/*
*
Check the HTTP status code, and switch the domain once if it's greater than or equal to 500
*/
func (rc *RongCloud) checkStatusCode(resp *http.Response) {
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		rc.ChangeURI()
	}

	return
}
