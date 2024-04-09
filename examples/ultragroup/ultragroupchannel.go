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

	// 创建频道
	_, err := rc.UltraGroupChannelCreate(ctx, &rongcloud.UltraGroupChannelCreateRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		Type:       rongcloud.IntPtr(0),
	})
	if err != nil {
		log.Fatalf("ultra group channel create error %s", err)
	}

	// 查询频道列表
	ultraGroupChannelResp, err := rc.UltraGroupChannelGet(ctx, &rongcloud.UltraGroupChannelGetRequest{
		GroupId: rongcloud.StringPtr("ug01"),
		Page:    rongcloud.IntPtr(1),
		Limit:   rongcloud.IntPtr(20),
	})
	if err != nil {
		log.Fatalf("ultra group channel get error %s", err)
	}
	ultraGroupChannelRespData, _ := json.Marshal(ultraGroupChannelResp)
	log.Printf("ultra group channel resp data: %s", ultraGroupChannelRespData)

	// 变更频道类型
	_, err = rc.UltraGroupChannelTypeChange(ctx, &rongcloud.UltraGroupChannelTypeChangeRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		Type:       rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("ultra group channel type change error %s", err)
	}

	// 删除频道
	_, err = rc.UltraGroupChannelDel(ctx, &rongcloud.UltraGroupChannelDelRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
	})
	if err != nil {
		log.Fatalf("ultra group channel del error %s", err)
	}
}
