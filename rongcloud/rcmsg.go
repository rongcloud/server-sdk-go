package rongcloud

import (
	"encoding/json"
	"net/url"
)

// RCMsg
// 消息类型，接受内置消息类型（见消息类型概述）或自定义消息的消息类型值。
// 注意：在自定义消息时，消息类型不可以 "RC:" 开头，以免与系统内置消息类型重名；消息类型长度不可超过 32 个字符。SDK 中必须已注册过该自定义消息，否则 SDK 收到该消息后将无法解析。
// 所发送消息的内容，单条消息最大 128k。
// 内置消息类型：将消息内容体 JSON 对象序列化为 JSON 字符串传入。消息内容 JSON 结构体详见用户内容类消息格式或其他内置消息类型的消息内容格式。
// 例如，文本消息内容 JSON 结构体内部包含 content 字段（此为 JSON 结构体内的 key 值，注意区分），则需要将 {"content":"Hello world!"} 序列化后的结果作为此处 content 字段的值。
// 自定义消息类型（objectName 字段必须指定为自定义消息类型）：如果发送自定义消息，该参数可自定义格式，不限于 JSON。
type RCMsg interface {
	ObjectName() string
	ToString() (string, error)
	EncodeValues(key string, v *url.Values) error
}

// MsgUserInfo 融云内置消息用户信息
type MsgUserInfo struct {
	ID       string `json:"id,omitempty"`   // 消息中携带的消息发送者的用户信息。一般情况下不建议在消息中携带用户信息。建议仅在直播场景下使用。
	Name     string `json:"name,omitempty"` // 消息发送者的用户昵称。
	Icon     string `json:"icon,omitempty"`
	Portrait string `json:"portrait,omitempty"` // 消息发送者的头象。
	Extra    string `json:"extra,omitempty"`    // 扩展信息，可以放置任意的数据内容。
}

// TXTMsg 消息
type TXTMsg struct {
	Content string       `json:"content,omitempty"` // [必传] 文字消息的文字内容，包括表情。
	User    *MsgUserInfo `json:"user,omitempty"`    // 消息中携带的消息发送者的用户信息。一般情况下不建议在消息中携带用户信息。建议仅在直播场景下使用。
	Extra   string       `json:"extra,omitempty"`   // 扩展信息，可以放置任意的数据内容，也可以去掉此属性。
}

func (m *TXTMsg) ObjectName() string {
	return "RC:TxtMsg"
}

func (m *TXTMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *TXTMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// ImgMsg 消息
type ImgMsg struct {
	Content  string       `json:"content,omitempty"`  // [必传] 图片缩略图进行 Base64 编码的结果值。Base64 字符串长度建议为 5k，最大不超过 10k。注意在 Base64 进行 Encode 后需要将所有 \r\n 和 \r 和 \n 替换成空。
	User     *MsgUserInfo `json:"user,omitempty"`     // 消息中携带的消息发送者的用户信息。一般情况下不建议在消息中携带用户信息。建议仅在直播场景下使用。
	ImageURI string       `json:"imageUri,omitempty"` // [必传] 图片上传到图片存储服务器后的地址。通过 IM Server API 发送图片消息时，需要自行上传文件到应用的文件服务器，生成图片地址后进行发送。
	Extra    string       `json:"extra,omitempty"`
}

func (m *ImgMsg) ObjectName() string {
	return "RC:ImgMsg"
}

func (m *ImgMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *ImgMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// InfoNtf 提示（小灰条）通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification
type InfoNtf struct {
	Message string       `json:"message,omitempty"`
	User    *MsgUserInfo `json:"user,omitempty"`
	Extra   string       `json:"extra,omitempty"`
}

func (m *InfoNtf) ObjectName() string {
	return "RC:InfoNtf"
}

func (m *InfoNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *InfoNtf) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// VCMsg 旧版语音消息, 推荐使用高清语音消息 HQVCMsg
// Deprecated
type VCMsg struct {
	Content  string       `json:"content,omitempty"`
	User     *MsgUserInfo `json:"user,omitempty"`
	Extra    string       `json:"extra,omitempty"`
	Duration interface{}  `json:"duration,omitempty"`
}

func (m *VCMsg) ObjectName() string {
	return "RC:VCMsg"
}

func (m *VCMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *VCMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// HQVCMsg 高清语音消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E9%AB%98%E6%B8%85%E8%AF%AD%E9%9F%B3%E6%B6%88%E6%81%AF%E5%86%85%E5%AE%B9%E7%BB%93%E6%9E%84%E4%BD%93
type HQVCMsg struct {
	RemoteUrl string       `json:"remoteUrl,omitempty"` // 必传参数, 媒体内容上传服务器后的网络地址。通过 IM 服务端 API 发送高质量语音消息时，需要自行生成 AAC 格式文件并上传文件到应用的文件服务器，生成地址后进行发送。
	Duration  int          `json:"duration,omitempty"`  // 必传参数, 语音消息的时长，最长为 60 秒（单位：秒）。
	User      *MsgUserInfo `json:"user,omitempty"`
	Extra     string       `json:"extra,omitempty"` // 扩展信息，可以放置任意的数据内容，也可以去掉此属性。
}

func (m *HQVCMsg) ObjectName() string {
	return "RC:HQVCMsg"
}

func (m *HQVCMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *HQVCMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// IMGTextMsg 图文消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E5%9B%BE%E6%96%87%E6%B6%88%E6%81%AF%E5%86%85%E5%AE%B9%E7%BB%93%E6%9E%84%E4%BD%93
type IMGTextMsg struct {
	Title    string       `json:"title,omitempty"`   // 消息的标题。
	Content  string       `json:"content,omitempty"` // 消息的文字内容。
	User     *MsgUserInfo `json:"user,omitempty"`
	Extra    string       `json:"extra,omitempty"`
	ImageUri string       `json:"imageUri,omitempty"` // 消息中图片地址，图片尺寸为：120 x 120 像素。通过 IM Server API 发送图片消息时，需要自行上传文件到应用的文件服务器，生成图片地址后进行发送。
	URL      string       `json:"url,omitempty"`      // 点击图片消息后跳转的 URL 地址。
}

func (m *IMGTextMsg) ObjectName() string {
	return "RC:ImgTextMsg"
}

func (m *IMGTextMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *IMGTextMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// FileMsg 文件消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E6%96%87%E4%BB%B6%E6%B6%88%E6%81%AF
type FileMsg struct {
	Name    string       `json:"name,omitempty"`    // 文件名称。如不传，调用客户端 SDK 中的 downloadMediaMessage 方法下载后会默认生成一个名称。
	Size    string       `json:"size,omitempty"`    // [必传参数] 文件大小，单位：Byte。
	Type    string       `json:"type,omitempty"`    // [必传参数] 文件类型
	FileURL string       `json:"fileUrl,omitempty"` // [必传参数] 文件的服务器地址。通过 IM Server API 发送文件消息时，需要自行上传文件到应用的文件服务器，生成文件地址后进行发送。
	User    *MsgUserInfo `json:"user,omitempty"`
}

func (m *FileMsg) ObjectName() string {
	return "RC:FileMsg"
}

func (m *FileMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *FileMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// LBSMsg 位置消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E4%BD%8D%E7%BD%AE%E6%B6%88%E6%81%AF
type LBSMsg struct {
	Content   string       `json:"content,omitempty"`
	Extra     string       `json:"extra,omitempty"`
	POI       string       `json:"poi,omitempty"`
	Latitude  float64      `json:"latitude,omitempty"`
	Longitude float64      `json:"longitude,omitempty"`
	User      *MsgUserInfo `json:"user,omitempty"`
}

func (m *LBSMsg) ObjectName() string {
	return "RC:LBSMsg"
}

func (m *LBSMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *LBSMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// ProfileNtf 资料变更通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification#%E8%B5%84%E6%96%99%E5%8F%98%E6%9B%B4%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
type ProfileNtf struct {
	Operation string       `json:"operation,omitempty"`
	Data      string       `json:"data,omitempty"`
	User      *MsgUserInfo `json:"user,omitempty"`
	Extra     string       `json:"extra,omitempty"`
}

func (m *ProfileNtf) ObjectName() string {
	return "RC:ProfileNtf"
}

func (m *ProfileNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *ProfileNtf) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// CMDNtf 命令提醒消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification#%E5%91%BD%E4%BB%A4%E6%8F%90%E9%86%92%E6%B6%88%E6%81%AF
type CMDNtf struct {
	Name string       `json:"operation,omitempty"`
	Data string       `json:"data,omitempty"`
	User *MsgUserInfo `json:"user,omitempty"`
}

func (m *CMDNtf) ObjectName() string {
	return "RC:CMDNtf"
}

func (m *CMDNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *CMDNtf) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// CMDMsg 命令消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-callback#%E5%91%BD%E4%BB%A4%E6%B6%88%E6%81%AF
type CMDMsg struct {
	Name string       `json:"name,omitempty"`
	Data string       `json:"data,omitempty"`
	User *MsgUserInfo `json:"user,omitempty"`
}

func (m *CMDMsg) ObjectName() string {
	return "RC:CMDMsg"
}

func (m *CMDMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *CMDMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// ContactNtf 联系人（好友）通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification#%E8%81%94%E7%B3%BB%E4%BA%BA%E5%A5%BD%E5%8F%8B%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
type ContactNtf struct {
	Operation    string       `json:"operation,omitempty"`    // [必传]联系人操作的指令，官方针对 operation 属性定义了 "Request", "AcceptResponse", "RejectResponse" 几个常量，也可以由开发者自行扩展。
	SourceUserID string       `json:"sourceUserId,omitempty"` // [必传]发出通知的用户 Id。
	TargetUserID string       `json:"targetUserId,omitempty"` // [必传]单聊会话为接收通知的用户 Id，群聊、聊天室会话为会话 Id。
	Message      string       `json:"message,omitempty"`      // [必传]表示请求或者响应消息，如添加理由或拒绝理由。
	Extra        string       `json:"extra,omitempty"`
	User         *MsgUserInfo `json:"user,omitempty"`
}

func (m *ContactNtf) ObjectName() string {
	return "RC:ContactNtf"
}

func (m *ContactNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *ContactNtf) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// GrpNtf 群组通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-notification#%E7%BE%A4%E7%BB%84%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
type GrpNtf struct {
	OperatorUserID string       `json:"operatorUserId,omitempty"` // [必传] 操作人用户 Id
	Operation      string       `json:"operation,omitempty"`      // [必传] 群组中各种通知的操作名称。
	Data           string       `json:"data,omitempty"`           // [必传] 操作数据
	Message        string       `json:"message,omitempty"`        // [必传] 消息内容
	Extra          string       `json:"extra,omitempty"`
	User           *MsgUserInfo `json:"user,omitempty"`
}

func (m *GrpNtf) ObjectName() string {
	return "RC:GrpNtf"
}

func (m *GrpNtf) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *GrpNtf) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}

// ChrmKVNotiMsg 聊天室属性通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-callback#%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%B1%9E%E6%80%A7%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
type ChrmKVNotiMsg struct {
	Type  int    `json:"type,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Extra string `json:"extra,omitempty"`
}

func (m *ChrmKVNotiMsg) ObjectName() string {
	return "RC:ChrmKVNotiMsg"
}

func (m *ChrmKVNotiMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *ChrmKVNotiMsg) EncodeValues(key string, v *url.Values) error {
	return MakeRCMsgUrlValues(m, key, v)
}
