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

	// 创建超级群
	_, err := rc.UltraGroupCreate(ctx, &rongcloud.UltraGroupCreateRequest{
		UserId:    rongcloud.StringPtr("u01"),
		GroupId:   rongcloud.StringPtr("ug01"),
		GroupName: rongcloud.StringPtr("ug01"),
	})
	if err != nil {
		log.Fatalf("ultra group create error %s", err)
	}

	// 超级群创建用户组
	_, err = rc.UltraGroupUserGroupAdd(ctx, &rongcloud.UltraGroupUserGroupAddRequest{
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

	// 超级群用户组添加用户
	_, err = rc.UltraGroupUserGroupUserAdd(ctx, &rongcloud.UltraGroupUserGroupUserAddRequest{
		GroupId:     rongcloud.StringPtr("ug01"),
		UserGroupId: rongcloud.StringPtr("usg01"),
		UserIds:     rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group user group user add error %s", err)
	}

	// 超级群查询用户所属用户组
	userUserGroupQueryResp, err := rc.UltraGroupUserUserGroupQuery(ctx, &rongcloud.UltraGroupUserUserGroupQueryRequest{
		GroupId:  rongcloud.StringPtr("ug01"),
		UserId:   rongcloud.StringPtr("u01"),
		Page:     rongcloud.IntPtr(1),
		PageSize: rongcloud.IntPtr(10),
	})
	if err != nil {
		log.Fatalf("ultra group user user group query error %s", err)
	}
	userUserGroupQueryRespData, _ := json.Marshal(userUserGroupQueryResp)
	log.Printf("user user group query resp data: %s", userUserGroupQueryRespData)

	// 超级群用户组移出用户
	_, err = rc.UltraGroupUserGroupUserDel(ctx, &rongcloud.UltraGroupUserGroupUserDelRequest{
		GroupId:     rongcloud.StringPtr("ug01"),
		UserGroupId: rongcloud.StringPtr("usg01"),
		UserIds:     rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group user group user del error %s", err)
	}
}
