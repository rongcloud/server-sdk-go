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

	// 查询用户所在群组
	userGroupQueryResp, err := rc.UserGroupQuery(ctx, &rongcloud.UserGroupQueryRequest{
		UserId: rongcloud.StringPtr("u01"),
		Page:   rongcloud.IntPtr(1),
		Size:   rongcloud.IntPtr(10),
	})
	if err != nil {
		log.Fatalf("user group query error %s", err)
	}
	userGroupQueryRespData, _ := json.Marshal(userGroupQueryResp)
	log.Printf("user group query resp data %s", userGroupQueryRespData)

	// 同步用户所在群组
	_, err = rc.GroupSync(ctx, &rongcloud.GroupSyncRequest{
		UserId: rongcloud.StringPtr("u01"),
		Groups: rongcloud.SyncGroups{
			{
				Id:   "grp01",
				Name: "grp01",
			}, {
				Id:   "grp02",
				Name: "grp02",
			},
		},
	})
	if err != nil {
		log.Fatalf("group sync error %s", err)
	}

	// 刷新群组信息
	_, err = rc.GroupRefresh(ctx, &rongcloud.GroupRefreshRequest{
		GroupId:   rongcloud.StringPtr("grp01"),
		GroupName: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group refresh error %s", err)
	}
}
