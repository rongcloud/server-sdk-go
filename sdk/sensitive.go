package sdk

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/httplib"
)

// ListWordFilterResult listWordFilter返回结果
type ListWordFilterResult struct {
	Words []SensitiveWord `json:"words"`
}

// SensitiveWord 敏感词
// Type 0: 敏感词替换 1: 敏感词屏蔽
type SensitiveWord struct {
	Type        string `json:"type"`
	Word        string `json:"word"`
	ReplaceWord string `json:"replaceWord"`
}

// SensitiveAdd 添加敏感词
/*
*@param  keyword:敏感词，最长不超过 32 个字符，格式为汉字、数字、字母
*@param  replace:敏感词替换，最长不超过 32 个字符， 敏感词屏蔽可以为空
*@param  sensitiveType:0: 敏感词替换 1: 敏感词屏蔽
*
*@return error
 */
func (rc *RongCloud) SensitiveAdd(keyword, replace string, sensitiveType int) error {
	if keyword == "" {
		return RCErrorNew(1002, "Paramer 'keyword' is required")
	}
	if replace == "" {
		return RCErrorNew(1002, "Paramer 'replace' is required")
	}
	req := httplib.Post(rc.rongCloudURI + "/sensitiveword/add." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Param("word", keyword)
	switch sensitiveType {
	case 0:
		req.Param("replaceWord", replace)
	case 1:

	default:
		return RCErrorNew(1002, "Paramer 'replace' is required")
	}

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err
}

// SensitivePost结构体,replaceWord为可选字段
type SensitivePost struct {
	Word        string `json:"word"`
	ReplaceWord string `json:"replaceWord,,omitempty"`
}

// SensitivePost数组结构体
type SensitivePostArray struct {
	Words []SensitivePost `json:"words"`
}

type SensitiveBatchAddResult struct {
	Code        int `json:"code"`
	RemainQuota int `json:"remainQuota"`
}

// SensitiveBatchAdd 批量添加敏感词
/*
*@param  words:敏感词数组
*
*@return error
 */
func (rc *RongCloud) SensitiveBatchAdd(words []SensitiveWord) (SensitiveBatchAddResult, error) {

	// 遍历SensitiveWord数组转换成SensitivePost数组
	var sensitivePostArray SensitivePostArray
	for _, v := range words {
		sensitivePostArray.Words = append(sensitivePostArray.Words, SensitivePost{
			Word:        v.Word,
			ReplaceWord: v.ReplaceWord,
		})
	}

	sensitivePostJson, err := json.Marshal(sensitivePostArray)
	if err != nil {
		return SensitiveBatchAddResult{}, RCErrorNew(1002, "Marshal 'sensitivePostJson' err")
	}

	req := httplib.Post(rc.rongCloudURI + "/sensitiveword/batch/add." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	req.Body(sensitivePostJson)
	req.Header("Content-Type", "application/json")

	res, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return SensitiveBatchAddResult{}, err
	}
	var ret SensitiveBatchAddResult
	if err := json.Unmarshal(res, &ret); err != nil {
		return SensitiveBatchAddResult{}, err
	}

	return ret, err
}

// SensitiveGetList 查询敏感词列表方法
/*
*@return ListWordFilterResult error
 */
func (rc *RongCloud) SensitiveGetList() (ListWordFilterResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/sensitiveword/list." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	resp, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
		return ListWordFilterResult{}, err
	}

	var ret ListWordFilterResult
	if err := json.Unmarshal(resp, &ret); err != nil {
		return ListWordFilterResult{}, err
	}
	return ret, err

}

// SensitiveRemove 移除敏感词方法（从敏感词列表中，移除某一敏感词。）
/*
*@param  keywords:每次最多删除 50 个敏感词，2 小时后生效
*
*@return error
 */
func (rc *RongCloud) SensitiveRemove(keywords []string) error {
	if len(keywords) == 0 {
		return RCErrorNew(1002, "Paramer 'keywords' is required")
	}

	req := httplib.Post(rc.rongCloudURI + "/sensitiveword/batch/delete." + ReqType)
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)
	for _, v := range keywords {
		req.Param("words", v)
	}

	_, err := rc.do(req)
	if err != nil {
		rc.urlError(err)
	}
	return err

}
