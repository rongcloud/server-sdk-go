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

	// 创建群组
	_, err := rc.GroupCreate(ctx, &rongcloud.GroupCreateRequest{
		UserId:    []string{"u01", "u02"},
		GroupId:   rongcloud.StringPtr("grp01"),
		GroupName: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group create error %s", err)
	}

	// 加入群组
	_, err = rc.GroupJoin(ctx, &rongcloud.GroupJoinRequest{
		UserIds:   []string{"u01", "u02"},
		GroupId:   rongcloud.StringPtr("grp01"),
		GroupName: nil,
	})
	if err != nil {
		log.Fatalf("group join error %s", err)
	}

	// 查询群组成员
	groupUserQueryResp, err := rc.GroupUserQuery(ctx, &rongcloud.GroupUserQueryRequest{
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group user query error: %s", err)
	}
	groupUserQueryRespData, _ := json.Marshal(groupUserQueryResp)
	log.Printf("group user query response data: %s", groupUserQueryRespData)

	// 退出群组
	_, err = rc.GroupQuit(ctx, &rongcloud.GroupQuitRequest{
		UserId:  []string{"u01", "u02"},
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group quit error %s", err)
	}

	// 解散群组
	_, err = rc.GroupDismiss(ctx, &rongcloud.GroupDismissRequest{
		UserId:  rongcloud.StringPtr("u01"),
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group dismiss error %s", err)
	}
}
