package rongcloud

import "context"

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
	UserId *string `url:"userId,omitempty"`
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
