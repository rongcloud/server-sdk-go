package rongcloud

import (
	"net/http"
	"time"
)

type rongCloudOption func(*RongCloud)

// WithRongCloudSMSURI 设置融云 SMS URI
func WithRongCloudSMSURI(rongCloudSMSURI string) rongCloudOption {
	return func(o *RongCloud) {
		o.rongCloudSMSURI = rongCloudSMSURI
	}
}

// WithRongCloudURI 设置融云 URI
func WithRongCloudURI(rongCloudURI string) rongCloudOption {
	return func(o *RongCloud) {
		o.rongCloudURI = rongCloudURI
	}
}

// WithTimeout http client参数, 设置超时时间，最小单位为秒
func WithTimeout(t time.Duration) rongCloudOption {
	return func(o *RongCloud) {
		o.timeout = t
	}
}

// WithKeepAlive http client参数, 连接保活时间，最小单位为秒
func WithKeepAlive(t time.Duration) rongCloudOption {
	return func(o *RongCloud) {
		o.keepAlive = t
	}
}

// WithMaxIdleConnsPerHost http client参数, 设置每个域名最大连接数
func WithMaxIdleConnsPerHost(n int) rongCloudOption {
	return func(o *RongCloud) {
		o.maxIdleConnsPerHost = n
	}
}

// WithTransport 自定义http client Transport, 优先级大于其他http client参数
func WithTransport(transport http.RoundTripper) rongCloudOption {
	return func(o *RongCloud) {
		o.globalTransport = transport
	}
}
