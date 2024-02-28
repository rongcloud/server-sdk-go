package rongcloud

import (
	"context"
	"encoding/json"
	"fmt"
)

type UserBlockPushPeriodSetRequest struct {
	// [必传] 用户 ID
	UserId *string `url:"userId,omitempty"`
	// [必传] 开始时间，精确到秒。格式为 HH:MM:SS，示例：22:00:00。注意：startTime 与应用所属数据中心有关。如您的 App 业务使用国内数据中心，请使用北京时间。如您的 App 业务使用海外数据中心，请使用 UTC 时间。
	StartTime *string `url:"startTime,omitempty"`
	// [必传] 免打扰时间窗口大小，单位为分钟。支持范围为 [0-1439] 的整数。0 表示未设置。
	Period *int `url:"period,omitempty"`
	// 免打扰级别。1：仅 @消息进行通知，普通消息不进行通知。0：所有消息都进行通知。5：所有消息都不进行通知。 默认 1。
	Level *int `url:"level,omitempty"`
}

type UserBlockPushPeriodSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserBlockPushPeriodSet 设置用户免打扰时段
// More details see https://doc.rongcloud.cn/imserver/server/v1/push/set-push-disabled-period
func (rc *RongCloud) UserBlockPushPeriodSet(ctx context.Context, req *UserBlockPushPeriodSetRequest) (*UserBlockPushPeriodSetResponse, error) {
	path := "/user/blockPushPeriod/set.json"
	resp := &UserBlockPushPeriodSetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserBlockPushPeriodGetRequest struct {
	// [必传] 用户 ID
	UserId *string `url:"userId,omitempty"`
}

type UserBlockPushPeriodGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Data               *UserBlockPushPeriodData `json:"data"`
}

type UserBlockPushPeriodData struct {
	// 开始时间，精确到秒。格式为 HH:MM:SS，示例：22:00:00。注意：startTime 与应用所属数据中心有关。如您的 App 业务使用国内数据中心，该时间为北京时间。如您的 App 业务使用海外数据中心，该时间为 UTC 时间。
	StartTime string `json:"startTime"`
	// 免打扰时间窗口大小，单位为分钟。范围为 [0-1439] 的整数。0 表示未设置。
	Period int `json:"period"`
	// 免打扰级别。1：仅 @消息进行通知，普通消息不进行通知。0：所有消息都进行通知。5：所有消息都不进行通知。
	UnPushLevel int `json:"unPushLevel"`
}

// UserBlockPushPeriodGet 查询用户免打扰时段
// More details see https://doc.rongcloud.cn/imserver/server/v1/push/get-push-disabled-period
func (rc *RongCloud) UserBlockPushPeriodGet(ctx context.Context, req *UserBlockPushPeriodGetRequest) (*UserBlockPushPeriodGetResponse, error) {
	path := "/user/blockPushPeriod/get.json"
	resp := &UserBlockPushPeriodGetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserBlockPushPeriodDeleteRequest struct {
	// [必传] 用户ID
	UserId *string `url:"userId,omitempty"`
}

type UserBlockPushPeriodDeleteResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserBlockPushPeriodDelete 删除用户免打扰时段
// More details see https://doc.rongcloud.cn/imserver/server/v1/push/delete-push-disabled-period
func (rc *RongCloud) UserBlockPushPeriodDelete(ctx context.Context, req *UserBlockPushPeriodDeleteRequest) (*UserBlockPushPeriodDeleteResponse, error) {
	path := "/user/blockPushPeriod/delete.json"
	resp := &UserBlockPushPeriodDeleteResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserGetTokenRequest struct {
	// [必传] App 自行定义的用户 ID，用于换取 Token。支持大小写英文字母与数字的组合，最大长度 64 字节。
	UserId *string `url:"userId,omitempty"`
	// [必传] 推送服务使用的用户名称。不区分符号、英文字符、中文字符，统一限制最多 64 个字符。注意：该 name 字段仅用于推送服务，作为在移动客户端推送通知中默认显示的用户名称。因为即时通讯服务端不提供用户信息托管服务，所以不支持客户端 SDK 主动获取该用户名称数据。
	Name *string `url:"name,omitempty"`
	// 用户头像 URI，最大长度 1024 字节。注意：因为即时通讯服务端不提供用户信息托管服务，所以不支持客户端 SDK 主动获取该用户头像数据。
	PortraitUri *string `url:"portraitUri,omitempty"`
}

type UserGetTokenResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	// 用户身份验证 Token，长度在 256 字节以内，可以保存应用内。Token 中携带 IM 服务动态导航地址，开发者不需要进行处理。
	Token string `json:"token"`
	// 返回输入参数中提供的用户 ID。
	UserId string `json:"userId"`
}

// UserGetToken 注册用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/register
func (rc *RongCloud) UserGetToken(ctx context.Context, req *UserGetTokenRequest) (*UserGetTokenResponse, error) {
	path := "/user/getToken.json"
	resp := &UserGetTokenResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserTokenExpireRequest struct {
	// [必传] 需要设置 Token 失效的用户 ID，支持设置多个最多不超过 20 个。
	UserId []string `url:"userId,omitempty"`
	// [必传] 过期时间戳精确到毫秒，该时间戳前用户获取的 Token 全部失效，使用时间戳之前的 Token 已经在连接中的用户不会立即失效，断开后无法进行连接。
	Time *int64 `url:"time,omitempty"`
}

type UserTokenExpireResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserTokenExpire 作废Token
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/expire
func (rc *RongCloud) UserTokenExpire(ctx context.Context, req *UserTokenExpireRequest) (*UserTokenExpireResponse, error) {
	path := "/user/token/expire.json"
	resp := &UserTokenExpireResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserRemarksSetRequest struct {
	// [必传] 用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 设置的目标用户推送备注名
	Remarks []*UserRemark `url:"-"`
}

type UserRemark struct {
	// [必传] 目标用户 ID。单次最多设置 100 个。
	Id string `json:"id,omitempty"`
	// [必传] 收到目标用户推送时显示的备注名。
	Remark string `json:"remark,omitempty"`
}

type UserRemarksSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserRemarksSet 设置用户级推送备注名
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/set-remark-for-push
func (rc *RongCloud) UserRemarksSet(ctx context.Context, req *UserRemarksSetRequest) (*UserRemarksSetResponse, error) {
	path := "/user/remarks/set.json"
	resp := &UserRemarksSetResponse{}
	params, err := makeUrlValues(req)
	if err != nil {
		return nil, err
	}
	if req.Remarks != nil {
		remarks, err := json.Marshal(req.Remarks)
		if err != nil {
			return nil, NewSDKError(fmt.Sprintf("json marshal remarks error %s", err))
		}
		params.Set("remarks", string(remarks))
	}
	httpResp, err := rc.postFormUrlencoded(ctx, path, params, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserRemarksGetRequest struct {
	// [必传] 用户 ID。
	UserId *string `url:"userId,omitempty"`
	// 页数，默认为第一页。
	Page *int `url:"page,omitempty"`
	// 每页条数，默认每页 50 条。
	Size *int `url:"size,omitempty"`
}

type UserRemarksGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Total              int           `json:"total"` // 用户的备注名总数。
	Users              []*UserRemark `json:"users"` // 单次最多返回 50 个用户备注名。
}

// UserRemarksGet 查询用户级推送备注名
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/get-remark-for-push
func (rc *RongCloud) UserRemarksGet(ctx context.Context, req *UserRemarksGetRequest) (*UserRemarksGetResponse, error) {
	path := "/user/remarks/get.json"
	resp := &UserRemarksGetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserRemarksDelRequest struct {
	// [必传] 操作者用户ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 需要删除推送备注名的用户 ID。
	TargetId *string `url:"targetId,omitempty"`
}

type UserRemarksDelResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserRemarksDel 删除用户级推送备注名
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/delete-remark-for-push
func (rc *RongCloud) UserRemarksDel(ctx context.Context, req *UserRemarksDelRequest) (*UserRemarksDelResponse, error) {
	path := "/user/remarks/del.json"
	resp := &UserRemarksDelResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserChatFBSetRequest struct {
	// [必传] 被禁言用户 ID，支持批量设置，最多不超过 1000 个。
	UserId []string `url:"userId,omitempty"`
	// [必传] 禁言状态，0 解除禁言、1 添加禁言
	State *int `url:"state,omitempty"`
	// [必传] 会话类型，目前支持单聊会话 PERSON
	Type *string `url:"type,omitempty"`
}

type UserChatFBSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserChatFBSet 设置用户单聊禁言
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/ban
func (rc *RongCloud) UserChatFBSet(ctx context.Context, req *UserChatFBSetRequest) (*UserChatFBSetResponse, error) {
	path := "/user/chat/fb/set.json"
	resp := &UserChatFBSetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserChatFBQueryListRequest struct {
	// 获取行数，默认为 100，最大支持 200 个。
	Num *int `url:"num,omitempty"`
	// 查询开始位置，默认为 0。
	Offset *int `url:"offset,omitempty"`
	// [必传] 会话类型，目前支持单聊会话 PERSON
	Type *string `url:"type,omitempty"`
}

type UserChatFBQueryListResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Total              int      `json:"total"`
	Users              []string `json:"users"`
}

// UserChatFBQueryList 查询单聊禁言用户列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/query-banned-list
func (rc *RongCloud) UserChatFBQueryList(ctx context.Context, req *UserChatFBQueryListRequest) (*UserChatFBQueryListResponse, error) {
	path := "/user/chat/fb/querylist.json"
	resp := &UserChatFBQueryListResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserWhitelistAddRequest struct {
	// [必传] 用户ID
	UserId *string `url:"userId,omitempty"`
	// [必传] 被加入白名单的用户 ID。单次可添加最多 20 个 whiteUserId。
	WhiteUserId []string `url:"whiteUserId,omitempty"`
}

type UserWhitelistAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserWhitelistAdd 添加用户到白名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/add-to-whitelist
func (rc *RongCloud) UserWhitelistAdd(ctx context.Context, req *UserWhitelistAddRequest) (*UserWhitelistAddResponse, error) {
	path := "/user/whitelist/add.json"
	resp := &UserWhitelistAddResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserWhitelistRemoveRequest struct {
	// [必传] 用户ID
	UserId *string `url:"userId,omitempty"`
	// [必传] 被移除的用户 ID，每次最多移除 20 个用户。
	WhiteUserId []string `url:"whiteUserId,omitempty"`
}

type UserWhitelistRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserWhitelistRemove 移除白名单中用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/remove-from-whitelist
func (rc *RongCloud) UserWhitelistRemove(ctx context.Context, req *UserWhitelistRemoveRequest) (*UserWhitelistRemoveResponse, error) {
	path := "/user/whitelist/remove.json"
	resp := &UserWhitelistRemoveResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserWhitelistQueryRequest struct {
	// [必传] 用户ID
	UserId *string `url:"userId,omitempty"`
}

type UserWhitelistQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []string `json:"users"` // 白名单用户数组。
}

// UserWhitelistQuery 查询白名单中用户列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/query-whitelist
func (rc *RongCloud) UserWhitelistQuery(ctx context.Context, req *UserWhitelistQueryRequest) (*UserWhitelistQueryResponse, error) {
	path := "/user/whitelist/query.json"
	resp := &UserWhitelistQueryResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserRefreshRequest struct {
	// [必传] 用户 ID，支持大小写英文字母与数字的组合，最大长度 64 字节。 userId 是用户在 App 中的唯一标识，必须保证在同一个 App 内不重复，重复的用户 ID 将被当作是同一用户。
	UserId *string `url:"userId,omitempty"`
	// 用户名称，最大长度 64 个字符（不区分符号、英文字符、中文字符，统一限制最多 64 个字符）。用来在 Push 推送时，显示用户的名称，不提供则不进行刷新。
	Name *string `url:"name,omitempty"`
	// 用户头像 URI，最大长度 1024 字节。
	PortraitUri *string `url:"portraitUri,omitempty"`
}

type UserRefreshResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserRefresh 修改用户信息
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/modify
func (rc *RongCloud) UserRefresh(ctx context.Context, req *UserRefreshRequest) (*UserRefreshResponse, error) {
	path := "/user/refresh.json"
	resp := &UserRefreshResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserBlockRequest struct {
	// [必传] 用户 ID，支持一次封禁多个用户，最多不超过 20 个。
	UserId []string `url:"userId,omitempty"`
	// [必传] 封禁时长，单位为分钟，最大值为 43200 分钟。
	Minute *int `url:"minute,omitempty"`
}

type UserBlockResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserBlock 封禁用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/block
func (rc *RongCloud) UserBlock(ctx context.Context, req *UserBlockRequest) (*UserBlockResponse, error) {
	path := "/user/block.json"
	resp := &UserBlockResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserBlockQueryRequest struct {
	// 分页获取封禁用户列表时当前页数，不传或传入 0 时不做分页处理，默认获取前 1000 个被封禁的用户列表，按封禁结束时间倒序排序。
	Page *int `url:"page,omitempty"`
	// 分页获取封禁用户列表时每页行数，不传时默认为 50 条。
	Size *int `url:"size,omitempty"`
}

type UserBlockQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []UserBlockQueryUser `json:"users"` // 被封禁用户数组。
}

type UserBlockQueryUser struct {
	UserId       string `json:"userId"`       // 被封禁用户 ID。
	BlockEndTime string `json:"blockEndTime"` // 封禁结束时间。
}

// UserBlockQuery 获取封禁用户列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/query-blocked-list
func (rc *RongCloud) UserBlockQuery(ctx context.Context, req *UserBlockQueryRequest) (*UserBlockQueryResponse, error) {
	path := "/user/block/query.json"
	resp := &UserBlockQueryResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserUnBlockRequest struct {
	// [必传] 用户 ID，支持一次解除多个用户，最多不超过 20 个。
	UserId []string `url:"userId,omitempty"`
}

type UserUnBlockResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserUnBlock 解除封禁
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/block
func (rc *RongCloud) UserUnBlock(ctx context.Context, req *UserUnBlockRequest) (*UserUnBlockResponse, error) {
	path := "/user/unblock.json"
	resp := &UserUnBlockResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserBlacklistAddRequest struct {
	// [必传] 用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 被加入黑名单的用户 ID。单次可添加最多 20 个 blackUserId。
	BlackUserId []string `url:"blackUserId,omitempty"`
}

type UserBlacklistAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserBlacklistAdd 添加用户到黑名单
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/add-to-blacklist
func (rc *RongCloud) UserBlacklistAdd(ctx context.Context, req *UserBlacklistAddRequest) (*UserBlacklistAddResponse, error) {
	path := "/user/blacklist/add.json"
	resp := &UserBlacklistAddResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserBlacklistRemoveRequest struct {
	// [必传] 用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 被移除黑名单的用户 ID，每次最多移除 20 个用户。
	BlackUserId []string `url:"blackUserId,omitempty"`
}

type UserBlacklistRemoveResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserBlacklistRemove 移除黑名单中用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/remove-from-blacklist
func (rc *RongCloud) UserBlacklistRemove(ctx context.Context, req *UserBlacklistRemoveRequest) (*UserBlacklistRemoveResponse, error) {
	path := "/user/blacklist/remove.json"
	resp := &UserBlacklistRemoveResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserBlacklistQueryRequest struct {
	// [必传] 用户 ID。
	UserId *string `url:"userId,omitempty"`
}

type UserBlacklistQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []string `json:"users"` // 黑名单用户数组。
}

// UserBlacklistQuery 获取黑名单用户列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/query-blacklist
func (rc *RongCloud) UserBlacklistQuery(ctx context.Context, req *UserBlacklistQueryRequest) (*UserBlacklistQueryResponse, error) {
	path := "/user/blacklist/query.json"
	resp := &UserBlacklistQueryResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserCheckOnlineRequest struct {
	// [必传] 用户 ID。
	UserId *string `url:"userId,omitempty"`
}

type UserCheckOnlineResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Status             string `json:"status"` // 用户状态。1 - 用户当前在线。0 - 用户当前不在线。
}

// UserCheckOnline 查询用户在线状态
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/check-online-status-by-user
func (rc *RongCloud) UserCheckOnline(ctx context.Context, req *UserCheckOnlineRequest) (*UserCheckOnlineResponse, error) {
	path := "/user/checkOnline.json"
	resp := &UserCheckOnlineResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserTagSetRequest struct {
	// [必传] 用户 ID。
	UserId *string `json:"userId,omitempty"`
	// [必传] 用户标签，一个用户最多添加 20 个标签，每个 tag 最大不能超过 40 个字节，标签中不能包含特殊字符。每次设置时需要传入用户的全量标签数据。传入空数组表示清除该用户的所有标签。
	Tags []string `json:"tags,omitempty"`
}

type UserTagSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserTagSet 设置用户标签
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/set-user-tag
func (rc *RongCloud) UserTagSet(ctx context.Context, req *UserTagSetRequest) (*UserTagSetResponse, error) {
	path := "/user/tag/set.json"
	resp := &UserTagSetResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserTagBatchSetRequest struct {
	// [必传] 用户 ID，一次最多支持 1000 个用户。传入的所有用户的标签都被会覆盖更新为 tags 中的标签。
	UserIds []string `json:"userIds,omitempty"`
	// [必传] 用户标签，一个用户最多添加 20 个标签，每个 tag 最大不能超过 40 个字节，标签中不能包含特殊字符。每次设置时需要传入用户的全量标签数据。传入空数组表示清除该用户的所有标签。
	Tags []string `json:"tags,omitempty"`
}

type UserTagBatchSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// UserTagBatchSet 批量设置用户标签
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/batch-set-user-tag
func (rc *RongCloud) UserTagBatchSet(ctx context.Context, req *UserTagBatchSetRequest) (*UserTagBatchSetResponse, error) {
	path := "/user/tag/batch/set.json"
	resp := &UserTagBatchSetResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserTagsGetRequest struct {
	// [必传] 用户 ID，一次最多支持 50 个用户。
	UserIds []string `url:"userIds,omitempty"`
}

type UserTagsGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Result             map[string][]string `json:"result"` // 用户所有的标签数组。userId:tags
	// {"code":200,"result":{"u02":["tag4","tag3"],"u01":["tag4","tag3"]}}
}

// UserTagsGet 获取用户标签
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/get-user-tag
func (rc *RongCloud) UserTagsGet(ctx context.Context, req *UserTagsGetRequest) (*UserTagsGetResponse, error) {
	path := "/user/tags/get.json"
	resp := &UserTagsGetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserDeactivateRequest struct {
	// [必传] 被注销用户 ID，最多一次 100 个。逗号分割
	UserId *string `url:"userId,omitempty"`
}

type UserDeactivateResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	OperateId          string `json:"operateId"` // 操作 ID，为当前操作的唯一标识。开通用户注销与激活状态回调后，回调请求正文中会携带此参数。
}

// UserDeactivate 注销用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/deactivate
func (rc *RongCloud) UserDeactivate(ctx context.Context, req *UserDeactivateRequest) (*UserDeactivateResponse, error) {
	path := "/user/deactivate.json"
	resp := &UserDeactivateResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserDeactivateQueryRequest struct {
	// 分页获取注销用户列表时的当前页数，默认 1，最小 1。
	PageNo *int `url:"pageNo,omitempty"`
	// 分页获取注销用户列表时的每页行数，默认 50，最小 1，最大 50。
	PageSize *int `url:"pageSize,omitempty"`
}

type UserDeactivateQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []string `json:"users"` // 已注销的用户 ID 列表
}

// UserDeactivateQuery 查询已注销用户
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/query-deactivated-list
func (rc *RongCloud) UserDeactivateQuery(ctx context.Context, req *UserDeactivateQueryRequest) (*UserDeactivateQueryResponse, error) {
	path := "/user/deactivate/query.json"
	resp := &UserDeactivateQueryResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserReactivateRequest struct {
	// [必传] 激活用户 ID，单次请求最多传入 100 个用户 ID。逗号分割
	UserId *string `url:"userId,omitempty"`
}

type UserReactivateResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	OperateId          string `json:"operateId"` // 操作 ID，为当前操作的唯一标识。开通用户注销与激活状态回调后，回调请求正文中会携带此参数。
}

// UserReactivate 重新激活用户 ID
// More details see https://doc.rongcloud.cn/imserver/server/v1/user/reactivate
func (rc *RongCloud) UserReactivate(ctx context.Context, req *UserReactivateRequest) (*UserReactivateResponse, error) {
	path := "/user/reactivate.json"
	resp := &UserReactivateResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
