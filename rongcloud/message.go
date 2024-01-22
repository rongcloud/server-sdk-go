package rongcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type RCMsg interface {
	ObjectName() string
	ToString() (string, error)
}

// MsgUserInfo 融云内置消息用户信息
type MsgUserInfo struct {
	ID       string `json:"id"`   // 消息中携带的消息发送者的用户信息。一般情况下不建议在消息中携带用户信息。建议仅在直播场景下使用。
	Name     string `json:"name"` // 消息发送者的用户昵称。
	Icon     string `json:"icon"`
	Portrait string `json:"portrait"` // 消息发送者的头象。
	Extra    string `json:"extra"`    // 扩展信息，可以放置任意的数据内容。
}

// TXTMsg 消息
type TXTMsg struct {
	Content string      `json:"content"`
	User    MsgUserInfo `json:"user"`
	Extra   string      `json:"extra"`
}

func (m *TXTMsg) ObjectName() string {
	return "RC:TxtMsg"
}

func (m *TXTMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// ImgMsg 消息
type ImgMsg struct {
	Content  string      `json:"content"`
	User     MsgUserInfo `json:"user"`
	ImageURI string      `json:"imageUri"`
	Extra    string      `json:"extra"`
}

func (m *ImgMsg) ObjectName() string {
	return "RC:ImgMsg"
}

func (m *ImgMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// InfoNtf 提示（小灰条）通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification
type InfoNtf struct {
	Message string      `json:"message"`
	User    MsgUserInfo `json:"user"`
	Extra   string      `json:"extra"`
}

func (m *InfoNtf) ObjectName() string {
	return "RC:InfoNtf"
}

func (m *InfoNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// VCMsg 旧版语音消息, 推荐使用高清语音消息 HQVCMsg
// Deprecated
type VCMsg struct {
	Content  string      `json:"content"`
	User     MsgUserInfo `json:"user"`
	Extra    string      `json:"extra"`
	Duration interface{} `json:"duration"`
}

func (m *VCMsg) ObjectName() string {
	return "RC:VCMsg"
}

func (m *VCMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// 高清语音消息 RC:HQVCMsg
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E9%AB%98%E6%B8%85%E8%AF%AD%E9%9F%B3%E6%B6%88%E6%81%AF%E5%86%85%E5%AE%B9%E7%BB%93%E6%9E%84%E4%BD%93
type HQVCMsg struct {
	RemoteUrl string      `json:"remoteUrl"` // 必传参数, 媒体内容上传服务器后的网络地址。通过 IM 服务端 API 发送高质量语音消息时，需要自行生成 AAC 格式文件并上传文件到应用的文件服务器，生成地址后进行发送。
	Duration  int         `json:"duration"`  // 必传参数, 语音消息的时长，最长为 60 秒（单位：秒）。
	User      MsgUserInfo `json:"user"`
	Extra     string      `json:"extra"` // 扩展信息，可以放置任意的数据内容，也可以去掉此属性。
}

func (m *HQVCMsg) ObjectName() string {
	return "RC:HQVCMsg"
}

func (m *HQVCMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// IMGTextMsg 图文消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E5%9B%BE%E6%96%87%E6%B6%88%E6%81%AF%E5%86%85%E5%AE%B9%E7%BB%93%E6%9E%84%E4%BD%93
type IMGTextMsg struct {
	Title    string      `json:"title"`   // 消息的标题。
	Content  string      `json:"content"` // 消息的文字内容。
	User     MsgUserInfo `json:"user"`
	Extra    string      `json:"extra"`
	ImageUri string      `json:"imageUri"` // 消息中图片地址，图片尺寸为：120 x 120 像素。通过 IM Server API 发送图片消息时，需要自行上传文件到应用的文件服务器，生成图片地址后进行发送。
	URL      string      `json:"url"`      // 点击图片消息后跳转的 URL 地址。
}

func (m *IMGTextMsg) ObjectName() string {
	return "RC:ImgTextMsg"
}

func (m *IMGTextMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// FileMsg 文件消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E6%96%87%E4%BB%B6%E6%B6%88%E6%81%AF
type FileMsg struct {
	Name    string      `json:"name,omitempty"`    // 文件名称。如不传，调用客户端 SDK 中的 downloadMediaMessage 方法下载后会默认生成一个名称。
	Size    string      `json:"size,omitempty"`    // [必传参数] 文件大小，单位：Byte。
	Type    string      `json:"type,omitempty"`    // [必传参数] 文件类型
	FileURL string      `json:"fileUrl,omitempty"` // [必传参数] 文件的服务器地址。通过 IM Server API 发送文件消息时，需要自行上传文件到应用的文件服务器，生成文件地址后进行发送。
	User    MsgUserInfo `json:"user,omitempty"`
}

func (m *FileMsg) ObjectName() string {
	return "RC:FileMsg"
}

func (m *FileMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// LBSMsg 位置消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E4%BD%8D%E7%BD%AE%E6%B6%88%E6%81%AF
type LBSMsg struct {
	Content   string      `json:"content"`
	Extra     string      `json:"extra"`
	POI       string      `json:"poi"`
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
	User      MsgUserInfo `json:"user"`
}

func (m *LBSMsg) ObjectName() string {
	return "RC:LBSMsg"
}

func (m *LBSMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// ProfileNtf 资料变更通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification#%E8%B5%84%E6%96%99%E5%8F%98%E6%9B%B4%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
type ProfileNtf struct {
	Operation string      `json:"operation"`
	Data      string      `json:"data"`
	User      MsgUserInfo `json:"user"`
	Extra     string      `json:"extra"`
}

func (m *ProfileNtf) ObjectName() string {
	return "RC:ProfileNtf"
}

func (m *ProfileNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// CMDNtf 命令提醒消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification#%E5%91%BD%E4%BB%A4%E6%8F%90%E9%86%92%E6%B6%88%E6%81%AF
type CMDNtf struct {
	Name string      `json:"operation"`
	Data string      `json:"data"`
	User MsgUserInfo `json:"user"`
}

func (m *CMDNtf) ObjectName() string {
	return "RC:CMDNtf"
}

func (m *CMDNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// CMDMsg 命令消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-callback#%E5%91%BD%E4%BB%A4%E6%B6%88%E6%81%AF
type CMDMsg struct {
	Name string      `json:"name"`
	Data string      `json:"data"`
	User MsgUserInfo `json:"user"`
}

func (m *CMDMsg) ObjectName() string {
	return "RC:CMDMsg"
}

func (m *CMDMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// ContactNtf 联系人（好友）通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification#%E8%81%94%E7%B3%BB%E4%BA%BA%E5%A5%BD%E5%8F%8B%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
type ContactNtf struct {
	Operation    string      `json:"operation"`    // [必传]联系人操作的指令，官方针对 operation 属性定义了 "Request", "AcceptResponse", "RejectResponse" 几个常量，也可以由开发者自行扩展。
	SourceUserID string      `json:"sourceUserId"` // [必传]发出通知的用户 Id。
	TargetUserID string      `json:"targetUserId"` // [必传]单聊会话为接收通知的用户 Id，群聊、聊天室会话为会话 Id。
	Message      string      `json:"message"`      // [必传]表示请求或者响应消息，如添加理由或拒绝理由。
	Extra        string      `json:"extra"`
	User         MsgUserInfo `json:"user"`
}

func (m *ContactNtf) ObjectName() string {
	return "RC:ContactNtf"
}

func (m *ContactNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// GrpNtf 群组通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification#%E7%BE%A4%E7%BB%84%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
type GrpNtf struct {
	OperatorUserID string      `json:"operatorUserId"` // [必传] 操作人用户 Id
	Operation      string      `json:"operation"`      // [必传] 群组中各种通知的操作名称。
	Data           string      `json:"data"`           // [必传] 操作数据
	Message        string      `json:"message"`        // [必传] 消息内容
	Extra          string      `json:"extra"`
	User           MsgUserInfo `json:"user"`
}

func (m *GrpNtf) ObjectName() string {
	return "RC:GrpNtf"
}

func (m *GrpNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

// MessagePrivatePublishRequest 单聊消息
type MessagePrivatePublishRequest struct {
	// 发送人用户 ID。
	//注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserId string `json:"fromUserId"`

	// 接收用户 ID，可以实现向多人发送消息，每次上限为 1000 人。
	ToUserId string `json:"toUserId"`

	RCMsg    RCMsg   `json:"-"`
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
	ExtraContent *map[string]string `json:"extraContent"`

	// 是否为静默消息，默认为 false，设为 true 时终端用户离线情况下不会收到通知提醒。
	DisablePush *bool `json:"disablePush"`

	// 配置消息的推送通知，如推送通知的标题等。disablePush 属性为 true 时此属性无效。
	PushExt *string `json:"pushExt"`
}

func (r *MessagePrivatePublishRequest) MakeFormData() (url.Values, error) {
	content, err := json.Marshal(r.RCMsg)
	if err != nil {
		return nil, err
	}
	res := url.Values{}
	res.Set("fromUserId", r.FromUserId)
	res.Set("toUserId", r.ToUserId)
	res.Set("objectName", r.RCMsg.ObjectName())
	res.Set("content", string(content))
	if r.PushData != nil {
		res.Set("pushData", *r.PushData)
	}
	if r.PushContent != nil {
		res.Set("pushContent", *r.PushContent)
	}
	if r.IsIncludeSender != nil {
		res.Set("isIncludeSender", strconv.Itoa(*r.IsIncludeSender))
	}
	if r.Count != nil {
		res.Set("count", strconv.Itoa(*r.Count))
	}
	if r.VerifyBlacklist != nil {
		res.Set("verifyBlacklist", strconv.Itoa(*r.VerifyBlacklist))
	}
	if r.IsPersisted != nil {
		res.Set("isPersisted", strconv.Itoa(*r.IsPersisted))
	}
	if r.ContentAvailable != nil {
		res.Set("contentAvailable", strconv.Itoa(*r.ContentAvailable))
	}
	if r.Expansion != nil {
		res.Set("expansion", strconv.FormatBool(*r.Expansion))
	}
	if r.ExtraContent != nil {
		b, err := json.Marshal(r.ExtraContent)
		if err != nil {
			return nil, fmt.Errorf("Marshal ExtraContent err: %w", err)
		}
		res.Set("extraContent", string(b))
	}
	if r.DisablePush != nil {
		res.Set("disablePush", strconv.FormatBool(*r.DisablePush))
	}
	if r.PushExt != nil {
		res.Set("pushExt", *r.PushExt)
	}
	return res, nil
}

type MessagePrivatePublishResponse struct {
	Code int `json:"code"` // 返回码，200 为正常。

	HttpResponseGetter `json:"-"`
}

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
	resp.HttpResponseGetter = &RawHttpResponseGetter{rawHttpResponseInternal: httpResp}
	return resp, nil
}
