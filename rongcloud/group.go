package rongcloud

import (
	"context"
	"fmt"
	"net/url"
)

type GroupRemarksSetRequest struct {
	// [必传] 群成员用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 群 ID。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 群成员推送备注。
	Remark *string `url:"remark,omitempty"`
}

type GroupRemarksSetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// GroupRemarksSet 设置群成员推送备注名
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/set-remark-for-group-push
func (rc *RongCloud) GroupRemarksSet(ctx context.Context, req *GroupRemarksSetRequest) (*GroupRemarksSetResponse, error) {
	path := "/group/remarks/set.json"
	resp := &GroupRemarksSetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type GroupRemarksDelRequest struct {
	// [必传] 群成员用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 群 ID。
	GroupId *string `url:"groupId,omitempty"`
}

type GroupRemarksDelResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// GroupRemarksDel 删除群成员推送备注名
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/delete-remark-for-group-push
func (rc *RongCloud) GroupRemarksDel(ctx context.Context, req *GroupRemarksDelRequest) (*GroupRemarksDelResponse, error) {
	path := "/group/remarks/del.json"
	resp := &GroupRemarksDelResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type GroupRemarksGetRequest struct {
	// [必传] 群成员用户 ID。
	UserId *string `url:"userId,omitempty"`
	// [必传] 群 ID。
	GroupId *string `url:"groupId,omitempty"`
}

type GroupRemarksGetResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Remark             string `json:"remark"` // 备注名称。
}

// GroupRemarksGet 查询群成员推送备注名
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/get-remark-for-group-push
func (rc *RongCloud) GroupRemarksGet(ctx context.Context, req *GroupRemarksGetRequest) (*GroupRemarksGetResponse, error) {
	path := "/group/remarks/get.json"
	resp := &GroupRemarksGetResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type GroupCreateRequest struct {
	// [必传] 要加入群的用户 ID，最多不超过 1000 个。
	UserId []string `url:"userId,omitempty"`
	// [必传] 创建群组 ID，最大长度 64 个字符。支持大小写英文字母与数字的组合。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 群组 ID 对应的名称，用于在发送群组消息显示远程 Push 通知时使用，如群组名称改变需要调用刷新群组信息接口同步。
	GroupName *string `url:"groupName,omitempty"`
}

type GroupCreateResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// GroupCreate 创建群组
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/create-group
func (rc *RongCloud) GroupCreate(ctx context.Context, req *GroupCreateRequest) (*GroupCreateResponse, error) {
	path := "/group/create.json"
	resp := &GroupCreateResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type GroupUserGagAddRequest struct {
	// [必传] 用户 ID，每次添加最多不超过 20 个用户。
	UserId []string `url:"userId,omitempty"`
	// 群组 ID，为空时则设置用户在加入的所有群组中都不能发送消息。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 禁言时长，以分钟为单位，最大值为 43200 分钟，为 0 表示永久禁言。
	Minute *int `url:"minute,omitempty"`
}

type GroupUserGagAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// GroupUserGagAdd 禁言指定群成员
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/gag-user
func (rc *RongCloud) GroupUserGagAdd(ctx context.Context, req *GroupUserGagAddRequest) (*GroupUserGagAddResponse, error) {
	path := "/group/user/gag/add.json"
	resp := &GroupUserGagAddResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type GroupUserGagListRequest struct {
	// 群组 ID，为空时则设置用户在加入的所有群组中都不能发送消息。
	GroupId *string `url:"groupId,omitempty"`
}

type GroupUserGagListResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Users              []*GagUser `json:"users"` // 禁言成员列表。
}

type GagUser struct {
	UserId string `json:"userId"` // 群成员 Id。
	Time   string `json:"time"`   // 解禁时间。精确到秒，格式为 YYYY-MM-DD HH:MM:SS，例如 2022-09-25 16:12:38。注意：time 的值与应用所属数据中心有关。如您的 App 业务使用国内数据中心，则 time 为北京时间。如您的 App 业务使用海外数据中心，则 time 为 UTC 时间。
}

// GroupUserGagList 查询群成员禁言列表
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/query-gagged-user
func (rc *RongCloud) GroupUserGagList(ctx context.Context, req *GroupUserGagListRequest) (*GroupUserGagListResponse, error) {
	path := "/group/user/gag/list.json"
	resp := &GroupUserGagListResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type GroupUserGagRollbackRequest struct {
	// [必传] 用户 ID，每次最多移除 20 个用户。
	UserId []string `url:"userId,omitempty"`
	// 群组 ID，为空时则移除用户在所有群组中的禁言设置。
	GroupId *string `url:"groupId,omitempty"`
}

type GroupUserGagRollbackResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// GroupUserGagRollback 取消指定群成员禁言
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/ungag-user
func (rc *RongCloud) GroupUserGagRollback(ctx context.Context, req *GroupUserGagRollbackRequest) (*GroupUserGagRollbackResponse, error) {
	path := "/group/user/gag/rollback.json"
	resp := &GroupUserGagRollbackResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type UserGroupQueryRequest struct {
	// [必传] 用户 ID。
	UserId *string `url:"userId,omitempty"`
	// 当前页数，在分页查询时使用。如果进行分页查询，页面大小默认为 50，可使用 size 调整页面大小。如无需分页可不传（或传 0），可获得用户加入的前 5000 个群组列表。
	Page *int `url:"page,omitempty"`
	// 页面大小，仅在 page 传入有效值时生效。默认每页 50 行，最大值 1000。
	Size *int `url:"size,omitempty"`
}

type UserGroupQueryResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Groups             []*Group `json:"groups"` // 用户加入的群信息数组。
}

type Group struct {
	Id   string `json:"id"`   // 群组 ID。
	Name string `json:"name"` // 群名称。
}

type SyncGroups []Group

func (r SyncGroups) EncodeValues(key string, v *url.Values) error {
	for _, grp := range r {
		v.Set(fmt.Sprintf("%s[%s]", key, grp.Id), grp.Name)
	}
	return nil
}

// UserGroupQuery 查询用户所在群组
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/query-group-by-user
func (rc *RongCloud) UserGroupQuery(ctx context.Context, req *UserGroupQueryRequest) (*UserGroupQueryResponse, error) {
	path := "/user/group/query.json"
	resp := &UserGroupQueryResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type GroupSyncRequest struct {
	// [必传] 被同步群信息的用户 ID。
	UserId *string `url:"userId,omitempty"`
	// 该用户所属的群信息，如群组 ID 已经存在，则同时刷新对应群组名称。此参数可传多个，参见下面示例。
	Groups SyncGroups `url:"group,omitempty"`
}

type GroupSyncResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// GroupSync 同步用户所在群组
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/sync-group
func (rc *RongCloud) GroupSync(ctx context.Context, req *GroupSyncRequest) (*GroupSyncResponse, error) {
	path := "/group/sync.json"
	resp := &GroupSyncResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type GroupRefreshRequest struct {
	// [必传] 群组 ID。
	GroupId *string `url:"groupId,omitempty"`
	// [必传] 群组名称。
	GroupName *string `url:"groupName,omitempty"`
}

type GroupRefreshResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// GroupRefresh 刷新群组信息
// More details see https://doc.rongcloud.cn/imserver/server/v1/group/refresh-group-info
func (rc *RongCloud) GroupRefresh(ctx context.Context, req *GroupRefreshRequest) (*GroupRefreshResponse, error) {
	path := "/group/refresh.json"
	resp := &GroupRefreshResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
