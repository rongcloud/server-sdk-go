package rongcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"syscall"
)

// 需要切换域名的网络错误
func isNetError(err error) bool {
	netErr, ok := err.(net.Error)
	if !ok {
		return false
	}
	// 超时
	if netErr.Timeout() {
		return true
	}

	var opErr *net.OpError
	opErr, ok = netErr.(*net.OpError)
	if !ok {
		//  url 错误
		urlErr, ok := netErr.(*url.Error)
		if !ok {
			return false
		}
		opErr, ok = urlErr.Err.(*net.OpError)
		if !ok {
			return false
		}
	}

	switch t := opErr.Err.(type) {
	case *net.DNSError:
		return true
	case *os.SyscallError:
		if errno, ok := t.Err.(syscall.Errno); ok {
			switch errno {
			case syscall.ECONNREFUSED:
				return true
			case syscall.ETIMEDOUT:
				return true
			}
		}
	}

	return false
}

// do http request
func (rc *RongCloud) do(b *http.Request, data interface{}) (*http.Response, error) {
	resp, err := rc.HttpClient.Do(b)
	if err != nil {
		if isNetError(err) {
			rc.ChangeURI()
		}
		return nil, err
	}
	if resp.Body == nil {
		return resp, nil
	}
	defer resp.Body.Close()

	rc.changeURIIfNeed(resp)
	body, err := io.ReadAll(resp.Body)
	//if err = checkHTTPResponseCode(resp, data); err != nil {
	//	return err
	//}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type CodeGetter interface {
	GetCode() int
	GetErrMsg() string
}

// v2 api error
//func checkHTTPResponseCode(resp *http.Response, codeGetter CodeGetter) error {
//	code := codeGetter.GetCode()
//	msg := codeGetter.GetErrMsg()
//	if code != 10000 && code != 200 {
//		return fmt.Errorf(msg)
//	}
//	return nil
//}

// 判断 http status code, 如果大于 500 就切换一次域名
func (rc *RongCloud) changeURIIfNeed(resp *http.Response) {
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		rc.ChangeURI()
	}

	return
}

// postJson
// @param ctx context with
// @param path url path, e.g. /a/b/c
// @param postBody any json able struct
// @param res response struct
func (rc *RongCloud) postJson(ctx context.Context, path string, postBody interface{}, res interface{}) (*http.Response, error) {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(postBody)
	if err != nil {
		return nil, err
	}
	return rc.doRequest(ctx, path, body, &res)
}

func (rc *RongCloud) doRequest(ctx context.Context, path string, body io.Reader, res interface{}) (*http.Response, error) {
	// TODO http header params from ctx
	requestUrl := fmt.Sprintf("%s%s", rc.rongCloudURI, path)
	var req *http.Request
	var err error
	if ctx != nil {
		req, err = http.NewRequest(http.MethodPost, requestUrl, body)
	} else {
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, body)
	}
	if err != nil {
		// TODO new request error more detail
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rc.fillHeader(req)
	resp, err := rc.do(req, &res)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// postFormUrlencoded
// @param ctx context with
// @param path url path, e.g. /a/b/c
// @param formParams x-www-form-urlencoded body
func (rc *RongCloud) postFormUrlencoded(ctx context.Context, path string, formParams url.Values, res interface{}) (*http.Response, error) {
	body := &bytes.Buffer{}
	body.WriteString(formParams.Encode())
	return rc.doRequest(ctx, path, body, &res)
}

type HttpResponseGetter interface {
	GetHttpResponse() *http.Response
}

type RawHttpResponseGetter struct {
	rawHttpResponseInternal *http.Response
}

func (r *RawHttpResponseGetter) GetHttpResponse() *http.Response {
	return r.rawHttpResponseInternal
}
