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

	// 加入超级群
	_, err = rc.UltraGroupJoin(ctx, &rongcloud.UltraGroupJoinRequest{
		UserId:  rongcloud.StringPtr("u02"),
		GroupId: rongcloud.StringPtr("ug01"),
	})

	// 查询用户是否为群成员
	ultraGroupMemberExistResp, err := rc.UltraGroupMemberExist(ctx, &rongcloud.UltraGroupMemberExistRequest{
		UserId:  rongcloud.StringPtr("u01"),
		GroupId: rongcloud.StringPtr("ug01"),
	})
	if err != nil {
		log.Fatalf("ultra group member exist error %s", err)
	}
	ultraGroupMemberExistRespData, _ := json.Marshal(ultraGroupMemberExistResp)
	log.Printf("ultra group member exist resp data %s", ultraGroupMemberExistRespData)

	// 退出超级群
	_, err = rc.UltraGroupQuit(ctx, &rongcloud.UltraGroupQuitRequest{
		UserId:  rongcloud.StringPtr("u02"),
		GroupId: rongcloud.StringPtr("ug01"),
	})
	if err != nil {
		log.Fatalf("ultra group quit error %s", err)
	}

	// 刷新超级群信息
	_, err = rc.UltraGroupRefresh(ctx, &rongcloud.UltraGroupRefreshRequest{
		GroupId:   rongcloud.StringPtr("ug01"),
		GroupName: rongcloud.StringPtr("ug01"),
	})
	if err != nil {
		log.Fatalf("ultra group refresh error %s", err)
	}

	// 解散超级群
	_, err = rc.UltraGroupDis(ctx, &rongcloud.UltraGroupDisRequest{GroupId: rongcloud.StringPtr("ug01")})
	if err != nil {
		log.Fatalf("ultra group dis error %s", err)
	}
}
