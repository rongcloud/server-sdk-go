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
	resp, err := rc.httpClient.Do(b)
	if err != nil {
		if isNetError(err) {
			rc.changeURI()
		}
		return resp, err
	}
	if resp.Body == nil {
		return resp, nil
	}
	rc.changeURIIfNeed(resp)
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return resp, err
	}
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewBuffer(body))
	codeRes := &CodeResult{}
	err = json.Unmarshal(body, &codeRes)
	if err != nil {
		// skip code result check failed
		return resp, nil
	}
	if codeRes.Code != 200 && codeRes.Code != 10000 {
		return resp, RCErrorNew(codeRes.Code, codeRes.ErrorMessage)
	}
	return resp, nil
}

// CodeResult 融云返回状态码和错误码
type CodeResult struct {
	Code         int    `json:"code"`                   // 返回码，200 为正常。
	ErrorMessage string `json:"errorMessage,omitempty"` // 错误信息
}

// RCErrorNew 创建新的err信息
func RCErrorNew(code int, text string) error {
	return CodeResult{code, text}
}

// Error 获取错误信息
func (e CodeResult) Error() string {
	return fmt.Sprintf("code:%d, errorMessage:%s", e.Code, e.ErrorMessage)
}

// ErrorCode 获取错误码
func (e CodeResult) ErrorCode() int {
	return e.Code
}

// 判断 http status code, 如果大于 500 就切换一次域名
func (rc *RongCloud) changeURIIfNeed(resp *http.Response) {
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		rc.changeURI()
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
	return rc.doRequest(ctx, path, body, &res, "application/json")
}

func (rc *RongCloud) doRequest(ctx context.Context, path string, body io.Reader, res interface{}, contentType string) (*http.Response, error) {
	requestUrl := fmt.Sprintf("%s%s", rc.rongCloudURI, path)
	var req *http.Request
	var err error
	if ctx == nil {
		req, err = http.NewRequest(http.MethodPost, requestUrl, body)
	} else {
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, body)
	}
	if err != nil {
		return nil, fmt.Errorf("new http request error %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	rc.fillHeader(ctx, req)
	resp, err := rc.do(req, &res)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// postFormUrlencoded
func (rc *RongCloud) postFormUrlencoded(ctx context.Context, path string, formParams url.Values, res interface{}) (*http.Response, error) {
	body := &bytes.Buffer{}
	body.WriteString(formParams.Encode())
	return rc.doRequest(ctx, path, body, &res, "application/x-www-form-urlencoded")
}

// postForm post form with x-www-form-urlencoded format
// @param ctx context
// @param path url path e.g.  /a/b/c
// @param req url request struct
// @param res response struct
func (rc *RongCloud) postForm(ctx context.Context, path string, req, res interface{}) (*http.Response, error) {
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	return rc.postFormUrlencoded(ctx, path, params, res)
}

type httpResponseGetter interface {
	GetHttpResponse() *http.Response
	GetRequestId() string
}

type rawHttpResponseGetter struct {
	rawHttpResponseInternal *http.Response
}

func newRawHttpResponseGetter(rawHttpResponseInternal *http.Response) *rawHttpResponseGetter {
	return &rawHttpResponseGetter{rawHttpResponseInternal: rawHttpResponseInternal}
}

func (r *rawHttpResponseGetter) GetHttpResponse() *http.Response {
	return r.rawHttpResponseInternal
}

func (r *rawHttpResponseGetter) GetRequestId() string {
	return r.rawHttpResponseInternal.Header.Get(RCRequestIdHeader)
}

type contextHttpHeaderKey string

func (c contextHttpHeaderKey) String() string {
	return RCHttpHeaderPrefix + string(c)
}

const (
	RCHttpHeaderPrefix = "rc-http-header-" // http header context prefix for inhibiting context conflict
	RCRequestIdHeader  = "X-Request-Id"
)

var (
	// ContextRequestIdKey represent RongCloud http request id, passing to server api
	ContextRequestIdKey = contextHttpHeaderKey(RCRequestIdHeader)
)

func AddHttpRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, ContextRequestIdKey, requestId)
}
