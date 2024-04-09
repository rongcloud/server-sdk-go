package rongcloud

import "context"

type ConversationTopSetRequest struct {
	// [必传] 用户 ID，会话所属的用户
	UserId *string `url:"userId,omitempty"`
	// [必传] 会话类型。支持的会话类型包括：1（二人会话）、3（群组会话）、6（系统会话）、10（超级群会话）。
	ConversationType *int `url:"conversationType,omitempty"`
	// [必传] 需要设置的目标 ID，根据会话类型不同为单聊用户 ID、群聊 ID、系统目标 ID
	TargetId *string `url:"targetId,omitempty"`
	// [必传] true 表示置顶，false 表示取消置顶。
	SetTop *bool `url:"setTop,omitempty"`
}

type ConversationTopSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ConversationTopSet 会话置顶
// More details see https://doc.rongcloud.cn/imserver/server/v1/conversation/top
func (rc *RongCloud) ConversationTopSet(ctx context.Context, req *ConversationTopSetRequest) (*ConversationTopSetResponse, error) {
	path := "/conversation/top/set.json"
	resp := &ConversationTopSetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ConversationNotificationSetRequest struct {
	// [必传] 会话类型。支持的会话类型包括：1（二人会话）、3（群组会话）、6（系统会话）、10（超级群会话）。
	ConversationType *int `url:"conversationType,omitempty"`
	// [必传] 设置消息免打扰的用户 ID。
	RequestId *string `url:"requestId,omitempty"`
	// [必传] 目标 ID，根据不同的会话类型（ConversationType），可能是用户 ID、群组 ID、超级群 ID 等。
	TargetId *string `url:"targetId,omitempty"`
	// 消息免打扰设置状态，0 表示关闭，1 表示开启。
	// 该字段已废弃，推荐使用 unpushLevel。isMuted 与 unpushLevel 只需要传一个。如果都传，使用 unpushLevel。
	// Deprecated
	IsMuted *int `url:"isMuted,omitempty"`
	// 超级群的会话频道 ID。
	// 如果传入频道 ID，则针对该指定频道设置消息免打扰级别。
	// 注意：2022.09.01 之前开通超级群业务的客户，如果不指定频道 ID，则默认传 "" 空字符串，即仅针对指定超级群会话（targetId）中不属于任何频道的消息设置免打扰状态级别。如需修改请提交工单。
	BusChannel *string `url:"busChannel,omitempty"`
	// -1： 全部消息通知
	//  0： 未设置（用户未设置情况下，默认以群 或者 APP级别的默认设置为准，如未设置则全部消息都通知）
	//  1： 仅针对 @ 消息进行通知
	//  2： 仅针对 @ 指定用户进行通知
	//  如：@张三 则张三可以收到推送，@所有人 时不会收到推送。
	//  4： 仅针对 @ 群全员进行通知，只接收 @所有人 的推送信息。
	//  5： 不接收通知
	UnpushLevel *int `url:"unpushLevel,omitempty"`
}

type ConversationNotificationSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ConversationNotificationSet 设置指定会话免打扰
// More details see https://doc.rongcloud.cn/imserver/server/v1/conversation/set-do-not-disturb-by-id
func (rc *RongCloud) ConversationNotificationSet(ctx context.Context, req *ConversationNotificationSetRequest) (*ConversationNotificationSetResponse, error) {
	path := "/conversation/notification/set.json"
	resp := &ConversationNotificationSetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ConversationNotificationGetRequest struct {
	// [必传]  会话类型。支持的会话类型包括：1（二人会话）、3（群组会话）、6（系统会话）、10（超级群会话）。如需查询超级群指定频道的免打扰设置，请传入频道 ID（busChannel）。
	ConversationType *int `url:"conversationType,omitempty"`
	// [必传] 设置消息免打扰的用户 ID。
	RequestId *string `url:"requestId,omitempty"`
	// [必传] 目标 ID，根据不同的会话类型（ConversationType），可能是用户 ID、群组 ID。
	TargetId *string `url:"targetId,omitempty"`
	// 超级群的会话频道 ID。
	//  如果传入频道 ID，则查询该频道的消息免打扰级别。
	//  注意：如果不指定频道 ID，则查询指定超级群会话（targetId）中不属于任何频道的消息的免打扰状态级别。
	BusChannel *string `url:"busChannel,omitempty"`
}

type ConversationNotificationGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	// 消息免打扰设置状态。
	// -1： 全部消息通知
	//  0： 未设置（用户未设置情况下，默认以群 或者 APP级别的默认设置为准，如未设置则全部消息都通知）
	//  1： 仅针对 @ 消息进行通知
	//  2： 仅针对 @ 指定用户进行通知
	//  4： 仅针对 @ 群全员进行通知。
	//  5： 不接收通知
	IsMuted int `json:"isMuted"`
}

// ConversationNotificationGet 查询指定会话免打扰
// More details see https://doc.rongcloud.cn/imserver/server/v1/conversation/get-do-not-disturb-by-id
func (rc *RongCloud) ConversationNotificationGet(ctx context.Context, req *ConversationNotificationGetRequest) (*ConversationNotificationGetResponse, error) {
	path := "/conversation/notification/get.json"
	resp := &ConversationNotificationGetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ConversationTypeNotificationSetRequest struct {
	// [必传] 会话类型。支持的会话类型包括：1（二人会话）、3（群组会话）、6（系统会话）、10（超级群会话）。
	ConversationType *int `url:"conversationType,omitempty"`
	// [必传] 设置消息免打扰的用户 ID。
	RequestId *string `url:"requestId,omitempty"`
	// [必传]  -1： 全部消息通知
	//0： 未设置（用户未设置情况下，默认以群或者 APP 级别的默认设置为准，如未设置则全部消息都通知）
	//1： 仅针对 @ 消息进行通知
	//2： 仅针对 @ 指定用户进行通知
	//如：@张三 则张三可以收到推送，@所有人 时不会收到推送。
	//4： 仅针对 @ 群全员进行通知，只接收 @所有人 的推送信息。
	//5： 不接收通知
	UnpushLevel *int `url:"unpushLevel,omitempty"`
}

type ConversationTypeNotificationSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ConversationTypeNotificationSet 设置指定会话类型免打扰
// More details see https://doc.rongcloud.cn/imserver/server/v1/conversation/set-do-not-disturb-by-type
func (rc *RongCloud) ConversationTypeNotificationSet(ctx context.Context, req *ConversationTypeNotificationSetRequest) (*ConversationTypeNotificationSetResponse, error) {
	path := "/conversation/type/notification/set.json"
	resp := &ConversationTypeNotificationSetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ConversationTypeNotificationGetRequest struct {
	// [必传] 会话类型。支持的会话类型包括：1（二人会话）、3（群组会话）、6（系统会话）、10（超级群会话）。
	ConversationType *int `url:"conversationType,omitempty"`
	// [必传] 设置消息免打扰的用户 ID。
	RequestId *string `url:"requestId,omitempty"`
}

type ConversationTypeNotificationGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	// -1： 全部消息通知
	//  0： 未设置（用户未设置情况下，默认以群 或者 APP级别的默认设置为准，如未设置则全部消息都通知）
	//  1： 仅针对 @ 消息进行通知
	//  2： 仅针对 @ 指定用户进行通知
	//  4： 仅针对 @ 群全员进行通知。
	//  5： 不接收通知
	IsMuted int `json:"isMuted"`
}

// ConversationTypeNotificationGet 查询指定会话类型免打扰
// More details see https://doc.rongcloud.cn/imserver/server/v1/conversation/get-do-not-disturb-by-type
func (rc *RongCloud) ConversationTypeNotificationGet(ctx context.Context, req *ConversationTypeNotificationGetRequest) (*ConversationTypeNotificationGetResponse, error) {
	path := "/conversation/type/notification/get.json"
	resp := &ConversationTypeNotificationGetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
