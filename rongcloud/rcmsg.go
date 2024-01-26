package rongcloud

import "encoding/json"

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

// ChrmKVNotiMsg 聊天室属性通知消息
// https://doc.rongcloud.cn/imserver/server/v1/message/objectname-callback#%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%B1%9E%E6%80%A7%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
type ChrmKVNotiMsg struct {
	Type  int    `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Extra string `json:"extra"`
}

func (m *ChrmKVNotiMsg) ObjectName() string {
	return "RC:ChrmKVNotiMsg"
}

func (m *ChrmKVNotiMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}
