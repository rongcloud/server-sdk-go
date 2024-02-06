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

	// 禁言指定超级群成员
	_, err := rc.UltraGroupUserBannedAdd(ctx, &rongcloud.UltraGroupUserBannedAddRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		UserIds:    rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group user banned add error %s", err)
	}

	// 查询超级群成员禁言列表
	ultraGroupUserBannedResp, err := rc.UltraGroupUserBannedGet(ctx, &rongcloud.UltraGroupUserBannedGetRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		Page:       rongcloud.IntPtr(1),
		PageSize:   rongcloud.IntPtr(20),
	})
	if err != nil {
		log.Fatalf("ultra group user banned get error %s", err)
	}
	ultraGroupUserBannedRespData, _ := json.Marshal(ultraGroupUserBannedResp)
	log.Printf("ultra group user banned resp data: %s", ultraGroupUserBannedRespData)

	// 取消指定超级群成员禁言
	_, err = rc.UltraGroupUserBannedDel(ctx, &rongcloud.UltraGroupUserBannedDelRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		UserIds:    rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group user banned del error %s", err)
	}
}
