package rongcloud

import (
	"context"
	"encoding/json"
	"fmt"
)

// TODO ultraGroupMsgModify?

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
	FromUserId   string            `json:"fromUserId"`   // 发送人用户 ID。
	GroupId      string            `json:"groupId"`      // 超级群 ID。
	SentTime     int64             `json:"sentTime"`     // 消息发送时间。Unix 时间戳，单位为毫秒。
	BusChannel   string            `json:"busChannel"`   // 频道 ID。
	MsgUID       string            `json:"msgUID"`       // 全局唯一消息 ID，即消息 UID。
	ObjectName   string            `json:"objectName"`   // 消息类型的唯一标识。
	Content      string            `json:"content"`      // 消息的内容。
	Expansion    string            `json:"expansion"`    // 是否为扩展消息。
	ExtraContent map[string]string `json:"extraContent"` // 消息扩展的内容，JSON 结构的 Key、Value 对，如：{"type":"3"}。Key 最大 32 个字符，支持大小写英文字母、数字、 特殊字符+ = - _ 的组合方式，不支持汉字。Value 最大 4096 个字符。
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

func (rc *RongCloud) UltraGroupMsgSend(ctx context.Context) {

}

type UltraGroupCreateRequest struct {
	UserId    *string `json:"userId,omitempty" url:"userId,omitempty"`    // [必传] 需要加入的用户 ID，创建后同时加入超级群。仅支持传入一个用户 ID。
	GroupId   *string `json:"groupId,omitempty" url:"userId,omitempty"`   //  [必传] 超级群 ID，最大长度 64 个字符。支持大小写英文字母与数字的组合。
	GroupName *string `json:"groupName,omitempty" url:"userId,omitempty"` // [必传] 超级群 ID 对应的名称，用于在发送群组消息显示远程 Push 通知时使用，如超级群名称改变需要调用刷新超级群信息接口同步。
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
