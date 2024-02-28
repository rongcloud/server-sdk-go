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

	// 添加私有频道成员
	_, err = rc.UltraGroupChannelPrivateUsersAdd(ctx, &rongcloud.UltraGroupChannelPrivateUsersAddRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		UserIds:    rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group channel private users add error: %s", err)
	}

	// 查询私有频道成员列表
	usersQueryResp, err := rc.UltraGroupChannelPrivateUsersGet(ctx, &rongcloud.UltraGroupChannelPrivateUsersGetRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		Page:       rongcloud.IntPtr(1),
		PageSize:   rongcloud.IntPtr(20),
	})
	if err != nil {
		log.Fatalf("ultra group channel private users get error %s", err)
	}
	usersQueryRespData, _ := json.Marshal(usersQueryResp)
	log.Printf("ultra group channel private users get data %s", usersQueryRespData)

	// 查询用户所属的私有频道
	userChannelQueryResp, err := rc.UltraGroupUserChannelQuery(ctx, &rongcloud.UltraGroupUserChannelQueryRequest{
		GroupId:  rongcloud.StringPtr("ug01"),
		UserId:   rongcloud.StringPtr("u01"),
		Page:     rongcloud.IntPtr(1),
		PageSize: rongcloud.IntPtr(10),
	})
	if err != nil {
		log.Fatalf("ultra group user channel query error %s", err)
	}
	userChannelQueryRespData, _ := json.Marshal(userChannelQueryResp)
	log.Printf("ultra group user channel query data %s", userChannelQueryRespData)

	// 删除私有频道成员
	_, err = rc.UltraGroupChannelPrivateUsersDel(ctx, &rongcloud.UltraGroupChannelPrivateUsersDelRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		UserIds:    rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group channel private users del error: %s", err)
	}

	// 解散超级群
	_, err = rc.UltraGroupDis(ctx, &rongcloud.UltraGroupDisRequest{GroupId: rongcloud.StringPtr("ug01")})
	if err != nil {
		log.Fatalf("ultra group dis error %s", err)
	}
}
