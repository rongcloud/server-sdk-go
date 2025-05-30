/*
 * @Descripttion:
 * @version:
 * @Author: ran.ding
 * @Date: 2019-09-02 18:29:55
 * @LastEditors: ran.ding
 * @LastEditTime: 2019-09-10 12:14:41
 */
package sdk

import (
	"fmt"
	"os"
	"testing"
)

func TestRongCloud_UserBlockPushPeriodDelete(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if err := rc.UserBlockPushPeriodDelete("u01"); err != nil {
		t.Errorf("UserTokenExpire err:%v", err)
		return
	}
	t.Log("suc")
}

func TestRongCloud_UserInfoGet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UserInfoGet("uid1"); err != nil {
		t.Errorf("UserInfoGet err:%v", err)
		return
	} else {
		t.Log(fmt.Sprintf("suc %+v", res))
	}
}

func TestRongCloud_UserBlockPushPeriodGet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UserBlockPushPeriodGet("u01"); err != nil {
		t.Errorf("UserTokenExpire err:%v", err)
		return
	} else {
		t.Log(fmt.Sprintf("suc %+v", res))
	}
}

func TestRongCloud_UserBlockPushPeriodSet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if err := rc.UserBlockPushPeriodSet("u01", "23:59:59", "120", ""); err != nil {
		t.Errorf("UserTokenExpire err:%v", err)
		return
	}
	t.Log("suc")
}

func TestRongCloud_UserTokenExpireResObj(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UserTokenExpireResObj("u01", 1619469955344); err != nil {
		t.Errorf("UserTokenExpire err:%v", err)
		return
	} else {
		t.Logf("user remark get suc:%+v", res)
	}
}

func TestRongCloud_UserTokenExpire(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UserTokenExpire("u01", 1619469955344); err != nil {
		t.Errorf("UserTokenExpire err:%v", err)
		return
	} else {
		t.Log("user remark get suc", string(res))
	}
}

func TestRongCloud_UserRemarksGetResObj(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UserRemarksGetResObj("u01", 1, 1); err != nil {
		t.Errorf("user remark get err:%v", err)
		return
	} else {
		t.Logf("user remark get suc :%+v", res)
	}
}

func TestRongCloud_UserRemarksGet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UserRemarksGet("u01", 1, 1); err != nil {
		t.Errorf("user remark get err:%v", err)
		return
	} else {
		t.Log("user remark get suc", string(res))
	}
}

func TestRongCloud_UserRemarksDel(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if err := rc.UserRemarksDel("u01", "qq"); err != nil {
		t.Errorf("user remark del err:%v", err)
		return
	}
	t.Log("user remark del suc")
}

func TestRongCloud_UserRemarksSet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if err := rc.UserRemarksSet("u01", []UserRemark{{
		Id:     "u01",
		Remark: "remark1",
	}}); err != nil {
		t.Errorf("user remark  set err:%v", err)
		return
	}
	t.Log("user remark set suc")
}

func TestRongCloud_UserChatFbQueryListResObj(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UserChatFbQueryListResObj(0, 0, "PERSON"); err != nil {
		t.Errorf("user chat fb set err:%v", err)
		return
	} else {
		t.Logf("user chat fb set suc:%+v", res)
	}
}

func TestRongCloud_UserChatFbQueryList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UserChatFbQueryList(0, 0, "PERSON"); err != nil {
		t.Errorf("user chat fb set err:%v", err)
		return
	} else {
		t.Log("user chat fb set suc", string(res))
	}
}

func TestRongCloud_UserChatFbSet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if err := rc.UserChatFbSet("u01", 0, "PERSON"); err != nil {
		t.Errorf("user chat fb set err:%v", err)
		return
	}
	t.Log("user chat fb set suc")
}

func TestQueryWhiteList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	_, err := rc.QueryWhiteList("123")
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		t.Log("PASS")
	}
}

func TestRemoveWhiteList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	if err := rc.RemoveWhiteList("123", []string{"234", "456"}); err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		t.Log("PASS")
	}
}

func TestAddWhiteList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	if err := rc.AddWhiteList("123", []string{"234", "345"}); err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		t.Log("PASS")
	}
}

func TestRongCloud_UserRegister(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	rep, err := rc.UserRegister(
		"1000014",
		"1000014",
		"https://img.theplace.cool/MTY1NjMxNDExNjQxMCMzMjAjcG5n.png&apiId=1",
	)
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_UserUpdate(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	rep := rc.UserUpdate("7Szq13MKRVortoknTAk7W8", "7Szq13MKRVortoknTAk7W8", "http://rongcloud.cn/portrait.jpg")
	t.Log(rep)
}

func TestRongCloud_BlockAdd(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	err := rc.BlockAdd(
		"4kIvGJmETlYqDoVFgWdYdM",
		5,
	)
	t.Log(err)
}

func TestRongCloud_BlockGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	rep, err := rc.BlockGetList()
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_BlockRemove(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := rc.BlockRemove(
		"4kIvGJmETlYqDoVFgWdYdM",
	)
	t.Log(err)
}

func TestRongCloud_BlacklistAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := rc.BlacklistAdd(
		"4kIvGJmETlYqDoVFgWdYdM",
		[]string{"u01"},
	)

	t.Log(err)
}

func TestRongCloud_BlacklistGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	req, err := rc.BlacklistGet(
		"4kIvGJmETlYqDoVFgWdYdM",
	)
	t.Log(err)
	t.Log(req)
}

func TestRongCloud_BlacklistRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := rc.BlacklistRemove(
		"4kIvGJmETlYqDoVFgWdYdM",
		[]string{"u01"},
	)
	t.Log(err)
}

func TestRongCloud_OnlineStatusCheck(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	status, err := rc.OnlineStatusCheck("4kIvGJmETlYqDoVFgWdYdM")
	t.Log(err)
	t.Log(status)
}

func TestRongCloud_TagSet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	err := rc.TagSet(
		Tag{
			UserID: "u01",
			Tags: []string{
				"Male",
			},
		})
	t.Log(err)
}

func TestRongCloud_TagBatchSet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	err := rc.TagBatchSet(
		TagBatch{
			UserIDs: []string{
				"u02",
				"u03",
			},
			Tags: []string{
				"Male",
				"bj",
			},
		})
	t.Log(err)
}

func TestRongCloud_TagGet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	result, err := rc.TagGet(
		[]string{
			"u01",
			"u02",
			"u03",
		},
	)
	if err == nil {
		for k, v := range result.Result {
			t.Log(k)
			for _, tag := range v {
				t.Log(tag)
			}
		}
	}
	t.Log(err)
}

func TestRongCloud_UserDeactivate(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	// Register user test
	result, err := rc.UserDeactivate(
		[]string{
			"u01",
			"u02",
		})
	if err != nil {
		t.Fatalf("UserDeactivate fail: %s", err)
	}
	t.Logf("res: %+v", result)
}

func TestRongCloud_UserDeactivateQuery(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	result, err := rc.UserDeactivateQuery(1, 50)
	if err != nil {
		t.Fatalf("UserDeactivateQuery fail: %s", err)
	}
	t.Logf("res: %+v", result)
}

func TestRongCloud_UserReactivate(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	result, err := rc.UserReactivate(
		[]string{
			"u01",
			"u02",
		})
	if err != nil {
		t.Fatalf("UserReactivate fail: %s", err)
	}
	t.Logf("res: %+v", result)
}

func TestRongCloud_UserProfileSet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := rc.UserProfileSet(
		"u01",
		"{\"birthday\":\"20011221\",\"level\":2}",
		"{\"ext_1\":\"testext\"}",
	)
	if err != nil {
		t.Fatalf("UserProfileSet fail: %s", err)
	}
	t.Logf("UserProfileSet: end")
}

func TestRongCloud_UserProfileClean(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := rc.UserProfileClean(
		"u01",
	)
	if err != nil {
		t.Fatalf("UserProfileClean fail: %s", err)
	}
	t.Logf("UserProfileClean: end")
}

func TestRongCloud_UserProfilQuery(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	result, err := rc.UserProfilQuery(
		1,
		20,
		0,
	)
	if err != nil {
		t.Fatalf("UserProfilQuery fail: %s", err)
	}
	t.Logf("res: %+v", result)
}

func TestRongCloud_UserProfilBatchQuery(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	result, err := rc.UserProfilBatchQuery(
		"rong_health_check.u58",
	)
	if err != nil {
		t.Fatalf("UserProfilBatchQuery fail: %s", err)
	}
	t.Logf("res: %+v", result)
}
