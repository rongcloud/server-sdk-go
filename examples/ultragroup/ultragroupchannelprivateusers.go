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

	// 添加私有频道成员
	_, err := rc.UltraGroupChannelPrivateUsersAdd(ctx, &rongcloud.UltraGroupChannelPrivateUsersAddRequest{
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

	// 删除私有频道成员
	_, err = rc.UltraGroupChannelPrivateUsersDel(ctx, &rongcloud.UltraGroupChannelPrivateUsersDelRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		UserIds:    rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group channel private users del error: %s", err)
	}
}
