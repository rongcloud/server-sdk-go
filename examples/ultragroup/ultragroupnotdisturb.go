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

	// 设置群/频道默认免打扰
	_, err := rc.UltraGroupNotDisturbSet(ctx, &rongcloud.UltraGroupNotDisturbSetRequest{
		GroupId:     rongcloud.StringPtr("ug01"),
		BusChannel:  rongcloud.StringPtr("channel01"),
		UnpushLevel: rongcloud.IntPtr(-1),
	})
	if err != nil {
		log.Fatalf("ultra group not disturb set error %s", err)
	}

	// 查询默认免打扰配置
	ultraGroupNotDisturbGetResp, err := rc.UltraGroupNotDisturbGet(ctx, &rongcloud.UltraGroupNotDisturbGetRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
	})
	if err != nil {
		log.Fatalf("ultra group not disturb get error %s", err)
	}
	ultraGroupNotDisturbGetRespData, _ := json.Marshal(ultraGroupNotDisturbGetResp)
	log.Printf("ultra group not disturb get resp data: %s", ultraGroupNotDisturbGetRespData)
}
