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
	"os"
	"testing"
)

func TestRongCloud_EntrustGroupCreate(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupId"] = "groupId03"
	params["name"] = "name01"
	params["owner"] = "owner01"
	// params["userIds"] = "userId01,userIds02"

	result, err := rc.EntrustGroupCreate(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupProfileUpdate(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupId"] = "groupId01"
	params["name"] = "name01"

	result, err := rc.EntrustGroupProfileUpdate(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupProfileQuery(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupIds"] = "groupId01,groupId02"

	result, err := rc.EntrustGroupProfileQuery(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupQuit(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupId"] = "groupId01"
	params["userIds"] = "user01,user02"

	result, err := rc.EntrustGroupQuit(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupMemberKick(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupId"] = "groupId01"
	params["userIds"] = "user01,user02"

	result, err := rc.EntrustGroupMemberKick(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupMemberKickAll(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["userId"] = "user01"

	result, err := rc.EntrustGroupMemberKickAll(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupDismiss(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupId"] = "groupId01"

	result, err := rc.EntrustGroupDismiss(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupJoin(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupId"] = "groupId01"
	params["userIds"] = "user01,user02"

	result, err := rc.EntrustGroupJoin(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupTransferOwner(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupId"] = "groupId01"
	params["newOwner"] = "user01"

	result, err := rc.EntrustGroupTransferOwner(params)

	t.Log(result)
	t.Logf("%v", err)
}

func TestRongCloud_EntrustGroupImport(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)

	params := make(map[string]interface{})
	params["groupId"] = "u02"
	params["name"] = "name01"
	params["owner"] = "u02"

	result, err := rc.EntrustGroupImport(params)

	t.Log(result)
	t.Logf("%v", err)
}
