package rongcloud

import (
	"errors"
	"fmt"
)

type SDKError struct {
	Msg string
}

func NewSDKError(msg string) error {
	return SDKError{Msg: msg}
}
func (e SDKError) Error() string {
	return fmt.Sprintf("", e.Msg)
}

// NewEncodeRequestError
// query parse values error help function
func NewEncodeRequestError(err error) error {
	return NewSDKError(fmt.Sprintf("encode request error %s", err))
}

func IsSDKError(err error) bool {
	return errors.Is(err, SDKError{})
}
