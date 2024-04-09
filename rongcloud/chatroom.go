package rongcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// chatroom 聊天室

type ChatroomUserExistRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传] 要查询的聊天室 ID
	UserId     *string `json:"userId"`     // [必传] 要查询的用户 ID
}

type ChatroomUserExistResponse struct {
	httpResponseGetter `json:"-"`
	CodeResult
	// 用户是否在聊天室中，true 表示在聊天室中，false 表示不在聊天室中。
	IsInChrm bool `json:"isInChrm"`
}

// ChatroomUserExist 查询是否在聊天室中
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/check-member
func (rc *RongCloud) ChatroomUserExist(ctx context.Context, req *ChatroomUserExistRequest) (*ChatroomUserExistResponse, error) {
	resp := &ChatroomUserExistResponse{}
	params := url.Values{}
	params.Set("chatroomId", StringValue(req.ChatroomId))
	params.Set("userId", StringValue(req.UserId))
	httpResponse, err := rc.postFormUrlencoded(ctx, "/chatroom/user/exist.json", params, &resp)
	if err != nil {
		return resp, err
	}
	resp.httpResponseGetter = &rawHttpResponseGetter{rawHttpResponseInternal: httpResponse}
	return resp, nil
}

type ChatroomCreateNewRequest struct {
	// [必传] 聊天室 ID。
	ChatroomId *string `json:"chatroomId"`
	// 指定聊天室的销毁类型。0：默认值，表示不活跃时销毁。默认情况下，所有聊天室的自动销毁方式均为不活跃时销毁，一旦不活跃长达到 60 分钟即被销毁，可通过 destroyTime 延长该时间。1：固定时间销毁，设置为该类型后，聊天室默认在创建 60 分钟后自动销毁，可通过 destroyTime 设置更长的存活时间。您也可以在聊天室创建成功后再设置，详见设置聊天室销毁类型。
	DestroyType *int `json:"destroyType"`
	// 设置聊天室销毁时间。在 destroyType=0 时，表示聊天室应在不活跃达到该时长时自动销毁。在 destroyType=1 时，表示聊天室应在创建以后存活时间达到该时长后自动销毁。单位为分钟，最小值 60 分钟，最大 10080 分钟（7 天）。如果未设置，默认 60 分钟。
	DestroyTime *int `json:"destroyTime"`
	// 是否禁言聊天室全体成员，默认 false。您也可以在聊天室创建成功后再设置，详见设置聊天室全体禁言。
	IsBan *bool `json:"isBan"`
	// 禁言白名单用户列表，支持批量设置，最多不超过 20 个。您也可以在聊天室创建成功后再设置，详见加入聊天室全体禁言白名单。
	WhiteUserIds []string `json:"whiteUserIds"`

	// 聊天室自定义属性的所属用户 ID。仅在开通聊天室自定义属性服务后可使用该字段，且必须与 entryInfo 字段一起使用。如果未开启服务，或者设置该字段时未同时传入 entryInfo，API 会返回创建失败。仅支持 1 个用户 ID。您也可以在聊天室创建成功后再设置，详见聊天室属性概述。
	EntryOwnerId *string           `json:"entryOwnerId"`
	EntryInfo    map[string]string `json:"entryInfo"`
}

type ChatroomCreateNewResponse struct {
	httpResponseGetter `json:"-"`
	CodeResult
}

// ChatroomCreateNew create new chatroom
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/create
func (rc *RongCloud) ChatroomCreateNew(ctx context.Context, req *ChatroomCreateNewRequest) (*ChatroomCreateNewResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.DestroyType != nil {
		params.Set("destroyType", strconv.Itoa(IntValue(req.DestroyType)))
	}
	if req.DestroyTime != nil {
		params.Set("destroyTime", strconv.Itoa(IntValue(req.DestroyTime)))
	}
	if req.IsBan != nil {
		params.Set("isBan", strconv.FormatBool(BoolValue(req.IsBan)))
	}
	if req.WhiteUserIds != nil {
		for _, item := range req.WhiteUserIds {
			params.Add("whiteUserIds", item)
		}
	}
	if req.EntryOwnerId != nil {
		params.Set("entryOwnerId", StringValue(req.EntryOwnerId))
	}
	if req.EntryInfo != nil {
		entryInfo, err := json.Marshal(req.EntryInfo)
		if err != nil {
			return nil, fmt.Errorf("marshal entryInfo err: %w", err)
		}
		params.Set("entryInfo", string(entryInfo))
	}

	resp := &ChatroomCreateNewResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/create_new.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ChatroomDestroySetRequest struct {
	// [必传] 聊天室 ID。
	ChatroomId *string `json:"chatroomId"`
	// 指定聊天室的销毁类型。0：默认值，表示不活跃时销毁。默认情况下，所有聊天室的自动销毁方式均为不活跃时销毁，一旦不活跃长达到 60 分钟即被销毁，可通过 destroyTime 延长该时间。1：固定时间销毁，设置为该类型后，聊天室默认在创建 60 分钟后自动销毁，可通过 destroyTime 设置更长的存活时间。您也可以在聊天室创建成功后再设置，详见设置聊天室销毁类型。
	DestroyType *int `json:"destroyType"`
	// 设置聊天室销毁时间。在 destroyType=0 时，表示聊天室应在不活跃达到该时长时自动销毁。在 destroyType=1 时，表示聊天室应在创建以后存活时间达到该时长后自动销毁。单位为分钟，最小值 60 分钟，最大 10080 分钟（7 天）。如果未设置，默认 60 分钟。
	DestroyTime *int `json:"destroyTime"`
}

type ChatroomDestroySetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomDestroySet 设置聊天室销毁类型
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/set-destroy-type
func (rc *RongCloud) ChatroomDestroySet(ctx context.Context, req *ChatroomDestroySetRequest) (*ChatroomDestroySetResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.DestroyType != nil {
		params.Set("destroyType", strconv.Itoa(IntValue(req.DestroyType)))
	}
	if req.DestroyTime != nil {
		params.Set("destroyTime", strconv.Itoa(IntValue(req.DestroyTime)))
	}
	resp := &ChatroomDestroySetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/destroy/set.json", params, resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ChatroomGetRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传] 聊天室 ID
}

type ChatroomGetResponse struct {
	CodeResult
	ChatroomId  string `json:"chatroomId"`  // 聊天室 ID
	CreateTime  int64  `json:"createTime"`  // 聊天室创建时间
	MemberCount int    `json:"memberCount"` // 聊天室当前人数
	DestroyType int    `json:"destroyType"` // 指定聊天室的销毁方式。0：默认值，表示不活跃时销毁。默认情况下，所有聊天室的自动销毁方式均为不活跃时销毁，一旦不活跃长达到 60 分钟即被销毁，可通过 destroyTime 延长该时间。1：固定时间销毁，设置为该类型后，聊天室默认在创建 60 分钟后自动销毁，可通过 destroyTime 设置更长的存活时间。
	DestroyTime int    `json:"destroyTime"` // 设置聊天室销毁等待时间。在 destroyType=0 时，表示聊天室应在不活跃达到该时长时自动销毁。在 destroyType=1 时，表示聊天室应在创建以后存活时间达到该时长后自动销毁。单位为分钟，最小值 60 分钟，最大 10080 分钟（7 天）。
	IsBan       bool   `json:"ban"`         // 是否已开启聊天室全体禁言，默认 false。

	httpResponseGetter
}

// ChatroomGet 查询聊天室信息
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/get
func (rc *RongCloud) ChatroomGet(ctx context.Context, req *ChatroomGetRequest) (*ChatroomGetResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	resp := &ChatroomGetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/get.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ChatroomEntrySetRequest struct {
	ChatroomId *string `url:"chatroomId,omitempty"` // [必传] 聊天室 ID
	UserId     *string `url:"userId,omitempty"`     // [必传] 操作用户 ID。通过 Server API 非聊天室中用户可以进行设置。
	Key        *string `url:"key,omitempty"`        // [必传] 聊天室属性名称，Key 支持大小写英文字母、数字、部分特殊符号 + = - _ 的组合方式，大小写敏感。最大长度 128 字符。每个聊天室中，最多允许设置 100 个属性 Key-Value 对。
	Value      *string `url:"value,omitempty"`      // [必传] 聊天室属性对应的值，最大长度 4096 个字符。
	AutoDelete *int    `url:"autoDelete,omitempty"` // 属性的操作用户退出聊天室后，是否删除此 Key 值。为 1 时删除此 Key 值和对应的 Value，为 0 时用户退出后不删除，默认为 0。
	RCMsg      RCMsg   `url:"-"`                    // 聊天室属性变化通知消息的消息类型，一般为内置消息类型 RC:chrmKVNotiMsg，也可以是其他自定义消息类型。如果传入该字段，则在聊天室属性变化时发送一条消息。
}

type ChatroomEntrySetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomEntrySet set chatroom entry kv (设置聊天室属性（KV）)
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/set-kv-entry
func (rc *RongCloud) ChatroomEntrySet(ctx context.Context, req *ChatroomEntrySetRequest) (*ChatroomEntrySetResponse, error) {
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	if req.RCMsg != nil {
		err = makeRCMsgUrlValues(req.RCMsg, params)
		if err != nil {
			return nil, err
		}
	}
	resp := &ChatroomEntrySetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/entry/set.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ChatroomEntryBatchSetRequest struct {
	ChatroomId *string `json:"chatroomId"` // 聊天室 ID

	AutoDelete *int `json:"autoDelete"` // 用户（entryOwnerId）退出聊天室后，是否删除此 Key 值。为 1 时删除此 Key 值和对应的 Value，为 0 时用户退出后不删除，默认为 0。

	EntryOwnerId *string `json:"entryOwnerId"` // 聊天室自定义属性的所属用户 ID

	// 聊天室自定义属性 KV 对，JSON 结构，一次最多 20 个 KV。Key 为属性名，支持大小写英文字母、数字、部分特殊符号 + = - _ 的组合方式，大小写敏感。最大长度 128 字符。Value 为属性值，最大长度 4096 个字符。
	EntryInfo map[string]string `json:"entryInfo"`
}

type ChatroomEntryBatchSetResponse struct {
	httpResponseGetter `json:"-"`
	CodeResult
}

// ChatroomEntryBatchSet batch set chatroom kv entry  批量设置聊天室属性（KV）
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/set-kv-entry-batch
func (rc *RongCloud) ChatroomEntryBatchSet(ctx context.Context, req *ChatroomEntryBatchSetRequest) (*ChatroomEntryBatchSetResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.AutoDelete != nil {
		params.Set("autoDelete", strconv.Itoa(IntValue(req.AutoDelete)))
	}
	if req.EntryOwnerId != nil {
		params.Set("entryOwnerId", StringValue(req.EntryOwnerId))
	}
	if req.EntryInfo != nil {
		entryInfo, err := json.Marshal(req.EntryInfo)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal entryInfo err %w", err)
		}
		params.Set("entryInfo", string(entryInfo))
	}
	resp := &ChatroomEntryBatchSetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/entry/batch/set.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ChatroomEntryRemoveRequest struct {
	ChatroomId *string `url:"chatroomId,omitempty"` // [必传] 聊天室 ID。
	UserId     *string `url:"userId,omitempty"`     // [必传] 操作用户 ID。通过 Server API 非聊天室中用户可以进行设置。
	Key        *string `url:"key,omitempty"`        // [必传] 聊天室属性名称，Key 支持大小写英文字母、数字、部分特殊符号 + = - _ 的组合方式，大小写敏感。最大长度 128 字。
	RCMsg      RCMsg   `url:"-"`                    // 通聊天室属性变化通知消息的消息类型，一般为内置消息类型 ChrmKVNotiMsg，也可以是其他自定义消息类型。如果传入该字段，则在聊天室属性变化时发送一条消息。
}

type ChatroomEntryRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomEntryRemove 删除聊天室属性（KV）
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/remove-kv-entry
func (rc *RongCloud) ChatroomEntryRemove(ctx context.Context, req *ChatroomEntryRemoveRequest) (*ChatroomEntryRemoveResponse, error) {
	path := "/chatroom/entry/remove.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	if req.RCMsg != nil {
		err = makeRCMsgUrlValues(req.RCMsg, params)
		if err != nil {
			return nil, err
		}
	}
	resp := &ChatroomEntryRemoveResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomEntryQueryRequest struct {
	ChatroomId *string  `json:"chatroomId"` // [必传] 聊天室 ID。
	Keys       []string `json:"keys"`       // 批量获取指定聊天室中的 Key 值，最多上限为 100 个，不传时获取全部 key 值。
}

type ChatroomEntryQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Keys               []*ChatroomEntry `json:"keys"`
}

type ChatroomEntry struct {
	Key         string `json:"key"`         // 设置的属性名。
	Value       string `json:"value"`       // 属性对应的内容。
	UserId      string `json:"userId"`      // 最后一次设置此 Key 的用户 ID。
	AutoDelete  string `json:"autoDelete"`  // 用户退出聊天室后是否删除此 Key，"0" 为不删除、"1"为删除。
	LastSetTime string `json:"lastSetTime"` // 最近一次设置 Key 的时间。
}

// ChatroomEntryQuery 查询聊天室属性（KV）
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-kv-entry
func (rc *RongCloud) ChatroomEntryQuery(ctx context.Context, req *ChatroomEntryQueryRequest) (*ChatroomEntryQueryResponse, error) {
	path := "/chatroom/entry/query.json"
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.Keys != nil {
		for _, key := range req.Keys {
			params.Set("keys", key)
		}
	}
	resp := &ChatroomEntryQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomDestroyRequest struct {
	ChatroomIds []string `json:"chatroomIds"` // 要销毁的聊天室的 ID。每次可销毁多个聊天室。
}

type ChatroomDestroyResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomDestroy 销毁聊天室
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/destroy
func (rc *RongCloud) ChatroomDestroy(ctx context.Context, req *ChatroomDestroyRequest) (*ChatroomDestroyResponse, error) {
	params := url.Values{}
	if req.ChatroomIds != nil {
		for _, id := range req.ChatroomIds {
			params.Add("chatroomId", id)
		}
	}
	resp := &ChatroomDestroyResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/destroy.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ChatroomUserQueryRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传] 要查询的聊天室 ID

	Count *int `json:"count"` // [必传] 要获取的聊天室成员信息数，最多返回 500 个成员信息

	Order *int `json:"order"` // [必传]加入聊天室的先后顺序， 1 为加入时间正序， 2 为加入时间倒序
}

type ChatroomUserQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`

	Total int      `json:"total"`
	Users []string `json:"users"`
	Id    string   `json:"id"`
	Time  string   `json:"time"`
}

// ChatroomUserQuery 获取聊天室成员
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-member-list
func (rc *RongCloud) ChatroomUserQuery(ctx context.Context, req *ChatroomUserQueryRequest) (*ChatroomUserQueryResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.Count != nil {
		params.Set("count", strconv.Itoa(IntValue(req.Count)))
	}
	if req.Order != nil {
		params.Set("order", strconv.Itoa(IntValue(req.Order)))
	}
	resp := &ChatroomUserQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/query.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ChatroomUsersExistRequest struct {
	ChatroomId *string  `json:"chatroomId"` // [必传] 要查询的聊天室 ID
	UserIds    []string `json:"userIds"`    // [必传] 要查询的用户 ID，每次最多不超过 1000 个用户 ID
}

type ChatroomUsersExistResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`

	Result []*ChatroomUserExistResult `json:"result"`
}

type ChatroomUserExistResult struct {
	UserId string `json:"userId"` // 聊天室中用户 ID。

	IsInChrm int `json:"isInChrm"` // 用户是否在聊天室中，1 表示在聊天室中，0 表示不在聊天室中。
}

// ChatroomUsersExist 批量查询是否在聊天室中
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/check-members
func (rc *RongCloud) ChatroomUsersExist(ctx context.Context, req *ChatroomUsersExistRequest) (*ChatroomUsersExistResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	resp := &ChatroomUsersExistResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/users/exist.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ChatroomUserBlockAddRequest struct {
	UserIds    []string `json:"userIds"`    // [必传]用户 ID，可同时封禁多个用户，最多不超过 20 个。
	ChatroomId *string  `json:"chatroomId"` // [必传] 聊天室 ID。
	Minute     *int     `json:"minute"`     // [必传] 封禁时长，以分钟为单位，最大值为43200分钟。
	Extra      *string  `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。
	NeedNotify *bool    `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：聊天室中所有成员，包括被封禁用户。
}

type ChatroomUserBlockAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBlockAdd 封禁聊天室用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/block-user
func (rc *RongCloud) ChatroomUserBlockAdd(ctx context.Context, req *ChatroomUserBlockAddRequest) (*ChatroomUserBlockAddResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	if req.Minute != nil {
		params.Set("minute", strconv.Itoa(IntValue(req.Minute)))
	}
	if req.Extra != nil {
		params.Set("extra", StringValue(req.Extra))
	}
	if req.NeedNotify != nil {
		params.Set("needNotify", strconv.FormatBool(BoolValue(req.NeedNotify)))
	}
	resp := &ChatroomUserBlockAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/block/add.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserBlockRollbackRequest struct {
	UserIds    []string `json:"userIds"`    // [必传] 用户 ID，可同时移除多个用户，最多不超过 20 个。
	ChatroomId *string  `json:"chatroomId"` // [必传] 聊天室 ID。
	Extra      *string  `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。
	NeedNotify *bool    `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：被解除封禁的成员。
}

type ChatroomUserBlockRollbackResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBlockRollback 解除封禁聊天室用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/unblock-user
func (rc *RongCloud) ChatroomUserBlockRollback(ctx context.Context, req *ChatroomUserBlockRollbackRequest) (*ChatroomUserBlockRollbackResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	if req.Extra != nil {
		params.Set("extra", StringValue(req.Extra))
	}
	if req.NeedNotify != nil {
		params.Set("needNotify", strconv.FormatBool(BoolValue(req.NeedNotify)))
	}
	resp := &ChatroomUserBlockRollbackResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/block/rollback.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserBlockListRequest struct {
	ChatroomId *string `json:"chatroomId"` //  [必传]聊天室 ID
}

type ChatroomUserBlockListResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []*ChatroomUserBlockUser `json:"users"` // 被封禁用户数组。
}

type ChatroomUserBlockUser struct {
	Time   string `json:"time"`   // 解禁时间。精确到秒，格式为 YYYY-MM-DD HH:MM:SS，例如 2022-09-25 16:12:38。注意：time 的值与应用所属数据中心有关。如您的 App 业务使用国内数据中心，则 time 为北京时间。如您的 App 业务使用海外数据中心，则 time 为 UTC 时间。
	UserId string `json:"userId"` // 被封禁用户 ID。
}

// ChatroomUserBlockList 查询聊天室封禁用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-blocked-user
func (rc *RongCloud) ChatroomUserBlockList(ctx context.Context, req *ChatroomUserBlockListRequest) (*ChatroomUserBlockListResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	resp := &ChatroomUserBlockListResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/block/list.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserBanAddRequest struct {
	UserIds    []string `json:"userIds"`    // [必传]用户 ID，可同时禁言多个用户，每次最多不超过 20 个。
	Minute     *int     `json:"minute"`     // [必传]禁言时长，以分钟为单位，最大值为 43200 分钟。
	Extra      *string  `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。
	NeedNotify *bool    `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：指定聊天室中所有成员。
}

type ChatroomUserBanAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBanAdd 全局禁言用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/gag-user-globally
func (rc *RongCloud) ChatroomUserBanAdd(ctx context.Context, req *ChatroomUserBanAddRequest) (*ChatroomUserBanAddResponse, error) {
	params := url.Values{}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	if req.Minute != nil {
		params.Set("minute", strconv.Itoa(IntValue(req.Minute)))
	}
	if req.Extra != nil {
		params.Set("extra", StringValue(req.Extra))
	}
	if req.NeedNotify != nil {
		params.Set("needNotify", strconv.FormatBool(BoolValue(req.NeedNotify)))
	}
	resp := &ChatroomUserBanAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/ban/add.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserBanRemoveRequest struct {
	UserIds []string `json:"userIds"` // [必传]用户 ID，可同时移除多个用户，每次最多不超过 20 个。

	Extra *string `json:"extra"` // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。

	NeedNotify *bool `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：被解除全局禁言的用户。
}

type ChatroomUserBanRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBanRemove 取消全局禁言用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/ungag-user-globally
func (rc *RongCloud) ChatroomUserBanRemove(ctx context.Context, req *ChatroomUserBanRemoveRequest) (*ChatroomUserBanRemoveResponse, error) {
	params := url.Values{}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	if req.Extra != nil {
		params.Set("extra", StringValue(req.Extra))
	}
	if req.NeedNotify != nil {
		params.Set("needNotify", strconv.FormatBool(BoolValue(req.NeedNotify)))
	}
	resp := &ChatroomUserBanRemoveResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/ban/remove.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserBanQueryRequest struct {
	// 暂无请求参数(预留)
}

type ChatroomUserBanQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []ChatroomUserBlockUser `json:"users"`
}

// ChatroomUserBanQuery 查询全局禁言用户列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-globally-gagged-user
func (rc *RongCloud) ChatroomUserBanQuery(ctx context.Context, req *ChatroomUserBanQueryRequest) (*ChatroomUserBanQueryResponse, error) {
	resp := &ChatroomUserBanQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/ban/query.json", nil, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserGagAddRequest struct {
	UserIds    []string `json:"userIds"`    // [必传]用户 ID，可同时禁言多个用户，最多不超过 20 个。
	ChatroomId *string  `json:"chatroomId"` // [必传]聊天室 ID。
	Minute     *int     `json:"minute"`     // [必传]禁言时长，以分钟为单位，最大值为 43200 分钟。
	Extra      *string  `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。
	NeedNotify *bool    `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：指定聊天室中所有成员。
}

type ChatroomUserGagAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserGagAdd 禁言指定聊天室用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/gag-user
func (rc *RongCloud) ChatroomUserGagAdd(ctx context.Context, req *ChatroomUserGagAddRequest) (*ChatroomUserGagAddResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	if req.Minute != nil {
		params.Set("minute", strconv.Itoa(IntValue(req.Minute)))
	}
	if req.Extra != nil {
		params.Set("extra", StringValue(req.Extra))
	}
	if req.NeedNotify != nil {
		params.Set("needNotify", strconv.FormatBool(BoolValue(req.NeedNotify)))
	}
	resp := &ChatroomUserGagAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/gag/add.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserGagRollbackRequest struct {
	UserIds []string `json:"userIds"` // [必传]用户 ID，可同时移除多个用户，最多不超过 20 个。

	ChatroomId *string `json:"chatroomId"` // [必传]聊天室 ID。

	Extra *string `json:"extra"` // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。

	NeedNotify *bool `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：指定聊天室中所有成员。

}

type ChatroomUserGagRollbackResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserGagRollback 取消禁言指定聊天室用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/ungag-user
func (rc *RongCloud) ChatroomUserGagRollback(ctx context.Context, req *ChatroomUserGagRollbackRequest) (*ChatroomUserGagRollbackResponse, error) {
	params := url.Values{}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.Extra != nil {
		params.Set("extra", StringValue(req.Extra))
	}
	if req.NeedNotify != nil {
		params.Set("needNotify", strconv.FormatBool(BoolValue(req.NeedNotify)))
	}
	resp := &ChatroomUserGagRollbackResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/gag/rollback.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserGagListRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传]聊天室 ID。
}

type ChatroomUserGagListResponse struct {
	CodeResult
	httpResponseGetter
	Users []*ChatroomUserBlockUser `json:"users"` // 被禁言用户数组。
}

// ChatroomUserGagList 查询聊天室禁言用户列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-gagged-user
func (rc *RongCloud) ChatroomUserGagList(ctx context.Context, req *ChatroomUserGagListRequest) (*ChatroomUserGagListResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	resp := &ChatroomUserGagListResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/gag/list.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomMessagePriorityAddRequest struct {
	ObjectNames []string `json:"objectNames"` // 低优先级的消息类型，每次最多提交 5 个，设置的消息类型最多不超过 20 个。
}

type ChatroomMessagePriorityAddResponse struct {
	CodeResult
	httpResponseGetter
}

// ChatroomMessagePriorityAdd 添加低级别消息类型
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/add-low-priority-message-type
func (rc *RongCloud) ChatroomMessagePriorityAdd(ctx context.Context, req *ChatroomMessagePriorityAddRequest) (*ChatroomMessagePriorityAddResponse, error) {
	params := url.Values{}
	if req.ObjectNames != nil {
		for _, name := range req.ObjectNames {
			params.Add("objectName", name)
		}
	}
	resp := &ChatroomMessagePriorityAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/message/priority/add.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomMessagePriorityRemoveRequest struct {
	ObjectNames []string `json:"objectNames"` // 低优先级的消息类型，每次最多提交 5 个
}

type ChatroomMessagePriorityRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomMessagePriorityRemove 移除低级别消息类型
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/remove-low-priority-message-type
func (rc *RongCloud) ChatroomMessagePriorityRemove(ctx context.Context, req *ChatroomMessagePriorityRemoveRequest) (*ChatroomMessagePriorityRemoveResponse, error) {
	params := url.Values{}
	if req.ObjectNames != nil {
		for _, name := range req.ObjectNames {
			params.Add("objectName", name)
		}
	}
	resp := &ChatroomMessagePriorityRemoveResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/message/priority/remove.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomMessagePriorityQueryRequest struct {
	// 暂无请求参数(预留)
}

type ChatroomMessagePriorityQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	ObjectNames        []string `json:"objectNames"` // 消息类型数组。
}

// ChatroomMessagePriorityQuery 查询低级别消息类型
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-low-priority-message-type
func (rc *RongCloud) ChatroomMessagePriorityQuery(ctx context.Context, req *ChatroomMessagePriorityQueryRequest) (*ChatroomMessagePriorityQueryResponse, error) {
	resp := &ChatroomMessagePriorityQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/message/priority/query.json", nil, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomKeepaliveAddRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传]聊天室 ID。
}

type ChatroomKeepaliveAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomKeepaliveAdd 保活聊天室
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/add-to-keep-alive
func (rc *RongCloud) ChatroomKeepaliveAdd(ctx context.Context, req *ChatroomKeepaliveAddRequest) (*ChatroomKeepaliveAddResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	resp := &ChatroomKeepaliveAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/keepalive/add.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomKeepaliveRemoveRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传]聊天室 ID。
}

type ChatroomKeepaliveRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomKeepaliveRemove 取消保活聊天室
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/remove-from-keep-alive
func (rc *RongCloud) ChatroomKeepaliveRemove(ctx context.Context, req *ChatroomKeepaliveRemoveRequest) (*ChatroomKeepaliveRemoveResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	resp := &ChatroomKeepaliveRemoveResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/keepalive/remove.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomKeepaliveQueryRequest struct {
	// 暂无请求参数(预留)
}

type ChatroomKeepaliveQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	ChatroomIds        []string `json:"chatroomIds,omitempty"` // 保活聊天室数组。
}

// ChatroomKeepaliveQuery 查询保活聊天室
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-keep-alive
func (rc *RongCloud) ChatroomKeepaliveQuery(ctx context.Context, req *ChatroomKeepaliveQueryRequest) (*ChatroomKeepaliveQueryResponse, error) {
	resp := &ChatroomKeepaliveQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/keepalive/query.json", nil, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomWhitelistAddRequest struct {
	ObjectNames []string `json:"objectnames"` // [必传] 消息标识，最多不超过 20 个，自定义消息类型，长度不超过 32 个字符
}

type ChatroomWhitelistAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomWhitelistAdd 加入聊天室消息白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/add-to-message-type-whitelist
func (rc *RongCloud) ChatroomWhitelistAdd(ctx context.Context, req *ChatroomWhitelistAddRequest) (*ChatroomWhitelistAddResponse, error) {
	params := url.Values{}
	if req.ObjectNames != nil {
		for _, name := range req.ObjectNames {
			params.Add("objectnames", name) // no typo
		}
	}
	resp := &ChatroomWhitelistAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/whitelist/add.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomWhitelistRemoveRequest struct {
	ObjectNames []string `json:"objectnames"` // [必传] 消息标识，最多不超过 20 个，自定义消息类型，长度不超过 32 个字符
}

type ChatroomWhitelistRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomWhitelistRemove 移出聊天室消息白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/remove-from-message-type-whitelist
func (rc *RongCloud) ChatroomWhitelistRemove(ctx context.Context, req *ChatroomWhitelistRemoveRequest) (*ChatroomWhitelistRemoveResponse, error) {
	params := url.Values{}
	if req.ObjectNames != nil {
		for _, name := range req.ObjectNames {
			params.Add("objectnames", name)
		}
	}
	resp := &ChatroomWhitelistRemoveResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/whitelist/delete.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomWhitelistQueryRequest struct {
	// 暂无请求参数(预留)
}

type ChatroomWhitelistQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	WhitelistMsgType   []string `json:"whitlistMsgType"` // 消息类型数组。
}

// ChatroomWhitelistQuery 查询聊天室消息白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-message-type-whitelist
func (rc *RongCloud) ChatroomWhitelistQuery(ctx context.Context, req *ChatroomWhitelistQueryRequest) (*ChatroomWhitelistQueryResponse, error) {
	resp := &ChatroomWhitelistQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/whitelist/query.json", nil, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserWhitelistAddRequest struct {
	ChatroomId *string  `json:"chatroomId"` // [必传] 聊天室 ID。
	UserIds    []string `json:"userId"`     // [必传] 聊天室中用户 ID，可提交多个。聊天室中白名单用户最多不超过 5 个。
}

type ChatroomUserWhitelistAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserWhitelistAdd 加入聊天室用户白名单
// https://doc.rongcloud.cn/imserver/server/v1/chatroom/add-to-user-whitelist
func (rc *RongCloud) ChatroomUserWhitelistAdd(ctx context.Context, req *ChatroomUserWhitelistAddRequest) (*ChatroomUserWhitelistAddResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	resp := &ChatroomUserWhitelistAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/whitelist/add.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserWhitelistRemoveRequest struct {
	ChatroomId *string  `json:"chatroomId"` // [必传] 聊天室 ID。
	UserIds    []string `json:"userId"`     // [必传] 聊天室中用户 ID，可提交多个。聊天室中白名单用户最多不超过 5 个。
}

type ChatroomUserWhitelistRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserWhitelistRemove 移出聊天室用户白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/remove-from-user-whitelist
func (rc *RongCloud) ChatroomUserWhitelistRemove(ctx context.Context, req *ChatroomUserWhitelistRemoveRequest) (*ChatroomUserWhitelistRemoveResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	resp := &ChatroomUserWhitelistRemoveResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/whitelist/remove.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserWhitelistQueryRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传] 聊天室 ID。
}

type ChatroomUserWhitelistQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []string `json:"users"` // 白名单用户数组。
}

// ChatroomUserWhitelistQuery 查询聊天室用户白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-user-whitelist
func (rc *RongCloud) ChatroomUserWhitelistQuery(ctx context.Context, req *ChatroomUserWhitelistQueryRequest) (*ChatroomUserWhitelistQueryResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	resp := &ChatroomUserWhitelistQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/whitelist/query.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomBanRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传] 需要设置为禁言的聊天室 ID。
	Extra      *string `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 NeedNotify 为 true 时有效。
	NeedNotify *bool   `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：指定聊天室中所有成员
}

func (req ChatroomBanRequest) toUrlValues() url.Values {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.Extra != nil {
		params.Set("extra", StringValue(req.Extra))
	}
	if req.NeedNotify != nil {
		params.Set("needNotify", strconv.FormatBool(BoolValue(req.NeedNotify)))
	}
	return params
}

type ChatroomBanAddRequest struct {
	ChatroomBanRequest `json:",inline"`
}

type ChatroomBanAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomBanAdd 设置聊天室全体禁言
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/ban-chatroom
func (rc *RongCloud) ChatroomBanAdd(ctx context.Context, req *ChatroomBanAddRequest) (*ChatroomBanAddResponse, error) {
	params := req.toUrlValues()
	resp := &ChatroomBanAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/ban/add.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomBanRollbackRequest struct {
	ChatroomBanRequest `json:",inline"`
}

type ChatroomBanRollbackResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomBanRollback 取消聊天室全体禁言
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/unban-chatroom
func (rc *RongCloud) ChatroomBanRollback(ctx context.Context, req *ChatroomBanRollbackRequest) (*ChatroomBanRollbackResponse, error) {
	params := req.toUrlValues()
	resp := &ChatroomBanRollbackResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/ban/rollback.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomBanQueryRequest struct {
	Size *int `json:"size"` // 获取聊天室禁言列表的每页条数，不传时默认为 50 条，上限为 1000 条。
	Page *int `json:"page"` // 当前页面数，不传时默认获取第 1 页。
}

type ChatroomBanQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	ChatroomIds        []string `json:"chatroomIds"` // 被全体禁言的聊天室数组。
}

// ChatroomBanQuery 查询聊天室全体禁言列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-banned-list
func (rc *RongCloud) ChatroomBanQuery(ctx context.Context, req *ChatroomBanQueryRequest) (*ChatroomBanQueryResponse, error) {
	params := url.Values{}
	if req.Size != nil {
		params.Set("size", strconv.Itoa(IntValue(req.Size)))
	}
	if req.Page != nil {
		params.Set("page", strconv.Itoa(IntValue(req.Page)))
	}
	resp := &ChatroomBanQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/ban/query.json", params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomBanCheckRequest struct {
	ChatroomId *string `json:"chatroomId"` // 要查询的聊天室 ID
}

type ChatroomBanCheckResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Status             int `json:"status"` // 禁言状态，1 为全体禁言、0 为非全体禁言
}

// ChatroomBanCheck 查询聊天室全体禁言状态
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-banned-state
func (rc *RongCloud) ChatroomBanCheck(ctx context.Context, req *ChatroomBanCheckRequest) (*ChatroomBanCheckResponse, error) {
	path := "/chatroom/ban/check.json"
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	resp := &ChatroomBanCheckResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserBanWhitelistRequest struct {
	ChatroomId *string  `json:"chatroomId"` // [必传] 聊天室 ID
	UserIds    []string `json:"userId"`     // [必传] 需要添加到白名单中的用户 ID，白名单中用户上限为 20 个，支持批量添加，单次添加上限不超过 20 个。
	Extra      *string  `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 NeedNotify 为 true 时有效。
	NeedNotify *bool    `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：指定聊天室中所有成员
}

func (req *ChatroomUserBanWhitelistRequest) toUrlValues() url.Values {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.UserIds != nil {
		for _, id := range req.UserIds {
			params.Add("userId", id)
		}
	}
	if req.Extra != nil {
		params.Set("extra", StringValue(req.Extra))
	}
	if req.NeedNotify != nil {
		params.Set("needNotify", strconv.FormatBool(BoolValue(req.NeedNotify)))
	}
	return params
}

type ChatroomUserBanWhitelistAddRequest struct {
	ChatroomUserBanWhitelistRequest `json:",inline"`
}

type ChatroomUserBanWhiteListAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBanWhitelistAdd 加入聊天室全体禁言白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/add-to-chatroom-ban-whitelist
func (rc *RongCloud) ChatroomUserBanWhitelistAdd(ctx context.Context, req *ChatroomUserBanWhitelistAddRequest) (*ChatroomUserBanWhiteListAddResponse, error) {
	path := "/chatroom/user/ban/whitelist/add.json"
	params := req.toUrlValues()
	resp := &ChatroomUserBanWhiteListAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserBanWhitelistRollbackRequest struct {
	ChatroomUserBanWhitelistRequest `json:",inline"`
}

type ChatroomUserBanWhiteListRollbackResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBanWhitelistRollback 移出聊天室全体禁言白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/remove-from-chatroom-ban-whitelist
func (rc *RongCloud) ChatroomUserBanWhitelistRollback(ctx context.Context, req *ChatroomUserBanWhitelistRollbackRequest) (*ChatroomUserBanWhiteListRollbackResponse, error) {
	path := "/chatroom/user/ban/whitelist/rollback.json"
	params := req.toUrlValues()
	resp := &ChatroomUserBanWhiteListRollbackResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserBanWhitelistQueryRequest struct {
	ChatroomId *string `json:"chatroomId"` // [必传] 聊天室 ID
}

type ChatroomUserBanWhitelistQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	UserIds            []string `json:"userIds"` // 聊天室中白名单用户数组。
}

// ChatroomUserBanWhitelistQuery 查询聊天室全体禁言白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-chatroom-ban-whitelist
func (rc *RongCloud) ChatroomUserBanWhitelistQuery(ctx context.Context, req *ChatroomUserBanWhitelistQueryRequest) (*ChatroomUserBanWhitelistQueryResponse, error) {
	path := "/chatroom/user/ban/whitelist/query.json"
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	resp := &ChatroomUserBanWhitelistQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
