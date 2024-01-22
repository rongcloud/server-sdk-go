package rongcloud

import (
	"context"
)

// chatroom 聊天室

type ChatroomUserExistRequest struct {
	ChatroomId string `json:"chatroomId"`
	UserId     string `json:"userId"`
}

type ChatroomUserExistResponse struct {
	// 200：成功。
	Code int `json:"code"`

	// 用户是否在聊天室中，true 表示在聊天室中，false 表示不在聊天室中。
	IsInChrm bool `json:"isInChrm"`

	HttpResponseGetter `json:"-"`
}

// ChatroomUserExist
// @param ctx context with tracing, custom http request header
// @param req ChatroomUserExistRequest
// @return ChatroomUserExistResponse, error
func (rc *RongCloud) ChatroomUserExist(ctx context.Context, req *ChatroomUserExistRequest) (*ChatroomUserExistResponse, error) {
	var (
		resp *ChatroomUserExistResponse
	)
	httpResponse, err := rc.postJson(ctx, "/chatroom/user/exist.json", req, &resp)
	if err != nil {
		return nil, err
	}
	resp.HttpResponseGetter = &RawHttpResponseGetter{rawHttpResponseInternal: httpResponse}
	return resp, nil
}
