package rongcloud

import "context"

type SensitiveWordAddRequest struct {
	// [必传] 敏感词，最长不超过 32 个字符，格式为汉字、数字、字母。
	Word *string `url:"word,omitempty"`
	// 替换后的词，最长不超过 32 个字符。如未设置，当消息中含有敏感词时，消息将被屏蔽，用户不会收到消息。如设置了，当消息中含有敏感词时，将被替换为指定的词进行发送。
	ReplaceWord *string `url:"replaceWord,omitempty"`
}

type SensitiveWordAddResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// SensitiveWordAdd 添加消息敏感词
// More details see https://doc.rongcloud.cn/imserver/server/v1/moderation/add-sensitive-word
func (rc *RongCloud) SensitiveWordAdd(ctx context.Context, req *SensitiveWordAddRequest) (*SensitiveWordAddResponse, error) {
	path := "/sensitiveword/add.json"
	resp := &SensitiveWordAddResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type SensitiveWordListRequest struct {
	// 查询敏感词的类型。0 为查询替换敏感词。1 为查询屏蔽敏感词。2 为查询全部敏感词。默认为 1。
	Type *string `url:"type,omitempty"`
}

type SensitiveWordListResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Words              []*SensitiveWord `json:"words"`
}

type SensitiveWord struct {
	Word        string `json:"word"`        // 敏感词内容。
	ReplaceWord string `json:"replaceWord"` // 替换敏感词的内容，为空时对应 Word 敏感词类型为屏蔽敏感词。
	Type        string `json:"type"`        // 0 为替换敏感词。1 为屏蔽敏感词。
}

// SensitiveWordList 查询消息敏感词
// More details see https://doc.rongcloud.cn/imserver/server/v1/moderation/query-sensitive-word
func (rc *RongCloud) SensitiveWordList(ctx context.Context, req *SensitiveWordListRequest) (*SensitiveWordListResponse, error) {
	path := "/sensitiveword/list.json"
	resp := &SensitiveWordListResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type SensitiveWordBatchDeleteRequest struct {
	// [必传] 敏感词数组，一次最多移除 50 个敏感词。
	Words []string `url:"words,omitempty"`
}

type SensitiveWordBatchDeleteResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
}

// SensitiveWordBatchDelete 批量移除消息敏感词
// More details see https://doc.rongcloud.cn/imserver/server/v1/moderation/batch-remove-sensitive-word
func (rc *RongCloud) SensitiveWordBatchDelete(ctx context.Context, req *SensitiveWordBatchDeleteRequest) (*SensitiveWordBatchDeleteResponse, error) {
	path := "/sensitiveword/batch/delete.json"
	resp := &SensitiveWordBatchDeleteResponse{}
	httpResp, err := rc.postForm(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
