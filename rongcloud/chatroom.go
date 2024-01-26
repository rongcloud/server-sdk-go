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
	ChatroomId *string `json:"chatroomId"`
	UserId     *string `json:"userId"`
}

type ChatroomUserExistResponse struct {
	httpResponseGetter `json:"-"`
	CodeResult
	// 用户是否在聊天室中，true 表示在聊天室中，false 表示不在聊天室中。
	IsInChrm bool `json:"isInChrm"`
}

// ChatroomUserExist checking whether the member is in the chatroom.
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

// ChatroomDestroySet set chatroom destroy type
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
	ChatroomId *string `json:"chatroomId"` // 聊天室 ID
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

// ChatroomGet get chatroom
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
	ChatroomId *string `json:"chatroomId"` // 聊天室 ID
	UserId     *string `json:"userId"`     // 操作用户 ID。通过 Server API 非聊天室中用户可以进行设置。
	Key        *string `json:"key"`        // 聊天室属性名称，Key 支持大小写英文字母、数字、部分特殊符号 + = - _ 的组合方式，大小写敏感。最大长度 128 字符。每个聊天室中，最多允许设置 100 个属性 Key-Value 对。
	Value      *string `json:"value"`      // 聊天室属性对应的值，最大长度 4096 个字符。
	AutoDelete *int    `json:"autoDelete"` // 属性的操作用户退出聊天室后，是否删除此 Key 值。为 1 时删除此 Key 值和对应的 Value，为 0 时用户退出后不删除，默认为 0。
	RCMsg      RCMsg   `json:"rcMsg"`      // 聊天室属性变化通知消息的消息类型，一般为内置消息类型 RC:chrmKVNotiMsg，也可以是其他自定义消息类型。如果传入该字段，则在聊天室属性变化时发送一条消息。
}

type ChatroomEntrySetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomEntrySet set chatroom entry kv
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/set-kv-entry
func (rc *RongCloud) ChatroomEntrySet(ctx context.Context, req *ChatroomEntrySetRequest) (*ChatroomEntrySetResponse, error) {
	params := url.Values{}
	if req.ChatroomId != nil {
		params.Set("chatroomId", StringValue(req.ChatroomId))
	}
	if req.UserId != nil {
		params.Set("userId", StringValue(req.UserId))
	}
	if req.Key != nil {
		params.Set("key", StringValue(req.Key))
	}
	if req.Value != nil {
		params.Set("value", StringValue(req.Value))
	}
	if req.AutoDelete != nil {
		params.Set("autoDelete", strconv.Itoa(IntValue(req.AutoDelete)))
	}
	if req.RCMsg != nil {
		params.Set("objectName", req.RCMsg.ObjectName())
		content, err := req.RCMsg.ToString()
		if err != nil {
			return nil, fmt.Errorf("RCMsg ToString() error: %w", err)
		}
		params.Set("content", content)
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

// ChatroomEntryBatchSet batch set chatroom kv entry
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

type ChatroomDestroyRequest struct {
	ChatroomIds []string `json:"chatroomIds"` // 要销毁的聊天室的 ID。每次可销毁多个聊天室。
}

type ChatroomDestroyResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomDestroy destroy chatroom
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
	ChatroomId *string `json:"chatroomId"` // 要查询的聊天室 ID

	Count *int `json:"count"` // 要获取的聊天室成员信息数，最多返回 500 个成员信息

	Order *int `json:"order"` // 加入聊天室的先后顺序， 1 为加入时间正序， 2 为加入时间倒序
}

type ChatroomUserQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`

	Total int      `json:"total"`
	Users []string `json:"users"`
	Id    string   `json:"id"`
	Time  string   `json:"time"`
}

// ChatroomUserQuery
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
	ChatroomId *string  `json:"chatroomId"` // 要查询的聊天室 ID
	UserIds    []string `json:"userIds"`    // 要查询的用户 ID，每次最多不超过 1000 个用户 ID
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

// ChatroomUsersExist
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

// ChatroomUserBlockAdd
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
	UserIds    []string `json:"userIds"`    // 用户 ID，可同时移除多个用户，最多不超过 20 个。
	ChatroomId *string  `json:"chatroomId"` // 聊天室 ID。
	Extra      *string  `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。
	NeedNotify *bool    `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：被解除封禁的成员。
}

type ChatroomUserBlockRollbackResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBlockRollback
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
	ChatroomId *string `json:"chatroomId"` // 聊天室 ID
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

// ChatroomUserBlockList
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
	UserIds    []string `json:"userIds"`    // 用户 ID，可同时禁言多个用户，每次最多不超过 20 个。
	Minute     *int     `json:"minute"`     // 禁言时长，以分钟为单位，最大值为 43200 分钟。
	Extra      *string  `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。
	NeedNotify *bool    `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：指定聊天室中所有成员。
}

type ChatroomUserBanAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBanAdd
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
	UserIds []string `json:"userIds"` // 用户 ID，可同时移除多个用户，每次最多不超过 20 个。

	Extra *string `json:"extra"` // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。

	NeedNotify *bool `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：被解除全局禁言的用户。
}

type ChatroomUserBanRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserBanRemove remove chatroom ban users
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

// ChatroomUserBanQuery
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-globally-gagged-user
func (rc *RongCloud) ChatroomUserBanQuery(ctx context.Context, req *ChatroomUserBanQueryRequest) (*ChatroomUserBanQueryResponse, error) {
	resp := &ChatroomUserBanQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "/chatroom/user/ban/query.json", nil, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type ChatroomUserGagAddRequest struct {
	UserIds    []string `json:"userIds"`    // 用户 ID，可同时禁言多个用户，最多不超过 20 个。
	ChatroomId *string  `json:"chatroomId"` // 聊天室 ID。
	Minute     *int     `json:"minute"`     // 禁言时长，以分钟为单位，最大值为 43200 分钟。
	Extra      *string  `json:"extra"`      // 通知携带的 JSON 格式的扩展信息，仅在 needNotify 为 true 时有效。
	NeedNotify *bool    `json:"needNotify"` // 是否通知成员。默认 false 不通知。如果为 true，客户端会触发相应回调方法（要求 Android/iOS IMLib ≧ 5.4.5；Web IMLib ≧ 5.7.9）。通知范围：指定聊天室中所有成员。
}

type ChatroomUserGagAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// ChatroomUserGagAdd gag chatroom user
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
	UserIds []string `json:"userIds"` // 用户 ID，可同时移除多个用户，最多不超过 20 个。

	ChatroomId *string `json:"chatroomId"` // 聊天室 ID。

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
	ChatroomId *string `json:"chatroomId"` // 聊天室 ID。
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

// ChatroomMessagePriorityQuery
// More details see https://doc.rongcloud.cn/imserver/server/v1/chatroom/query-low-priority-message-type
func (rc *RongCloud) ChatroomMessagePriorityQuery(ctx context.Context, req *ChatroomMessagePriorityQueryRequest) (*ChatroomMessagePriorityQueryResponse, error) {
	resp := &ChatroomMessagePriorityQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, "", nil, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
