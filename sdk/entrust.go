/*
 * @Descripttion:
 * @version:
 * @Author: xujie
 * @Date: 2025-01-09
 * @LastEditors: xujie
 * @LastEditTime: 2025-01-09
 */
package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

type EntrustBaseResult struct {
	// 返回码，200 为正常。
	Code int `json:"code"`
}

// EntrustGroupCreate 创建群组
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupCreate(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/create.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err

}

// EntrustGroupProfileUpdate 群组资料设置
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupProfileUpdate(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/profile/update.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err

}

type ProfileResult struct {
	GroupId         string `json:"groupId"`
	Name            string `json:"name"`
	GroupProfile    string `json:"groupProfile"`
	GroupExtProfile string `json:"groupExtProfile"`
	Permissions     string `json:"permissions"`
	Owner           string `json:"owner"`
	CreateTime      int    `json:"createTime"`
	MemberCount     int    `json:"memberCount"`
}

type EntrustGroupProfileQueryResult struct {
	Code         int             `json:"code"`
	ProfileArray []ProfileResult `json:"profiles"`
}

// EntrustGroupProfileQuery 群组资料查询
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupProfileQuery(params map[string]interface{}) (EntrustGroupProfileQueryResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/profile/query.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustGroupProfileQueryResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err

}

// EntrustGroupQuit 退出群组
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupQuit(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/quit.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}

// EntrustGroupMemberKick 踢出群组
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupMemberKick(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/member/kick.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}

// EntrustGroupMemberKickAll 踢出所有群组
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupMemberKickAll(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/member/kick/all.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}

// EntrustGroupDismiss 解散群组
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupDismiss(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/dismiss.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}

// EntrustGroupJoin 加入群组
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupJoin(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/join.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}

// EntrustGroupTransferOwner 转移群组
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupTransferOwner(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/transfer/owner.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}

// EntrustGroupImport 导入群组
// *
// @param: params map形式，value支持string，int
//
// *//
func (rc *RongCloud) EntrustGroupImport(params map[string]interface{}) (EntrustBaseResult, error) {

	req := httplib.Post(rc.rongCloudURI + "/entrust/group/import.json")
	req.SetTimeout(time.Second*rc.timeout, time.Second*rc.timeout)
	rc.fillHeader(req)

	result := EntrustBaseResult{}

	for key, value := range params {

		switch v := value.(type) {
		case string:
			req.Param(key, v)
		case int:
			req.Param(key, strconv.Itoa(v))
		default:
			formatted := fmt.Sprintf("unsupport type : %s", key)
			return result, errors.New(formatted)
		}
	}

	res, err := rc.do(req)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	return result, err
}
