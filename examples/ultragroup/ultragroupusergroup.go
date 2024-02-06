package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	// 创建用户组
	_, err := rc.UltraGroupUserGroupAdd(ctx, &rongcloud.UltraGroupUserGroupAddRequest{
		GroupId: rongcloud.StringPtr("ug01"),
		UserGroups: []*rongcloud.UltraGroupUserGroup{
			{
				UserGroupId: "usg01",
			}, {
				UserGroupId: "usg02",
			},
		},
	})
	if err != nil {
		log.Fatalf("ultra group user group add error %s", err)
	}

	// 查询用户组列表
	userGroupQueryResp, err := rc.UltraGroupUserGroupQuery(ctx, &rongcloud.UltraGroupUserGroupQueryRequest{
		GroupId:  rongcloud.StringPtr("ug01"),
		Page:     rongcloud.IntPtr(1),
		PageSize: rongcloud.IntPtr(10),
	})
	if err != nil {
		log.Fatalf("ultra group user group query error %s", err)
	}
	userGroupQueryRespData, _ := json.Marshal(userGroupQueryResp)
	log.Printf("ultra group user query resp data: %s", userGroupQueryRespData)

	// 删除用户组
	_, err = rc.UltraGroupUserGroupDel(ctx, &rongcloud.UltraGroupUserGroupDelRequest{
		GroupId:      rongcloud.StringPtr("ug01"),
		UserGroupIds: rongcloud.StringPtr("usg01,usg02"),
	})
	if err != nil {
		log.Fatalf("ultra group user group del error %s", err)
	}
}
