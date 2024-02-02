package rongcloud

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

func makeUrlValues(v interface{}) (url.Values, error) {
	values, err := query.Values(v)
	if err != nil {
		return values, NewEncodeRequestError(err)
	}
	return values, nil
}

func MakeRCMsgUrlValues(rcMsg RCMsg, key string, v *url.Values) error {
	v.Set("objectName", rcMsg.ObjectName())
	content, err := rcMsg.ToString()
	if err != nil {
		return NewSDKError(fmt.Sprintf("%s RCMsg.ToString() error %s", rcMsg.ObjectName(), err))
	}
	v.Set("content", content)
	return nil
}
