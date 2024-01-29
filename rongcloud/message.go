package rongcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// MessagePrivatePublishRequest 单聊消息
type MessagePrivatePublishRequest struct {
	// 发送人用户 ID。
	//注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserId *string `json:"fromUserId"`

	// 接收用户 ID，可以实现向多人发送消息，每次上限为 1000 人。
	ToUserId *string `json:"toUserId"`

	RCMsg    RCMsg   `json:"-"`        // 消息类型参数的封装, 例如: TXTMsg(文本消息), HQVCMsg(高清语音消息)
	PushData *string `json:"pushData"` // iOS 平台收到推送消息时，可从 payload 中获取 APNs 推送数据，对应字段名为 appData（提示：rc 字段中默认携带了消息基本信息）。Android 平台收到推送消息时对应字段名为 appData。

	//指定收件人离线时触发的远程推送通知中的通知内容。注意：对于部分消息类型，该字段是否有值决定了是否触发远程推送通知。
	//如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中的用户内容类消息格式，可不填写该字段，远程推送通知默认使用服务端预置的推送通知内容。
	//如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中通知类、信令类（"撤回命令消息" 除外），且需要支持远程推送通知，则必须填写 pushContent，否则收件人不在线时无法收到远程推送通知。如无需触发远程推送，可不填该字段。
	//如果消息类型为自定义消息类型，且需要支持远程推送通知，则必须填写 pushContent 字段，否则收件人不在线时无法收到远程推送通知。
	//如果消息类型为自定义消息类型，但不需要支持远程推送通知（例如通过自定义消息类型实现的 App 业务层操作指令），可将 pushContent 字段留空禁用远程推送通知
	PushContent      *string `json:"pushContent"`
	IsIncludeSender  *int    `json:"isIncludeSender"`  // 是否向发件人客户端同步已发消息。1 表示同步，默认值为 0，即不同步。
	Count            *int    `json:"count"`            // 仅目标用户为 iOS 设备有效，Push 时用来控制桌面角标未读消息数，只有在 toUserId 为一个用户 ID 时有效，客户端获取远程推送内容时为 badge。具体参见 iOS 客户端文档「APNs 推送开发指南」目录下的集成 APNs 远程推送。为 -1 时不改变角标数，传入相应数字表示把角标数改为指定的数字，最大不超过 9999。
	VerifyBlacklist  *int    `json:"verifyBlacklist"`  // 是否过滤接收用户黑名单列表，0 表示为不过滤、 1 表示为过滤，默认为 0。
	IsPersisted      *int    `json:"isPersisted"`      // 是否需要为收件人在历史消息云端存储服务中存储此条消息。0 表示不存储；1 表示存储。默认值为 1，存储（依赖单群聊消息云端存储服务）。
	ContentAvailable *int    `json:"contentAvailable"` // 仅目标用户为 iOS 设备时有效，应用处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看知识库文档。1 表示为开启，0 表示为关闭，默认为 0。
	Expansion        *bool   `json:"expansion"`        // 是否为可扩展消息，默认为 false，设为 true 时终端在收到该条消息后，可对该条消息设置扩展信息。移动端 SDK 4.0.3 版本、Web 端 3.0.7 版本支持此功能。

	// 仅在 expansion 为 true 时有效。
	//自定义的消息扩展信息，该字段接受 JSON 字符串格式的键值对（key-value pairs）。请注意区别于消息体内的 extra 字段，extraContent 的值在消息发送后可修改，修改方式请参见服务端 API 接口文档消息扩展，或参考各客户端「消息扩展」接口文档。
	//KV 详细要求：以 Key、Value 的方式进行设置，如：{"type":"3"}。Key 最大 32 个字符，支持大小写英文字母、数字、 特殊字符+ = - _ 的组合方式，不支持汉字。Value 最大 4096 个字符。单次可设置最多 100 对 KV 扩展信息，单条消息最多可设置 300 对 KV 扩展信息。
	ExtraContent map[string]string `json:"extraContent"`

	// 是否为静默消息，默认为 false，设为 true 时终端用户离线情况下不会收到通知提醒。
	DisablePush *bool `json:"disablePush"`

	// 配置消息的推送通知，如推送通知的标题等。disablePush 属性为 true 时此属性无效。
	PushExt *string `json:"pushExt"`
}

func (r *MessagePrivatePublishRequest) MakeFormData() (url.Values, error) {
	res := url.Values{}
	if r.FromUserId != nil {
		res.Set("fromUserId", StringValue(r.FromUserId))
	}
	if r.ToUserId != nil {
		res.Set("toUserId", StringValue(r.ToUserId))
	}
	if r.RCMsg != nil {
		content, err := json.Marshal(r.RCMsg)
		if err != nil {
			return nil, err
		}
		res.Set("objectName", r.RCMsg.ObjectName())
		res.Set("content", string(content))
	}
	if r.PushData != nil {
		res.Set("pushData", StringValue(r.PushData))
	}
	if r.PushContent != nil {
		res.Set("pushContent", StringValue(r.PushContent))
	}
	if r.IsIncludeSender != nil {
		res.Set("isIncludeSender", strconv.Itoa(IntValue(r.IsIncludeSender)))
	}
	if r.Count != nil {
		res.Set("count", strconv.Itoa(IntValue(r.Count)))
	}
	if r.VerifyBlacklist != nil {
		res.Set("verifyBlacklist", strconv.Itoa(IntValue(r.VerifyBlacklist)))
	}
	if r.IsPersisted != nil {
		res.Set("isPersisted", strconv.Itoa(IntValue(r.IsPersisted)))
	}
	if r.ContentAvailable != nil {
		res.Set("contentAvailable", strconv.Itoa(IntValue(r.ContentAvailable)))
	}
	if r.Expansion != nil {
		res.Set("expansion", strconv.FormatBool(BoolValue(r.Expansion)))
	}
	if r.ExtraContent != nil {
		b, err := json.Marshal(r.ExtraContent)
		if err != nil {
			return nil, fmt.Errorf("Marshal ExtraContent err: %w", err)
		}
		res.Set("extraContent", string(b))
	}
	if r.DisablePush != nil {
		res.Set("disablePush", strconv.FormatBool(BoolValue(r.DisablePush)))
	}
	if r.PushExt != nil {
		res.Set("pushExt", StringValue(r.PushExt))
	}
	return res, nil
}

type MessagePrivatePublishResponse struct {
	Code int `json:"code"` // 返回码，200 为正常。

	httpResponseGetter `json:"-"`
}

// MessagePrivatePublish 发送单聊普通消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/send-private
func (rc *RongCloud) MessagePrivatePublish(ctx context.Context, req *MessagePrivatePublishRequest) (*MessagePrivatePublishResponse, error) {
	var resp *MessagePrivatePublishResponse
	data, err := req.MakeFormData()
	if err != nil {
		return nil, err
	}
	httpResp, err := rc.postFormUrlencoded(ctx, "/message/private/publish.json", data, &resp)
	if err != nil {
		return nil, err
	}
	resp.httpResponseGetter = &rawHttpResponseGetter{rawHttpResponseInternal: httpResp}
	return resp, nil
}
