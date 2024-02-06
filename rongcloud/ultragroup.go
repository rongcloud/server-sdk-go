package rongcloud

import (
	"context"
	"encoding/json"
	"fmt"
)

// TODO ultraGroupMsgModify?
// TODO /v2/ultragroups?

type UltraGroupMsgGetRequest struct {
	GroupId *string          `json:"groupId,omitempty" url:"groupId,omitempty"` // [必传] 超级群 ID
	Msgs    []*UltraGroupMsg `json:"msgs,omitempty" url:"-"`                    // [必传] 消息的查询参数，单次请求最多获取 20 条消息。
}

type UltraGroupMsg struct {
	MsgUID     string `json:"msgUID,omitempty"`     // [必传] 全局唯一 ID，即消息 UID。
	BusChannel string `json:"busChannel,omitempty"` // 消息所在的超级群频道 ID。
}

type UltraGroupMsgGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Data               []*UltraGroupMsgData `json:"data"` // 消息数组。
}

type UltraGroupMsgData struct {
	FromUserId   string `json:"fromUserId"`   // 发送人用户 ID。
	GroupId      string `json:"groupId"`      // 超级群 ID。
	SentTime     int64  `json:"sentTime"`     // 消息发送时间。Unix 时间戳，单位为毫秒。
	BusChannel   string `json:"busChannel"`   // 频道 ID。
	MsgUID       string `json:"msgUID"`       // 全局唯一消息 ID，即消息 UID。
	ObjectName   string `json:"objectName"`   // 消息类型的唯一标识。
	Content      string `json:"content"`      // 消息的内容。
	Expansion    bool   `json:"expansion"`    // 是否为扩展消息。
	ExtraContent string `json:"extraContent"` // 消息扩展的内容，JSON 结构的 Key、Value 对，如：{"type":"3"}。Key 最大 32 个字符，支持大小写英文字母、数字、 特殊字符+ = - _ 的组合方式，不支持汉字。Value 最大 4096 个字符。
}

// UltraGroupMsgGet 获取指定超级群消息内容
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/get-messages-by-uid
func (rc *RongCloud) UltraGroupMsgGet(ctx context.Context, req *UltraGroupMsgGetRequest) (*UltraGroupMsgGetResponse, error) {
	path := "/ultragroup/msg/get.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	if req.Msgs != nil {
		msgs, err := json.Marshal(req.Msgs)
		if err != nil {
			return nil, NewSDKError(fmt.Sprintf("marshal msgs error %s", err))
		}
		params.Set("msgs", string(msgs))
	}
	resp := &UltraGroupMsgGetResponse{}
	httpRes, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpRes)
	return resp, err
}

type UltraGroupCreateRequest struct {
	UserId    *string `url:"userId,omitempty"`    // [必传] 需要加入的用户 ID，创建后同时加入超级群。仅支持传入一个用户 ID。
	GroupId   *string `url:"groupId,omitempty"`   //  [必传] 超级群 ID，最大长度 64 个字符。支持大小写英文字母与数字的组合。
	GroupName *string `url:"groupName,omitempty"` // [必传] 超级群 ID 对应的名称，用于在发送群组消息显示远程 Push 通知时使用，如超级群名称改变需要调用刷新超级群信息接口同步。
}

type UltraGroupCreateResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupCreate 创建超级群
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/create-ultragroup
func (rc *RongCloud) UltraGroupCreate(ctx context.Context, req *UltraGroupCreateRequest) (*UltraGroupCreateResponse, error) {
	path := "/ultragroup/create.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupCreateResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupDisRequest struct {
	// [必传] 要解散的超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
}

type UltraGroupDisResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupDis 解散超级群
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/dismiss-ultragroup
func (rc *RongCloud) UltraGroupDis(ctx context.Context, req *UltraGroupDisRequest) (*UltraGroupDisResponse, error) {
	path := "/ultragroup/dis.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupDisResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupRefreshRequest struct {
	// [必传] 群组 ID。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 群组名称。
	GroupName *string `url:"groupName,omitempty"`
}

type UltraGroupRefreshResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupRefresh 刷新超级群信息
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/refresh-ultragroup-info
func (rc *RongCloud) UltraGroupRefresh(ctx context.Context, req *UltraGroupRefreshRequest) (*UltraGroupRefreshResponse, error) {
	path := "/ultragroup/refresh.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupRefreshResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupChannelGetRequest struct {
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// 当前页数，默认为 1。
	Page *int `url:"page,omitempty"`
	// 每页条数，默认为 20 条，最大不超过 100 条。
	Limit *int `url:"limit,omitempty"`
}

type UltraGroupChannelGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	ChannelList        []*UltraGroupChannel `json:"channelList"` // 频道列表
}

type UltraGroupChannel struct {
	ChannelId  string `json:"channelId"`  // 频道 ID
	CreateTime string `json:"createTime"` // 频道创建时间
	Type       int    `json:"type"`       // 频道类型。0 表示公有频道。1 表示私有频道。
}

// UltraGroupChannelGet 查询频道列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/get-channel-list
func (rc *RongCloud) UltraGroupChannelGet(ctx context.Context, req *UltraGroupChannelGetRequest) (*UltraGroupChannelGetResponse, error) {
	path := "/ultragroup/channel/get.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupChannelGetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupHisMsgMsgIdQueryRequest struct {
	// [必传] 超级群 ID，请确保该超级群 ID 已存在。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 频道 ID。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 消息的 UID。返回结果中会包含该 msgUID 的消息。例如，该接口默认查询该消息前面 10 条和后面 10 条消息，默认情况下返回 21 条消息。
	MsgUID *string `url:"msgUID,omitempty"`
	// 需要查询的上文消息数量。例如传入 20，即表示需要获取 msgUID 前的 20 消息数量。取值范围：0-50，默认 10。0 表示不需要获取 msgUID 前的消息。
	PrevNum *int `url:"prevNum,omitempty"`
	// 需要查询的下文消息数量。例如传入 20，即表示需要获取 msgUID 前的 20 消息数量。取值范围：0-50，默认 10。0 表示不需要获取 msgUID 后的消息。
	LastNum *int `url:"lastNum,omitempty"`
}

type UltraGroupHisMsgMsgIdQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Data               []*UltraGroupHisMsg `json:"data"` // 查询结果。按消息时间戳升序排列。
}

type UltraGroupHisMsg struct {
	GroupId          string            `json:"groupId"`          // 超级群 ID。
	BusChannel       string            `json:"busChannel"`       // 超级群频道 ID。
	FromUserId       string            `json:"fromUserId"`       // 消息发送方用户 ID。
	MsgUID           string            `json:"msgUID"`           // 消息 UID。
	MsgTime          int64             `json:"msgTime"`          // 发送消息的时间戳。Unix 时间戳（毫秒）。
	ObjectName       string            `json:"objectName"`       // 消息类型。详见消息类型概述。
	ConversationType int               `json:"conversationType"` // 会话类型。超级群的会话类型为 10。
	Content          string            `json:"content"`          // 消息内容，JSON 格式。具体结构可通过 objectName 字段在 消息类型概述 中查询。
	Expansion        bool              `json:"expansion"`        // 消息是否已被设置为可扩展消息。true 表示可扩展。false 表示不可扩展。
	ExtraContent     map[string]string `json:"extraContent"`     // 消息扩展的内容，JSON 结构的 Key、Value 对，如：{"type":"3"}。Key 最大 32 个字符，支持大小写英文字母、数字、 特殊字符+ = - _ 的组合方式，不支持汉字。Value 最大 4096 个字符。
}

// UltraGroupHismsgMsgIdQuery 搜索超级群消息上下文
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/query-history-by-message-uid
func (rc *RongCloud) UltraGroupHismsgMsgIdQuery(ctx context.Context, req *UltraGroupHisMsgMsgIdQueryRequest) (*UltraGroupHisMsgMsgIdQueryResponse, error) {
	path := "/ultragroup/hismsg/msgid/query.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupHisMsgMsgIdQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupHismsgQueryRequest struct {
	// [必传] 超级群 ID，请确保该超级群 ID 已存在。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 频道 ID。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 查询开始时间。Unix 时间戳（毫秒）。
	StartTime *int64 `url:"startTime,omitempty"`
	// [必传] 查询结束时间。Unix 时间戳（毫秒）。需要保证比 startTime 大，且两者之间时间跨度最大14 天。
	EndTime *int64 `url:"endTime,omitempty"`
	// 发送者用户 ID。不传该字段，查所有用户发送的历史消息。如果传入该参数，表示只查该用户 ID 发的历史消息。
	FromUserId *string `url:"fromUserId,omitempty"`
	// 分页返回的页面大小。默认 20 条，最大 100 条。
	PageSize *int `url:"pageSize,omitempty"`
}

type UltraGroupHismsgQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Data               []*UltraGroupMsgData `json:"data"` // 查询结果。按消息时间戳升序排列。
}

// UltraGroupHismsgQuery 搜索超级群消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/query-history-messages
func (rc *RongCloud) UltraGroupHismsgQuery(ctx context.Context, req *UltraGroupHismsgQueryRequest) (*UltraGroupHismsgQueryResponse, error) {
	path := "/ultragroup/hismsg/query.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupHismsgQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupChannelCreateRequest struct {
	// [必传] 超级群 ID，请确保该超级群 ID 已存在。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 频道 ID，支持大小写字母、数字的组合方式，长度不超过 20 个字符。
	BusChannel *string `url:"busChannel,omitempty"`
	// 频道类型。0 表示公有频道（默认）。1 表示私有频道。
	Type *int `url:"type,omitempty"`
}

type UltraGroupChannelCreateResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupChannelCreate 创建频道
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/create-channel
func (rc *RongCloud) UltraGroupChannelCreate(ctx context.Context, req *UltraGroupChannelCreateRequest) (*UltraGroupChannelCreateResponse, error) {
	path := "/ultragroup/channel/create.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupChannelCreateResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupChannelDelRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 频道 ID，支持大小写字母、数字的组合方式，长度不超过 20 个字符。
	BusChannel *string `url:"busChannel,omitempty"`
}

type UltraGroupChannelDelResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupChannelDel 删除频道
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/delete-channel
func (rc *RongCloud) UltraGroupChannelDel(ctx context.Context, req *UltraGroupChannelDelRequest) (*UltraGroupChannelDelResponse, error) {
	path := "/ultragroup/channel/del.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupChannelDelResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupChannelTypeChangeRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 频道 ID
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 频道类型，0:公有频道，1:私有频道
	Type *int `url:"type,omitempty"`
}

type UltraGroupChannelTypeChangeResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupChannelTypeChange 变更频道类型
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/change-channel-type
func (rc *RongCloud) UltraGroupChannelTypeChange(ctx context.Context, req *UltraGroupChannelTypeChangeRequest) (*UltraGroupChannelTypeChangeResponse, error) {
	path := "/ultragroup/channel/type/change.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupChannelTypeChangeResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupChannelPrivateUsersGetRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 频道 ID
	BusChannel *string `url:"busChannel,omitempty"`
	// 查询页码，默认值 1
	Page *int `url:"page,omitempty"`
	// 查询每页条数，可选值 1-100。默认 50。
	PageSize *int `url:"pageSize,omitempty"`
}

type UltraGroupChannelPrivateUsersGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []string `json:"users"` // 用户 ID 列表
}

// UltraGroupChannelPrivateUsersGet 查询私有频道成员列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/get-private-channel-users
func (rc *RongCloud) UltraGroupChannelPrivateUsersGet(ctx context.Context, req *UltraGroupChannelPrivateUsersGetRequest) (*UltraGroupChannelPrivateUsersGetResponse, error) {
	path := "/ultragroup/channel/private/users/get.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupChannelPrivateUsersGetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupChannelPrivateUsersAddRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 频道 ID
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 群内的用户 ID。单次不超过 20 个。多个用户间用英文 ',' 分割
	UserIds *string `url:"userIds"`
}

type UltraGroupChannelPrivateUsersAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupChannelPrivateUsersAdd 添加私有频道成员
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/add-to-private-channel-users
func (rc *RongCloud) UltraGroupChannelPrivateUsersAdd(ctx context.Context, req *UltraGroupChannelPrivateUsersAddRequest) (*UltraGroupChannelPrivateUsersAddResponse, error) {
	path := "/ultragroup/channel/private/users/add.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupChannelPrivateUsersAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupChannelPrivateUsersDelRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 频道 ID
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 群内的用户 ID。单次不超过 20 个。多个用户间用英文 ',' 分割
	UserIds *string `url:"userIds"`
}

type UltraGroupChannelPrivateUsersDelResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupChannelPrivateUsersDel 删除私有频道成员
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/delete-from-private-channel-users
func (rc *RongCloud) UltraGroupChannelPrivateUsersDel(ctx context.Context, req *UltraGroupChannelPrivateUsersDelRequest) (*UltraGroupChannelPrivateUsersDelResponse, error) {
	path := "/ultragroup/channel/private/users/del.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupChannelPrivateUsersDelResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupJoinRequest struct {
	// [必传] 要加入群的用户 ID。单次仅支持 1 个用户。
	UserId *string `url:"userId,omitempty"`
	// [必传] 要加入的群 ID。
	GroupId *string `url:"groupId,omitempty"`
}

type UltraGroupJoinResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupJoin 加入超级群
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/join-ultragroup
func (rc *RongCloud) UltraGroupJoin(ctx context.Context, req *UltraGroupJoinRequest) (*UltraGroupJoinResponse, error) {
	path := "/ultragroup/join.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupJoinResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupQuitRequest struct {
	// [必传] 要退出群的用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 要退出的群 ID。
	GroupId *string `url:"groupId,omitempty"`
}

type UltraGroupQuitResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupQuit 退出超级群
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/quit-ultragroup
func (rc *RongCloud) UltraGroupQuit(ctx context.Context, req *UltraGroupQuitRequest) (*UltraGroupQuitResponse, error) {
	path := "/ultragroup/quit.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupQuitResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupMessageExpansionSetRequest struct {
	// [必传] 消息唯一标识 ID，可通过全量消息路由功能获取。详见全量消息路由。
	MsgUID *string `url:"msgUID,omitempty"`
	// [必传] 操作者用户 ID，即需要为指定消息（msgUID）删除扩展信息的用户 ID。
	UserId *string `url:"userId,omitempty"`
	// 超级群频道 ID。具体使用要求如下：
	// 如果发送消息时指定了频道 ID，则必传频道 ID，否则无法成功设置扩展。
	// 如果发送消息时未指定频道 ID，则不可传入频道 ID，否则无法成功设置扩展。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 需要删除的扩展信息的 Key 值，一次最多可以删除 100 个扩展信息
	ExtraKeyVal *string `url:"extraKeyVal,omitempty"`
}

type UltraGroupMessageExpansionSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupMessageExpansionSet 设置超级群消息扩展
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/set-expansion
func (rc *RongCloud) UltraGroupMessageExpansionSet(ctx context.Context, req *UltraGroupMessageExpansionSetRequest) (*UltraGroupMessageExpansionSetResponse, error) {
	path := "/ultragroup/message/expansion/set.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupMessageExpansionSetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupMessageExpansionDeleteRequest struct {
	// [必传] 消息唯一标识 ID，可通过全量消息路由功能获取。详见全量消息路由。
	MsgUID *string `url:"msgUID,omitempty"`
	// [必传] 操作者用户 ID，即需要为指定消息（msgUID）删除扩展信息的用户 ID。
	UserId *string `url:"userId,omitempty"`
	// 超级群频道 ID。具体使用要求如下：
	// 如果发送消息时指定了频道 ID，则必传频道 ID，否则无法成功设置扩展。
	// 如果发送消息时未指定频道 ID，则不可传入频道 ID，否则无法成功设置扩展。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 需要删除的扩展信息的 Key 值，一次最多可以删除 100 个扩展信息
	ExtraKey *string `url:"extraKey,omitempty"`
}

type UltraGroupMessageExpansionDeleteResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupMessageExpansionDelete 删除超级群消息扩展
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/delete-expansion
func (rc *RongCloud) UltraGroupMessageExpansionDelete(ctx context.Context, req *UltraGroupMessageExpansionDeleteRequest) (*UltraGroupMessageExpansionDeleteResponse, error) {
	path := "/ultragroup/message/expansion/delete.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupMessageExpansionDeleteResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupMessageExpansionQueryRequest struct {
	// [必传] 消息唯一标识 ID，可通过全量消息路由功能获取。详见全量消息路由。
	MsgUID *string `url:"msgUID,omitempty"`
	// 超级群频道 ID。具体使用要求如下：
	// 如果发送消息时指定了频道 ID，则必传频道 ID，否则无法成功获取扩展。
	// 如果发送消息时未指定频道 ID，则不可传入频道 ID，否则无法成功获取扩展
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// 页数，默认返回 300 个扩展信息。
	PageNo *int `url:"pageNo,omitempty"`
}

type UltraGroupMessageExpansionQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	ExtraContent       map[string]*MessageExpansionQueryExtraContentValue `json:"extraContent"`
}

// UltraGroupMessageExpansionQuery 获取超级群消息扩展
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/get-expansion
func (rc *RongCloud) UltraGroupMessageExpansionQuery(ctx context.Context, req *UltraGroupMessageExpansionQueryRequest) (*UltraGroupMessageExpansionQueryResponse, error) {
	path := "/ultragroup/message/expansion/query.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupMessageExpansionQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupMemberExistRequest struct {
	// [必传] 要查询的用户 ID
	UserId *string `url:"userId,omitempty"`
	// [必传] 要查询的超级群 ID
	GroupId *string `url:"groupId,omitempty"`
}

type UltraGroupMemberExistResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Status             bool `json:"status"` // 用户是否在超级群中。true 表示在超级群中，false 表示不在超级群中。
}

// UltraGroupMemberExist 查询用户是否为群成员
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/query-member
func (rc *RongCloud) UltraGroupMemberExist(ctx context.Context, req *UltraGroupMemberExistRequest) (*UltraGroupMemberExistResponse, error) {
	path := "/ultragroup/member/exist.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupMemberExistResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupNotDisturbSetRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID，不传时为群的默认免打扰设置
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] -1：全部消息通知
	// 0：未设置（用户未设置时为此状态，为全部消息都通知，在此状态下，如设置了超级群默认状态以超级群的默认设置为准）
	// 1：仅针对 @ 消息进行通知，包括 @指定用户 和 @所有人
	// 2：仅针对 @ 指定用户消息进行通知，且仅通知被 @ 的指定的用户进行通知。
	// 如：@张三 则张三可以收到推送，@所有人 时不会收到推送。
	// 4：仅针对 @群全员进行通知，只接收 @所有人 的推送信息
	// 5：不接收通知，即使为 @ 消息也不推送通知。
	UnpushLevel *int `url:"unpushLevel,omitempty"`
}

type UltraGroupNotDisturbSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupNotDisturbSet 设置群/频道默认免打扰
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/set-do-not-disturb
func (rc *RongCloud) UltraGroupNotDisturbSet(ctx context.Context, req *UltraGroupNotDisturbSetRequest) (*UltraGroupNotDisturbSetResponse, error) {
	path := "/ultragroup/notdisturb/set.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupNotDisturbSetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupNotDisturbGetRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID，不传时为群的默认免打扰设置
	BusChannel *string `url:"busChannel,omitempty"`
}

type UltraGroupNotDisturbGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	// 超级群 ID
	GroupId string `json:"groupId"`
	// 频道 ID，不传时为群的默认免打扰设置
	BusChannel string `json:"busChannel"`
	// -1：全部消息通知
	//  0：未设置（用户未设置时为此状态，为全部消息都通知，在此状态下，如设置了超级群默认状态以超级群的默认设置为准）
	//  1：仅针对 @ 消息进行通知，包括 @指定用户 和 @所有人
	//  2：仅针对 @ 指定用户消息进行通知，且仅通知被 @ 的指定的用户进行通知。
	//  如：@张三 则张三可以收到推送，@所有人 时不会收到推送。
	//  4：仅针对 @群全员进行通知，只接收 @所有人 的推送信息
	//  5：不接收通知，即使为 @ 消息也不推送通知。
	UnpushLevel int `json:"unpushLevel"`
}

// UltraGroupNotDisturbGet 查询默认免打扰配置
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/get-do-not-disturb
func (rc *RongCloud) UltraGroupNotDisturbGet(ctx context.Context, req *UltraGroupNotDisturbGetRequest) (*UltraGroupNotDisturbGetResponse, error) {
	path := "/ultragroup/notdisturb/get.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupNotDisturbGetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserBannedAddRequest struct {
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 用户 ID 列表，每次最多不超过 20 个用户，以逗号分隔。
	UserIds *string `url:"userIds,omitempty"`
}

type UltraGroupUserBannedAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupUserBannedAdd 禁言指定超级群成员
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/ban-user
func (rc *RongCloud) UltraGroupUserBannedAdd(ctx context.Context, req *UltraGroupUserBannedAddRequest) (*UltraGroupUserBannedAddResponse, error) {
	path := "/ultragroup/userbanned/add.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupUserBannedAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserBannedGetRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID。
	BusChannel *string `url:"busChannel,omitempty"`
	// 当前页码，默认获取第一页。
	Page *int `url:"page,omitempty"`
	// 每页条数，默认每页 50 条。
	PageSize *int `url:"pageSize,omitempty"`
}

type UltraGroupUserBannedGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []*UltraGroupUserBannedGetUser `json:"users"`
}

type UltraGroupUserBannedGetUser struct {
	Id string `json:"id"`
}

// UltraGroupUserBannedGet 查询超级群成员禁言列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/query-banned-user-list
func (rc *RongCloud) UltraGroupUserBannedGet(ctx context.Context, req *UltraGroupUserBannedGetRequest) (*UltraGroupUserBannedGetResponse, error) {
	path := "/ultragroup/userbanned/get.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupUserBannedGetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserBannedDelRequest struct {
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 用户 ID 列表，每次最多不超过 20 个用户，以逗号分隔。
	UserIds *string `url:"userIds,omitempty"`
}

type UltraGroupUserBannedDelResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupUserBannedDel 取消指定超级群成员禁言
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/unban-user
func (rc *RongCloud) UltraGroupUserBannedDel(ctx context.Context, req *UltraGroupUserBannedDelRequest) (*UltraGroupUserBannedDelResponse, error) {
	path := "/ultragroup/userbanned/del.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupUserBannedDelResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupGlobalBannedSetRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] true 为禁言，false 为取消禁言
	Status *bool `url:"status,omitempty"`
}

type UltraGroupGlobalBannedSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupGlobalBannedSet 设置超级群全体禁言
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/set-ultragroup-global-ban
func (rc *RongCloud) UltraGroupGlobalBannedSet(ctx context.Context, req *UltraGroupGlobalBannedSetRequest) (*UltraGroupGlobalBannedSetResponse, error) {
	path := "/ultragroup/globalbanned/set.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupGlobalBannedSetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupGlobalBannedGetRequest struct {
	// [必传] 超级群 ID
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID。
	BusChannel *string `url:"busChannel,omitempty"`
}

type UltraGroupGlobalBannedGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Status             bool `json:"status"` // true 为禁言，false 未禁言
}

// UltraGroupGlobalBannedGet 查询超级群全体禁言
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/query-ultragroup-global-ban
func (rc *RongCloud) UltraGroupGlobalBannedGet(ctx context.Context, req *UltraGroupGlobalBannedGetRequest) (*UltraGroupGlobalBannedGetResponse, error) {
	path := "/ultragroup/globalbanned/get.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupGlobalBannedGetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupBannedWhitelistAddRequest struct {
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID 字段。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 用户 ID 列表，最多不超过 20 个。以逗号分割
	UserIds *string `url:"userIds,omitempty"`
}

type UltraGroupBannedWhitelistAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupBannedWhitelistAdd 加入超级群全体禁言白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/add-to-global-ban-whitelist
func (rc *RongCloud) UltraGroupBannedWhitelistAdd(ctx context.Context, req *UltraGroupBannedWhitelistAddRequest) (*UltraGroupBannedWhitelistAddResponse, error) {
	path := "/ultragroup/banned/whitelist/add.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupBannedWhitelistAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupBannedWhitelistDelRequest struct {
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID 字段。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 用户 ID 列表，最多不超过 20 个。以逗号分割
	UserIds *string `url:"userIds,omitempty"`
}

type UltraGroupBannedWhitelistDelResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupBannedWhitelistDel 移出超级群全体禁言白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/remove-from-global-ban-whitelist
func (rc *RongCloud) UltraGroupBannedWhitelistDel(ctx context.Context, req *UltraGroupBannedWhitelistDelRequest) (*UltraGroupBannedWhitelistDelResponse, error) {
	path := "/ultragroup/banned/whitelist/del.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupBannedWhitelistDelResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupBannedWhitelistGetRequest struct {
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// 频道 ID 字段。
	BusChannel *string `url:"busChannel,omitempty"`
	// 当前页码,默认1。
	Page *int `url:"page,omitempty"`
	// 每页条数，默认50。 上限 200
	PageSize *int `url:"pageSize,omitempty"`
}

type UltraGroupBannedWhitelistGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []*UltraGroupBannedWhitelistUser `json:"users"` // 用户列表
}

type UltraGroupBannedWhitelistUser struct {
	Id string `json:"id"` // 用户 ID
}

// UltraGroupBannedWhitelistGet 查询超级群全体禁言白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/query-global-ban-whitelist
func (rc *RongCloud) UltraGroupBannedWhitelistGet(ctx context.Context, req *UltraGroupBannedWhitelistGetRequest) (*UltraGroupBannedWhitelistGetResponse, error) {
	path := "/ultragroup/banned/whitelist/get.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupBannedWhitelistGetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserGroupAddRequest struct {
	// [必传] 超级群 ID，请确保超级群 ID 存在。
	GroupId *string `json:"groupId,omitempty"`
	// [必传] 用户组信息列表，最大长度为 10。
	UserGroups []*UltraGroupUserGroup `json:"userGroups,omitempty"`
}

type UltraGroupUserGroup struct {
	// [必传] 用户组 ID，支持大小写字母、数字的组合方式，长度不超过 64 个字符。
	UserGroupId string `json:"userGroupId"`
}

type UltraGroupUserGroupAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupUserGroupAdd 超级群创建用户组
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/create-usergroup
func (rc *RongCloud) UltraGroupUserGroupAdd(ctx context.Context, req *UltraGroupUserGroupAddRequest) (*UltraGroupUserGroupAddResponse, error) {
	path := "/ultragroup/usergroup/add.json"
	resp := &UltraGroupUserGroupAddResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserGroupDelRequest struct {
	// [必传] 超级群 ID，请确保超级群 ID 存在。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 用户组 ID 列表，多个 ID 以逗号分隔。单次请求最大长度为 10 个，否则全部失败。
	UserGroupIds *string `url:"userGroupIds,omitempty"`
}

type UltraGroupUserGroupDelResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupUserGroupDel 超级群删除用户组
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/delete-usergroup
func (rc *RongCloud) UltraGroupUserGroupDel(ctx context.Context, req *UltraGroupUserGroupDelRequest) (*UltraGroupUserGroupDelResponse, error) {
	path := "/ultragroup/usergroup/del.json"
	resp := &UltraGroupUserGroupDelResponse{}
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserGroupQueryRequest struct {
	// [必传] 超级群 ID，请确保超级群 ID 存在。
	GroupId *string `url:"groupId,omitempty"`
	// 查询页码，默认 1。
	Page *int `url:"page,omitempty"`
	// 每页条数，默认 10，最多为 50。
	PageSize *int `url:"pageSize,omitempty"`
}

type UltraGroupUserGroupQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	UserGroups         []*UltraGroupUserGroup `json:"userGroups"` // 用户组列表
}

// UltraGroupUserGroupQuery 超级群查询用户组列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/query-usergroup
func (rc *RongCloud) UltraGroupUserGroupQuery(ctx context.Context, req *UltraGroupUserGroupQueryRequest) (*UltraGroupUserGroupQueryResponse, error) {
	path := "/ultragroup/usergroup/query.json"
	resp := &UltraGroupUserGroupQueryResponse{}
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserGroupUserAddRequest struct {
	// [必传] 超级群 ID，请确保超级群 ID 存在。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 用户组 ID，请确保用户组 ID 存在。
	UserGroupId *string `url:"userGroupId,omitempty"`
	// [必传] 群内用户 ID 列表，多个 ID 以逗号分隔。单次不得超过 20人，请确保用户 ID 存在，否则全部失败。
	UserIds *string `url:"userIds,omitempty"`
}

type UltraGroupUserGroupUserAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupUserGroupUserAdd 超级群用户组添加用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/add-user-to-usergroup
func (rc *RongCloud) UltraGroupUserGroupUserAdd(ctx context.Context, req *UltraGroupUserGroupUserAddRequest) (*UltraGroupUserGroupUserAddResponse, error) {
	path := "/ultragroup/usergroup/user/add.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupUserGroupUserAddResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserGroupUserDelRequest struct {
	// [必传] 超级群 ID，请确保超级群 ID 存在。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 用户组 ID，请确保用户组 ID 存在。
	UserGroupId *string `url:"userGroupId,omitempty"`
	// [必传] 群内用户 ID 列表，多个 ID 以逗号分隔。单次不得超过 20人，请确保用户 ID 存在，否则全部失败。
	UserIds *string `url:"userIds,omitempty"`
}

type UltraGroupUserGroupUserDelResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UltraGroupUserGroupUserDel 超级群用户组移出用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/remove-user-from-usergroup
func (rc *RongCloud) UltraGroupUserGroupUserDel(ctx context.Context, req *UltraGroupUserGroupUserDelRequest) (*UltraGroupUserGroupUserDelResponse, error) {
	path := "/ultragroup/usergroup/user/del.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupUserGroupUserDelResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UltraGroupUserUserGroupQueryRequest struct {
	// [必传] 超级群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 用户 ID。
	UserId *string `url:"userId,omitempty"`
	// 查询页码，默认 1。
	Page *int `url:"page,omitempty"`
	// 每页条数，默认 10，最多为 50。
	PageSize *int `url:"pageSize,omitempty"`
}

type UltraGroupUserUserGroupQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Data               []string `json:"data"` // 用户组 ID 列表
}

// UltraGroupUserUserGroupQuery 超级群查询用户所属用户组
// More details see https://doc.rongcloud.cn/imserver/server/v1/ultragroup/query-usergroup-by-user
func (rc *RongCloud) UltraGroupUserUserGroupQuery(ctx context.Context, req *UltraGroupUserUserGroupQueryRequest) (*UltraGroupUserUserGroupQueryResponse, error) {
	path := "/ultragroup/user/usergroup/query.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &UltraGroupUserUserGroupQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
