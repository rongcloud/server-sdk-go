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

	RCMsg    RCMsg   `json:"-"`        // 消息类型参数的SDK封装, 例如: TXTMsg(文本消息), HQVCMsg(高清语音消息)
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
	ContentAvailable *int    `json:"contentAvailable"` // 仅目标用户为 iOS 设备时有效，应用处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看[知识库文档](https://help.rongcloud.cn/t/topic/855)。1 表示为开启，0 表示为关闭，默认为 0。
	Expansion        *bool   `json:"expansion"`        // 是否为可扩展消息，默认为 false，设为 true 时终端在收到该条消息后，可对该条消息设置扩展信息。移动端 SDK 4.0.3 版本、Web 端 3.0.7 版本支持此功能。

	// 仅在 expansion 为 true 时有效。
	//自定义的消息扩展信息，该字段接受 JSON 字符串格式的键值对（key-value pairs）。请注意区别于消息体内的 extra 字段，extraContent 的值在消息发送后可修改，修改方式请参见服务端 API 接口文档消息扩展，或参考各客户端「消息扩展」接口文档。
	//KV 详细要求：以 Key、Value 的方式进行设置，如：{"type":"3"}。Key 最大 32 个字符，支持大小写英文字母、数字、 特殊字符+ = - _ 的组合方式，不支持汉字。Value 最大 4096 个字符。单次可设置最多 100 对 KV 扩展信息，单条消息最多可设置 300 对 KV 扩展信息。
	ExtraContent map[string]string `json:"extraContent"`

	// 是否为静默消息，默认为 false，设为 true 时终端用户离线情况下不会收到通知提醒。
	DisablePush *bool `json:"disablePush"`

	// 配置消息的推送通知，如推送通知的标题等。disablePush 属性为 true 时此属性无效。json string
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
		err := makeRCMsgUrlValues(r.RCMsg, res)
		if err != nil {
			return nil, err
		}
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

const (
	ConversationTypePrivate = "1" // 二人会话
	ConversationTypeGroup   = "3" // 群组会话
)

type MessageExpansionSetRequest struct {
	// [必传] 消息唯一标识 ID，可通过全量消息路由功能获取。详见[全量消息路由](https://doc.rongcloud.cn/imserver/server/v1/message/sync)。
	MsgUID *string `url:"msgUID,omitempty"`
	// [必传] 操作者用户 ID，即需要为指定消息（msgUID）设置扩展信息的用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 会话类型。支持的会话类型包括："1" ConversationTypePrivate（二人会话）、"3" ConversationTypeGroup（群组会话）。
	ConversationType *string `url:"conversationType,omitempty"`
	// [必传] 目标 ID，根据不同的 conversationType，可能是用户 ID 或群组 ID。
	TargetId *string `url:"targetId,omitempty"`
	// [必传] 消息扩展的内容，JSON 结构，以 Key、Value 的方式进行设置，如：{"type":"3"}。Key 最大 32 个字符，支持大小写英文字母、数字、 特殊字符+ = - _ 的组合方式，不支持汉字。Value 最大 4096 个字符。单条消息可设置 300 个扩展信息，一次最多可以设置 100 个。
	ExtraKeyVal map[string]string `url:"-"`
	// 删除操作会生成一条「扩展操作消息」。该字段指定是否将该「扩展操作消息」同步到发件人（扩展操作者）的客户端。1 表示同步，默认值为 0，即不同步。注意，仅设置该参数无法确保发件人客户端一定能获取到该条已发消息，您可能还需要启用其他服务。详见[发件人客户端如何同步已发消息](https://doc.rongcloud.cn/imserver/server/v1/message/how-to-sync-to-sender-client)。
	IsSyncSender *int `url:"isSyncSender,omitempty"`
}

type MessageExpansionSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageExpansionSet 设置单群聊消息扩展
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/set-expansion
func (rc *RongCloud) MessageExpansionSet(ctx context.Context, req *MessageExpansionSetRequest) (*MessageExpansionSetResponse, error) {
	path := "/message/expansion/set.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	if req.ExtraKeyVal != nil {
		kv, err := json.Marshal(req.ExtraKeyVal)
		if err != nil {
			return nil, NewSDKError(fmt.Sprintf("marshal ExtraKeyVal error %s", err))
		}
		params.Set("extraKeyVal", string(kv))
	}

	resp := &MessageExpansionSetResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageExpansionDeleteRequest struct {
	// [必传] 消息唯一标识 ID，可通过全量消息路由功能获取。详见[全量消息路由](https://doc.rongcloud.cn/imserver/server/v1/message/sync)。
	MsgUID *string `url:"msgUID,omitempty"`
	// [必传] 操作者用户 ID，即需要为指定消息（msgUID）设置扩展信息的用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 会话类型。支持的会话类型包括："1" ConversationTypePrivate（二人会话）、"3" ConversationTypeGroup（群组会话）。
	ConversationType *string `url:"conversationType,omitempty"`
	// [必传] 目标 ID，根据不同的 conversationType，可能是用户 ID 或群组 ID。
	TargetId *string `url:"targetId,omitempty"`
	// [必传] 需要删除的扩展信息的 Key 值，一次最多可以删除 100 个扩展信息。
	ExtraKey []string `url:"-"`
	// 删除操作会生成一条「扩展操作消息」。该字段指定是否将该「扩展操作消息」同步到发件人（扩展操作者）的客户端。1 表示同步，默认值为 0，即不同步。注意，仅设置该参数无法确保发件人客户端一定能获取到该条已发消息，您可能还需要启用其他服务。详见[发件人客户端如何同步已发消息](https://doc.rongcloud.cn/imserver/server/v1/message/how-to-sync-to-sender-client)。
	IsSyncSender *int `url:"isSyncSender,omitempty"`
}

type MessageExpansionDeleteResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageExpansionDelete 删除单群聊消息扩展
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/delete-expansion
func (rc *RongCloud) MessageExpansionDelete(ctx context.Context, req *MessageExpansionDeleteRequest) (*MessageExpansionDeleteResponse, error) {
	path := "/message/expansion/delete.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	if req.ExtraKey != nil {
		extraKeys, err := json.Marshal(req.ExtraKey)
		if err != nil {
			return nil, NewSDKError(fmt.Sprintf("marshal extraKey error %s", err))
		}
		params.Set("extraKey", string(extraKeys))
	}
	resp := &MessageExpansionDeleteResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageExpansionQueryRequest struct {
	// [必传] 消息唯一标识 ID，可通过全量消息路由功能获取。详见[全量消息路由](https://doc.rongcloud.cn/imserver/server/v1/message/sync)
	MsgUID *string `url:"msgUID,omitempty"`
	// 页数，默认返回 300 个扩展信息。
	PageNo *int `url:"pageNo,omitempty"`
}

type MessageExpansionQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	// 消息扩展的内容
	ExtraContent map[string]*MessageExpansionQueryExtraContentValue `json:"extraContent"`
}

type MessageExpansionQueryExtraContentValue struct {
	Value     string `json:"v"`  // value
	Timestamp int64  `json:"ts"` // 版本号
}

// MessageExpansionQuery 获取单群聊消息扩展
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/get-expansion
func (rc *RongCloud) MessageExpansionQuery(ctx context.Context, req *MessageExpansionQueryRequest) (*MessageExpansionQueryResponse, error) {
	path := "/message/expansion/query.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &MessageExpansionQueryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageUltraGroupPublishRequest struct {
	// [必传] 发送人用户 ID，通过 Server API 非群成员也可以向群组中发送消息。 注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserId *string `json:"fromUserId,omitempty"`
	// [必传] 接收群 ID，提供多个本参数可以实现向多群发送消息，最多不超过 3 个超级群。
	ToGroupIds []string `json:"toGroupIds,omitempty"`
	// [必传] 消息类型，接受内置消息类型（见消息类型概述）或自定义消息的消息类型值。
	// 注意：在自定义消息时，消息类型不可以 "RC:" 开头，以免与系统内置消息类型重名；消息类型长度不可超过 32 个字符。SDK 中必须已注册过该自定义消息，否则 SDK 收到该消息后将无法解析。
	ObjectName *string `json:"objectName,omitempty"`
	// [必传] 所发送消息的内容，单条消息最大 128k。
	//内置消息类型：将消息内容体 JSON 对象序列化为 JSON 字符串传入。消息内容 JSON 结构体详见用户内容类消息格式或其他内置消息类型的消息内容格式。
	//例如，文本消息内容 JSON 结构体内部包含 content 字段（此为 JSON 结构体内的 key 值，注意区分），则需要将 {"content":"Hello world!"} 序列化后的结果作为此处 content 字段的值。
	//自定义消息类型（objectName 字段必须指定为自定义消息类型）：如果发送自定义消息，该参数可自定义格式，不限于 JSON。
	Content *string `json:"content,omitempty"`
	// 指定收件人离线时触发的远程推送通知中的通知内容。注意：对于部分消息类型，该字段是否有值决定了是否触发远程推送通知。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中的用户内容类消息格式，可不填写该字段，远程推送通知默认使用服务端预置的推送通知内容。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中通知类、信令类（"撤回命令消息" 除外），且需要支持远程推送通知，则必须填写 pushContent，否则收件人不在线时无法收到远程推送通知。如无需触发远程推送，可不填该字段。
	// 如果消息类型为自定义消息类型，且需要支持远程推送通知，则必须填写 pushContent 字段，否则收件人不在线时无法收到远程推送通知。
	// 如果消息类型为自定义消息类型，但不需要支持远程推送通知（例如通过自定义消息类型实现的 App 业务层操作指令），可将 pushContent 字段留空禁用远程推送通知
	PushContent *string `json:"pushContent,omitempty" `
	// iOS 平台收到推送消息时，可从 payload 中获取 APNs 推送数据，对应字段名为 appData（提示：rc 字段中默认携带了消息基本信息）。Android 平台收到推送消息时对应字段名为 appData。
	PushData *string `json:"pushData,omitempty" `
	// 是否需要为收件人在历史消息云端存储服务中存储此条消息。0 表示不存储；1 表示存储。默认值为 1，存储（依赖单群聊消息云端存储服务）。
	IsPersisted *int `json:"isPersisted,omitempty" `
	// 用户未在线时是否计入未读消息数。0 表示为不计数、1 表示为计数，默认为 1。
	IsCounted *int `json:"isCounted,omitempty" `
	// 是否为 @ 消息，不传时默认为非 @ 消息（效果等于传 0）。如果需要发送 @ 消息，必须指定为 1，且必须在消息内容字段（content）内部携带 @ 相关信息（mentionedInfo，可参考下方请求示例）。关于 mentionedInfo 结构的详细说明，参见如何发送 @ 消息。
	IsMentioned *int `json:"isMentioned,omitempty" `
	// 仅目标用户为 iOS 设备时有效，应用处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看[知识库文档](https://help.rongcloud.cn/t/topic/855)。1 表示为开启，0 表示为关闭，默认为 0。
	ContentAvailable *int `json:"contentAvailable,omitempty" `
	// 配置消息的推送通知，如推送通知的标题等。disablePush 属性为 true 时此属性无效。json string
	PushExt *string `json:"pushExt,omitempty" `
	// 频道 Id，发消息时会对群 ID 下的频道 ID 做合法性校验，如果群 ID 下无此频道 ID 则服务端不会下发此条消息。支持大小写英文字母、数字的组合方式，不支持特殊字符，最长为 20 个字符。
	BusChannel *string `json:"busChannel,omitempty" `
	// 是否为可扩展消息，默认为 false，设为 true 时终端在收到该条消息后，可对该条消息设置扩展信息
	Expansion *bool `json:"expansion,omitempty" `
	// 仅在 expansion 为 true 时有效。
	//自定义的消息扩展信息，该字段接受 JSON 字符串格式的键值对（key-value pairs）。请注意区别于消息体内的 extra 字段，extraContent 的值在消息发送后可修改，修改方式请参见服务端 API 接口文档消息扩展，或参考各客户端「消息扩展」接口文档。
	//KV 详细要求：以 Key、Value 的方式进行设置，如：{"type":"3"}。Key 最大 32 个字符，支持大小写英文字母、数字、 特殊字符+ = - _ 的组合方式，不支持汉字。Value 最大 4096 个字符。单次可设置最多 100 对 KV 扩展信息，单条消息最多可设置 300 对 KV 扩展信息。
	ExtraContent *string `json:"extraContent,omitempty" `
}

type MessageUltraGroupPublishResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageUltraGroupPublish 发送超级群消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/send-ultragroup
func (rc *RongCloud) MessageUltraGroupPublish(ctx context.Context, req *MessageUltraGroupPublishRequest) (*MessageUltraGroupPublishResponse, error) {
	path := "/message/ultragroup/publish.json"
	resp := &MessageUltraGroupPublishResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageRecallRequest struct {
	// [必传] 消息发送人用户 ID。
	FromUserId *string `url:"fromUserId,omitempty"`
	// [必传] 会话类型。支持的会话类型包括：1（二人会话）、3（群组会话）、4（聊天室会话）、6（系统会话）、10（超级群会话）。
	ConversationType *int `url:"conversationType,omitempty"`
	// [必传] 目标 ID，根据不同的会话类型（ConversationType），可能是用户 ID、群组 ID、聊天室 ID、超级群 ID，系统目标 ID。
	TargetId *string `url:"targetId,omitempty"`
	// 超级群频道 ID，仅适用于撤回超级群消息。使用要求如下：
	// 如果发送消息时指定了频道 ID，则撤回时必须指定频道 ID，否则无法撤回。
	// 如果发送消息时未指定频道 ID，则撤回时不可指定频道 ID，否则无法撤回。
	// 客户端发送超级群消息时，频道 ID 对应字段名称为 channelId。
	BusChannel *string `url:"busChannel,omitempty"`
	// [必传] 消息唯一标识。
	// 可通过全量消息路由服务获取消息唯一标识，对应名称为 msgUID。
	// 全量消息路由服务不支持的消息，目前只能通过历史消息日志获取，对应字段名称为 msgUID。
	MessageUID *string `url:"messageUID,omitempty"`
	// [必传] 消息发送时间。
	// 可通过全量消息路由服务获取消息发送时间，对应名称为 msgTimestamp。
	// 全量消息路由服务不支持的消息，目前只能通过历史消息日志获取，对应字段名称为 dateTime。
	SentTime *int64 `url:"sentTime,omitempty"`
	// 是否为管理员，默认为 0，设为 1 时，IMKit 收到此条消息后，小灰条默认显示为“管理员 撤回了一条消息”。
	IsAdmin *int `url:"isAdmin,omitempty"`
	// 指定移动端接收方是否需要在本地删除原始消息记录及显示撤回消息提示，默认为 0。
	// 为 0 时，移动端接收方仅将原始消息内容替换为撤回提示（小灰条通知），不删除该原始消息记录。
	// 为 1 时，移动端接收方会删除原始消息记录，不显示撤回提示（小灰条通知）。
	// 注意：即时通讯服务端历史消息不保存已撤回的超级群消息记录。如果 isDelete 设置为 0，撤回超级群消息后，移动端本地会存有记录（显示为撤回提示），而 Web 端无记录，可能会造成用户体验差异。
	IsDelete *int `url:"isDelete,omitempty"`
	// 是否为静默撤回，默认为 false，设为 true 时终端用户离线情况下不会收到撤回通知提醒。该字段不支持聊天室、超级群会话类型。
	DisablePush *bool `url:"disablePush,omitempty"`
	// 扩展信息，可以放置任意的数据内容。不支持超级群会话（conversationType 为 10）。
	Extra *string `url:"extra,omitempty"`
}

type MessageRecallResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageRecall 撤回消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/msgrecall
func (rc *RongCloud) MessageRecall(ctx context.Context, req *MessageRecallRequest) (*MessageRecallResponse, error) {
	path := "/message/recall.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &MessageRecallResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageBroadcastRequest struct {
	// [必传] 发送人用户 ID。
	// 注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserId *string `url:"fromUserId,omitempty"`
	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息), HQVCMsg(高清语音消息)
	RCMsg RCMsg `url:"-"`
	// 指定收件人离线时触发的远程推送通知中的通知内容。注意：对于部分消息类型，该字段是否有值决定了是否触发远程推送通知。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中的用户内容类消息格式，可不填写该字段，远程推送通知默认使用服务端预置的推送通知内容。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中通知类、信令类（"撤回命令消息" 除外），且需要支持远程推送通知，则必须填写 pushContent，否则收件人不在线时无法收到远程推送通知。如无需触发远程推送，可不填该字段。
	// 如果消息类型为自定义消息类型，且需要支持远程推送通知，则必须填写 pushContent 字段，否则收件人不在线时无法收到远程推送通知。
	// 如果消息类型为自定义消息类型，但不需要支持远程推送通知（例如通过自定义消息类型实现的 App 业务层操作指令），可将 pushContent 字段留空禁用远程推送通知。
	PushContent *string `url:"pushContent,omitempty"`
	// iOS 平台收到推送消息时，可从 payload 中获取 APNs 推送数据，对应字段名为 appData（提示：rc 字段中默认携带了消息基本信息）。Android 平台收到推送消息时对应字段名为 appData。
	PushData *string `url:"pushData,omitempty"`
	// 针对 iOS 平台，对 SDK 处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。 允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看知识库文档。1 表示为开启，0 表示为关闭，默认为 0
	ContentAvailable *int `url:"contentAvailable,omitempty"`
	// 配置消息的推送通知，如推送通知的标题等。disablePush 属性为 true 时此属性无效。具体请查看下方 pushExt 参数说明。json string
	PushExt *string `url:"pushExt,omitempty"`
}

type MessageBroadcastResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageBroadcast 发送全量用户落地通知
// More details see https://doc.rongcloud.cn/imserver/server/v1/system/send-message-broadcast-to-all
func (rc *RongCloud) MessageBroadcast(ctx context.Context, req *MessageBroadcastRequest) (*MessageBroadcastResponse, error) {
	path := "/message/broadcast.json"
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
	resp := &MessageBroadcastResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type StatusMessagePrivatePublishRequest struct {
	// [必传] 发送人用户 ID。
	FromUserId *string `url:"fromUserId,omitempty"`
	// [必传] 接收用户 ID，支持向多人发送消息，每次上限为 1000 人。
	ToUserIds []string `url:"toUserId,omitempty"`
	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息), HQVCMsg(高清语音消息)
	RCMsg RCMsg `url:"-"`
	// 是否过滤发送人黑名单列表，0 表示为不过滤、 1 表示为过滤，默认为 0 不过滤。
	VerifyBlacklist *int `url:"verifyBlacklist,omitempty"`
	// 是否向发件人客户端同步已发消息。1 表示同步，默认值为 0，即不同步。注意，该接口用于发送状态消息，因此仅支持在发件人已登陆客户端（在线）的情况下同步已发消息。
	IsIncludeSender *int `url:"isIncludeSender,omitempty"`
}

type StatusMessagePrivatePublishResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// StatusMessagePrivatePublish 发送单聊状态消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/send-status-private
func (rc *RongCloud) StatusMessagePrivatePublish(ctx context.Context, req *StatusMessagePrivatePublishRequest) (*StatusMessagePrivatePublishResponse, error) {
	path := "/statusmessage/private/publish.json"
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
	resp := &StatusMessagePrivatePublishResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessagePrivatePublishTemplateRequest struct {
	// [必传] 发送人用户 ID。
	// 注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserId *string `json:"fromUserId"`
	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息)
	RCMsg RCMsg `json:"-"`
	// [必传] SDK封装消息模板
	MessageTemplate []*MessageTemplate `json:"-"`
	// 是否向发件人客户端同步已发消息。1 表示同步，默认值为 0，即不同步。注意，仅设置该参数无法确保发件人客户端一定能获取到该条已发消息，您可能还需要启用其他服务。详见发件人客户端如何同步已发消息。
	IsIncludeSender *int `json:"isIncludeSender"`
	// 是否过滤发送人黑名单列表，0 为不过滤、 1 为过滤，默认为 0 不过滤。
	VerifyBlacklist *int `json:"verifyBlacklist,omitempty"`
	// 是否需要为收件人在历史消息云端存储服务中存储此条消息。0 表示不存储；1 表示存储。默认值为 1，存储（依赖单群聊消息云端存储服务）。
	// 此属性不影响离线消息功能，用户未在线时都会转为离线消息?存储。
	// 提示：一般情况下（第 1、2 种情况），客户端是否存储消息不依赖此参数。以下第 3 种情况属于例外：
	// 如果消息属于内置消息类型，客户端 SDK 会根据消息类型本身的存储属性标识判断是否存入本地数据库。详见消息类型概述。
	// 如果消息属于自定义消息类型，则客户端 SDK 会根据该类型在客户端上注册时的存储属性标识判断是否需要存入本地数据库。
	// 如果消息属于客户端 App 上未注册自定义消息类型（例如客户端使用的 App 版本过旧），则客户端 SDK 会根据当前参数值确定是否将消息存储在本地。但因消息类型未注册，客户端无法解析显示该消息。
	IsPersisted *int `json:"isPersisted,omitempty"`
	// 仅目标用户为 iOS 设备时有效，对 SDK 处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。 允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看知识库文档。1 表示为开启，0 表示为关闭，默认为 0
	ContentAvailable *int `json:"contentAvailable,omitempty"`
	// 是否为可扩展消息，默认为 false，设为 true 时终端在收到该条消息后，可对该条消息设置扩展信息。
	Expansion *bool `json:"expansion,omitempty"`
	// 是否为静默消息，默认为 false，设为 true 时终端用户离线情况下不会收到通知提醒。
	DisablePush *bool `json:"disablePush,omitempty"`
	// 配置消息的推送通知，如推送通知的标题等。disablePush 属性为 true 时此属性无效。json string
	PushExt *string `json:"pushExt,omitempty"`
}

func (r *MessagePrivatePublishTemplateRequest) MarshalJSON() ([]byte, error) {
	req := messagePublishTemplateRequest{
		FromUserId:       r.FromUserId,
		IsIncludeSender:  r.IsIncludeSender,
		VerifyBlacklist:  r.VerifyBlacklist,
		IsPersisted:      r.IsPersisted,
		ContentAvailable: r.ContentAvailable,
		Expansion:        r.Expansion,
		DisablePush:      r.DisablePush,
		PushExt:          r.PushExt,
	}
	if r.RCMsg != nil {
		req.ObjectName = r.RCMsg.ObjectName()
		content, err := r.RCMsg.ToString()
		if err != nil {
			return nil, fmt.Errorf("%s RcMsg.ToString() error %s", req.ObjectName, err)
		}
		req.Content = content
	}
	if r.MessageTemplate != nil {
		var values []map[string]string
		var toUserIds []string
		var pushContent []string
		var pushData []string
		for _, template := range r.MessageTemplate {
			values = append(values, template.Value)
			toUserIds = append(toUserIds, template.ToUserId)
			pushContent = append(pushContent, template.PushContent)
			pushData = append(pushData, template.PushData)
		}
		req.Values = values
		req.ToUserId = toUserIds
		req.PushContent = pushContent
		req.PushData = pushData
	}
	return json.Marshal(req)
}

type messagePublishTemplateRequest struct {
	FromUserId       *string             `json:"fromUserId,omitempty"`
	ToUserId         []string            `json:"toUserId,omitempty"`
	ObjectName       string              `json:"objectName,omitempty"`
	Content          string              `json:"content,omitempty"`
	Values           []map[string]string `json:"values,omitempty"`
	PushContent      []string            `json:"pushContent,omitempty"`
	PushData         []string            `json:"pushData,omitempty"`
	IsIncludeSender  *int                `json:"isIncludeSender,omitempty"`
	VerifyBlacklist  *int                `json:"verifyBlacklist,omitempty"`
	IsPersisted      *int                `json:"isPersisted,omitempty"`
	ContentAvailable *int                `json:"contentAvailable,omitempty"`
	Expansion        *bool               `json:"expansion,omitempty"`
	DisablePush      *bool               `json:"disablePush,omitempty"`
	PushExt          *string             `json:"pushExt,omitempty"`
}

type MessageTemplate struct {
	// [必传] 接收用户 ID。
	ToUserId string
	// [必传] 为消息内容（content）、推送通知内容（pushContent）、推送数据（pushData）中的标识位（标识位示例：{d}）提供对应的值。
	Value map[string]string
	// [必传] 指定收件人离线时触发的远程推送通知中的通知内容。注意：对于部分消息类型，该字段是否有值决定了是否触发远程推送通知。支持定义模板标识位，使用 values 中的值进行替换。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义的消息类型，填写该字段后，离线推送通知中显示模板定义的推送内容，而非消息类型的默认推送内容。
	// 如果消息类型为自定义消息类型，且需要支持远程推送通知，则必须填写 pushContent 字段，否则收件人不在线时无法收到远程推送通知。
	// 如果消息类型为自定义消息类型，但不需要支持远程推送通知（例如通过自定义消息类型实现的 App 业务层操作指令），可将 pushContent 字段对应数组传空值禁用离线推送。
	PushContent string
	// iOS 平台收到推送消息时，可从 payload 中获取 APNs 推送数据，对应字段名为 appData（提示：rc 字段中默认携带了消息基本信息）。Android 平台收到推送消息时对应字段名为 appData。
	PushData string
}

type MessagePrivatePublishTemplateResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessagePrivatePublishTemplate 发送单聊模板消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/send-private-template
func (rc *RongCloud) MessagePrivatePublishTemplate(ctx context.Context, req *MessagePrivatePublishTemplateRequest) (*MessagePrivatePublishTemplateResponse, error) {
	path := "/message/private/publish_template.json"
	resp := &MessagePrivatePublishTemplateResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageGroupPublishRequest struct {
	// [必传] 发送人用户 ID，通过 Server API 非群成员也可以向群组中发送消息。
	//注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserId *string `url:"fromUserId,omitempty"`

	// [必传] 接收消息的群 ID。支持最多 3 个群组 ID。发送群聊定向消息时，仅支持传入一个群组 ID。
	ToGroupId []string `url:"toGroupId,omitempty"`

	// 发送群聊定向消息时，接收消息的群成员用户 ID 列表，群中其他用户无法收到该定向消息。仅当 toGroupId 传入单个群组 ID 时有效。
	ToUserId []string `url:"toUserId,omitempty"`

	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息)
	RCMsg RCMsg `url:"-"`

	// 指定收件人离线时触发的远程推送通知中的通知内容。注意：对于部分消息类型，该字段是否有值决定了是否触发远程推送通知。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中的用户内容类消息格式，可不填写该字段，远程推送通知默认使用服务端预置的推送通知内容。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中通知类、信令类（"撤回命令消息" 除外），且需要支持远程推送通知，则必须填写 pushContent，否则收件人不在线时无法收到远程推送通知。如无需触发远程推送，可不填该字段。
	// 如果消息类型为自定义消息类型，且需要支持远程推送通知，则必须填写 pushContent 字段，否则收件人不在线时无法收到远程推送通知。
	// 如果消息类型为自定义消息类型，但不需要支持远程推送通知（例如通过自定义消息类型实现的 App 业务层操作指令），可将 pushContent 字段留空禁用远程推送通知。
	PushContent *string `url:"pushContent,omitempty"`

	// iOS 平台收到推送消息时，可从 payload 中获取 APNs 推送数据，对应字段名为 appData（提示：rc 字段中默认携带了消息基本信息）。Android 平台收到推送消息时对应字段名为 appData。
	PushData *string `url:"pushData,omitempty"`

	// 是否向发件人客户端同步已发消息。1 表示同步，默认值为 0，即不同步。注意，仅设置该参数无法确保发件人客户端一定能获取到该条已发消息，您可能还需要启用其他服务。详见发件人[客户端如何同步已发消息](https://doc.rongcloud.cn/imserver/server/v1/message/how-to-sync-to-sender-client)。
	IsIncludeSender *int `url:"isIncludeSender,omitempty"`

	// 是否需要为收件人在历史消息云端存储服务中存储此条消息。0 表示不存储；1 表示存储。默认值为 1，存储（依赖单群聊消息云端存储服务）。
	// 注意：即使已开通单群聊消息云存储功能，群组定向消息也不会存入服务端历史消息记录。如有需要，请提交工单申请开通群定向消息云存储服务。
	// 此属性不影响离线消息功能，用户未在线时都会转为离线消息?存储。
	// 提示：一般情况下（第 1、2 种情况），客户端是否存储消息不依赖此参数。以下第 3 种情况属于例外：
	// 如果消息属于内置消息类型，客户端 SDK 会根据消息类型本身的存储属性标识判断是否存入本地数据库。详见消息类型概述。
	// 如果消息属于自定义消息类型，则客户端 SDK 会根据该类型在客户端上注册时的存储属性标识判断是否需要存入本地数据库。
	// 如果消息属于客户端 App 上未注册自定义消息类型（例如客户端使用的 App 版本过旧），则客户端 SDK 会根据当前参数值确定是否将消息存储在本地。但因消息类型未注册，客户端无法解析显示该消息。
	IsPersisted *int `url:"isPersisted,omitempty"`

	// 是否为 @ 消息，不传时默认为非 @ 消息（效果等于传 0）。如果需要发送 @ 消息，必须指定为 1，且必须在消息内容字段（content）内部携带 @ 相关信息（mentionedInfo，可参考下方请求示例）。关于 mentionedInfo 结构的详细说明，参见如何发送 @ 消息。
	IsMentioned *int `url:"isMentioned,omitempty"`

	// 仅目标用户为 iOS 设备时有效，应用处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看[知识库文档](https://help.rongcloud.cn/t/topic/855)。1 表示为开启，0 表示为关闭，默认为 0。
	ContentAvailable *int `url:"contentAvailable,omitempty"`

	// 是否为可扩展消息，默认为 false，设为 true 时终端在收到该条消息后，可对该条消息设置扩展信息。移动端 SDK 4.0.3 版本、Web 端 3.0.7 版本支持此功能。仅当 toGroupId 传入单个群组 ID 时有效。
	Expansion *bool `url:"expansion,omitempty"`

	// 仅在 expansion 为 true 时有效。
	//自定义的消息扩展信息，该字段接受 JSON 字符串格式的键值对（key-value pairs）。请注意区别于消息体内的 extra 字段，extraContent 的值在消息发送后可修改，修改方式请参见服务端 API 接口文档消息扩展，或参考各客户端「消息扩展」接口文档。
	//KV 详细要求：以 Key、Value 的方式进行设置，如：{"type":"3"}。Key 最大 32 个字符，支持大小写英文字母、数字、 特殊字符+ = - _ 的组合方式，不支持汉字。Value 最大 4096 个字符。单次可设置最多 100 对 KV 扩展信息，单条消息最多可设置 300 对 KV 扩展信息。
	ExtraContent map[string]string `url:"-"`

	// 是否为静默消息，默认为 false，设为 true 时终端用户离线情况下不会收到通知提醒。仅当 toGroupId 传入单个群组 ID 时有效。
	DisablePush *bool `url:"disablePush,omitempty"`

	// 配置消息的推送通知，如推送通知的标题等。disablePush 属性为 true 时此属性无效。仅当 toGroupId 传入单个群组 ID 时有效。json string
	PushExt *string `url:"pushExt,omitempty"`
}

type MessageGroupPublishResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageGroupPublish 发送群聊消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/send-group
func (rc *RongCloud) MessageGroupPublish(ctx context.Context, req *MessageGroupPublishRequest) (*MessageGroupPublishResponse, error) {
	path := "/message/group/publish.json"
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
	if req.ExtraContent != nil {
		extraContent, err := json.Marshal(req.ExtraContent)
		if err != nil {
			return nil, NewSDKError(fmt.Sprintf("extraContent marshal error %s", err))
		}
		params.Set("extraContent", string(extraContent))
	}
	resp := &MessageGroupPublishResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type StatusMessageGroupPublishRequest struct {
	// [必传] 发送人用户 ID，通过 Server API 非群成员也可以向群组中发送消息。
	FromUserId *string `url:"fromUserId,omitempty"`

	// [必传] 接收群ID，提供多个本参数可以实现向多群发送消息，最多不超过 3 个群组。
	ToGroupId []string `url:"toGroupId,omitempty"`

	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息)
	RCMsg RCMsg `url:"-"`

	// 是否向发件人客户端同步已发消息。1 表示同步，默认值为 0，即不同步。注意，该接口用于发送状态消息，因此仅支持在发件人已登陆客户端（在线）的情况下同步已发消息。
	IsIncludeSender *int `url:"isIncludeSender,omitempty"`
}

type StatusMessageGroupPublishResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// StatusMessageGroupPublish 发送群聊状态消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/send-status-group
func (rc *RongCloud) StatusMessageGroupPublish(ctx context.Context, req *StatusMessageGroupPublishRequest) (*StatusMessageGroupPublishResponse, error) {
	path := "/statusmessage/group/publish.json"
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
	resp := &StatusMessageGroupPublishResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageChatroomPublishRequest struct {
	// [必传] 发送人用户 ID。
	FromUserId *string `url:"fromUserId,omitempty"`

	// [必传] 接收聊天室 ID，提供多个本参数可以实现向多个聊天室发送消息，建议最多不超过 10 个聊天室。
	ToChatroomId []string `url:"toChatroomId,omitempty"`

	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息)
	RCMsg RCMsg `url:"-"`

	// 是否需要为收件人在历史消息云端存储服务中存储此条消息。0 表示不存储；1 表示存储。默认值为 1，存储（依赖聊天室消息云端存储服务）。
	// 一般情况下（第 1、2 种情况），客户端是否存储消息不依赖此参数。以下第 3 种情况属于例外：
	// 如果消息属于内置消息类型，客户端 SDK 会根据消息类型本身的存储属性标识判断是否存入本地数据库。详见消息类型概述。
	// 如果消息属于自定义消息类型，则客户端 SDK 会根据该类型在客户端上注册时的存储属性标识判断是否需要存入本地数据库。
	// 如果消息属于客户端 App 上未注册自定义消息类型（例如客户端使用的 App 版本过旧），则客户端 SDK 会根据当前参数值确定是否将消息存储在本地。但因消息类型未注册，客户端无法解析显示该消息。
	// 注意：客户端会在用户退出聊天室时自动清除本地的聊天室历史消息记录。
	IsPersisted *int `url:"isPersisted,omitempty"`

	// 是否向发件人客户端同步已发消息。1 表示同步，默认值为 0，即不同步。注意，仅设置该参数无法确保发件人客户端一定能获取到该条已发消息，您可能还需要启用其他服务。详见发件人客户端如何同步已发消息。
	IsIncludeSender *int `url:"isIncludeSender,omitempty"`
}

type MessageChatroomPublishResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageChatroomPublish 发送聊天室消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/send-chatroom
func (rc *RongCloud) MessageChatroomPublish(ctx context.Context, req *MessageChatroomPublishRequest) (*MessageChatroomPublishResponse, error) {
	path := "/message/chatroom/publish.json"
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
	resp := &MessageChatroomPublishResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageChatroomBroadcastRequest struct {
	// [必传] 发送人用户 ID。
	FromUserId *string `url:"fromUserId,omitempty"`

	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息)
	RCMsg RCMsg `url:"-"`

	// 是否向发件人客户端同步已发消息。1 表示同步，默认值为 0，即不同步。注意，仅设置该参数无法确保发件人客户端一定能获取到该条已发消息，您可能还需要启用其他服务。详见发件人客户端如何同步已发消息。
	IsIncludeSender *int `url:"isIncludeSender,omitempty"`
}

type MessageChatroomBroadcastResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageChatroomBroadcast 发送全体聊天室广播消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/send-chatroom-broadcast
func (rc *RongCloud) MessageChatroomBroadcast(ctx context.Context, req *MessageChatroomBroadcastRequest) (*MessageChatroomBroadcastResponse, error) {
	path := "/message/chatroom/broadcast.json"
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
	resp := &MessageChatroomBroadcastResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageOnlineBroadcastRequest struct {
	// [必传] 发送人用户 ID。
	FromUserId *string `url:"fromUserId,omitempty"`

	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息)
	RCMsg RCMsg `url:"-"`
}

type MessageOnlineBroadcastResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageOnlineBroadcast 发送在线用户广播
// More details see https://doc.rongcloud.cn/imserver/server/v1/system/send-message-broadcast-online
func (rc *RongCloud) MessageOnlineBroadcast(ctx context.Context, req *MessageOnlineBroadcastRequest) (*MessageOnlineBroadcastResponse, error) {
	path := "/message/online/broadcast.json"
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
	resp := &MessageOnlineBroadcastResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageSystemPublishRequest struct {
	// [必传] 发送人用户 ID，通过 Server API 非群成员也可以向群组中发送消息。
	//注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserId *string `url:"fromUserId,omitempty"`

	// [必传] 接收用户 ID，提供多个本参数可以实现向多用户发送系统消息，上限为 100 人。
	ToUserId []string `url:"toUserId,omitempty"`

	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息)
	RCMsg RCMsg `url:"-"`

	// 指定收件人离线时触发的远程推送通知中的通知内容。注意：对于部分消息类型，该字段是否有值决定了是否触发远程推送通知。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中的用户内容类消息格式，可不填写该字段，远程推送通知默认使用服务端预置的推送通知内容。
	// 如果消息类型（objectName 字段）为即时通讯服务预定义消息类型中通知类、信令类（"撤回命令消息" 除外），且需要支持远程推送通知，则必须填写 pushContent，否则收件人不在线时无法收到远程推送通知。如无需触发远程推送，可不填该字段。
	// 如果消息类型为自定义消息类型，且需要支持远程推送通知，则必须填写 pushContent 字段，否则收件人不在线时无法收到远程推送通知。
	// 如果消息类型为自定义消息类型，但不需要支持远程推送通知（例如通过自定义消息类型实现的 App 业务层操作指令），可将 pushContent 字段留空禁用远程推送通知。
	PushContent *string `url:"pushContent,omitempty"`

	// iOS 平台收到推送消息时，可从 payload 中获取 APNs 推送数据，对应字段名为 appData（提示：rc 字段中默认携带了消息基本信息）。Android 平台收到推送消息时对应字段名为 appData。
	PushData *string `url:"pushData,omitempty"`

	// 是否需要为收件人在历史消息云端存储服务中存储此条消息。0 表示不存储；1 表示存储。默认值为 1，存储（依赖单群聊消息云端存储服务）。
	// 注意：即使已开通单群聊消息云存储功能，群组定向消息也不会存入服务端历史消息记录。如有需要，请提交工单申请开通群定向消息云存储服务。
	// 此属性不影响离线消息功能，用户未在线时都会转为离线消息?存储。
	// 提示：一般情况下（第 1、2 种情况），客户端是否存储消息不依赖此参数。以下第 3 种情况属于例外：
	// 如果消息属于内置消息类型，客户端 SDK 会根据消息类型本身的存储属性标识判断是否存入本地数据库。详见消息类型概述。
	// 如果消息属于自定义消息类型，则客户端 SDK 会根据该类型在客户端上注册时的存储属性标识判断是否需要存入本地数据库。
	// 如果消息属于客户端 App 上未注册自定义消息类型（例如客户端使用的 App 版本过旧），则客户端 SDK 会根据当前参数值确定是否将消息存储在本地。但因消息类型未注册，客户端无法解析显示该消息。
	IsPersisted *int `url:"isPersisted,omitempty"`

	// 仅目标用户为 iOS 设备时有效，应用处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看[知识库文档](https://help.rongcloud.cn/t/topic/855)。1 表示为开启，0 表示为关闭，默认为 0。
	ContentAvailable *int `url:"contentAvailable,omitempty"`

	// 是否为静默消息，默认为 false，设为 true 时终端用户离线情况下不会收到通知提醒。
	DisablePush *bool `url:"disablePush,omitempty"`

	// 配置消息的推送通知，如推送通知的标题等。disablePush 属性为 true 时此属性无效。json string
	PushExt *string `url:"pushExt,omitempty"`
}

type MessageSystemPublishResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageSystemPublish 发送系统通知普通消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/system/send-private
func (rc *RongCloud) MessageSystemPublish(ctx context.Context, req *MessageSystemPublishRequest) (*MessageSystemPublishResponse, error) {
	path := "/message/system/publish.json"
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
	resp := &MessageSystemPublishResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageSystemPublishTemplateRequest struct {
	// [必传] 发送人用户 ID。
	// 注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserId *string `json:"fromUserId,omitempty"`
	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息)
	RCMsg RCMsg `json:"-"`
	// [必传] SDK封装消息模板
	MessageTemplate []*MessageTemplate `json:"-"`
	// 针对 iOS 平台，对 SDK 处于后台暂停状态时为静默推送，是 iOS7 之后推出的一种推送方式。 允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看知识库文档。1 表示为开启，0 表示为关闭，默认为 0
	ContentAvailable *int `json:"contentAvailable,omitempty"`
	// 是否为静默消息，默认为 false，设为 true 时终端用户离线情况下不会收到通知提醒。
	DisablePush *bool `json:"disablePush,omitempty"`
}

func (r *MessageSystemPublishTemplateRequest) MarshalJSON() ([]byte, error) {
	req := messagePublishTemplateRequest{
		FromUserId:       r.FromUserId,
		ContentAvailable: r.ContentAvailable,
		DisablePush:      r.DisablePush,
	}
	if r.RCMsg != nil {
		req.ObjectName = r.RCMsg.ObjectName()
		content, err := r.RCMsg.ToString()
		if err != nil {
			return nil, fmt.Errorf("%s RcMsg.ToString() error %s", req.ObjectName, err)
		}
		req.Content = content
	}
	if r.MessageTemplate != nil {
		var values []map[string]string
		var toUserIds []string
		var pushContent []string
		var pushData []string
		for _, template := range r.MessageTemplate {
			values = append(values, template.Value)
			toUserIds = append(toUserIds, template.ToUserId)
			pushContent = append(pushContent, template.PushContent)
			pushData = append(pushData, template.PushData)
		}
		req.Values = values
		req.ToUserId = toUserIds
		req.PushContent = pushContent
		req.PushData = pushData
	}
	return json.Marshal(req)
}

type MessageSystemPublishTemplateResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageSystemPublishTemplate 发送系统通知模板消息
// More details see https://doc.rongcloud.cn/imserver/server/v1/system/send-private-template
func (rc *RongCloud) MessageSystemPublishTemplate(ctx context.Context, req *MessageSystemPublishTemplateRequest) (*MessageSystemPublishTemplateResponse, error) {
	path := "/message/system/publish_template.json"
	resp := &MessageSystemPublishTemplateResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageHistoryRequest struct {
	// [必传] 指定时间，精确到某天某小时，格式为 YYYYMMDDHH。例如 2014010101 表示需要获取 2014 年 1 月 1 日凌晨 1 点至 2 点的数据。
	// 注意：date 的值与应用所属数据中心有关。如您的 App 业务使用新加坡数据中心，则获取消息日志时使用的时间（date），及日志中的消息时间（dateTime）均为 UTC 时间。如您仍需根据北京时间下载数据，请自行转换处理。如要下载北京时间 2019120109 的日志，需要输入 2019120101。
	Date *string `url:"date,omitempty"`
}

type MessageHistoryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	// 历史记录下载地址。如没有消息记录数据时，则 url 值为空。
	Url string `json:"url"`
	// 历史记录时间。
	Date string `json:"date"`
}

// MessageHistory 获取历史消息日志
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/get-message-history-log
func (rc *RongCloud) MessageHistory(ctx context.Context, req *MessageHistoryRequest) (*MessageHistoryResponse, error) {
	path := "/message/history.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &MessageHistoryResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type MessageHistoryDeleteRequest struct {
	// [必传] 指定时间，精确到某天某小时，格式为 YYYYMMDDHH。例如 2014010101 表示需要删除 2014 年 1 月 1 日凌晨 1 点至 2 点的数据。返回成功后，消息记录文件将在随后的 10 分钟内被永久删除。
	// 注意：date 的值与应用所属数据中心有关。如您的 App 业务使用新加坡数据中心，则获取消息日志时使用的时间（date），及日志中的消息时间（dateTime）均为 UTC 时间。如您仍需根据北京时间下载数据，请自行转换处理。如要下载北京时间 2019120109 的日志，需要输入 2019120101。
	Date *string `url:"date,omitempty"`
}

type MessageHistoryDeleteResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// MessageHistoryDelete 删除历史消息日志
// More details see https://doc.rongcloud.cn/imserver/server/v1/message/delete-message-history-log
func (rc *RongCloud) MessageHistoryDelete(ctx context.Context, req *MessageHistoryDeleteRequest) (*MessageHistoryDeleteResponse, error) {
	path := "/message/history/delete.json"
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	resp := &MessageHistoryDeleteResponse{}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
