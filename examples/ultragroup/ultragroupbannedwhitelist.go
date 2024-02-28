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

	// 加入超级群全体禁言白名单
	_, err := rc.UltraGroupBannedWhitelistAdd(ctx, &rongcloud.UltraGroupBannedWhitelistAddRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		UserIds:    rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group banned whitelist add error %s", err)
	}

	// 查询群组全体禁言白名单
	ugBannedWhitelistResp, err := rc.UltraGroupBannedWhitelistGet(ctx, &rongcloud.UltraGroupBannedWhitelistGetRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		Page:       rongcloud.IntPtr(1),
		PageSize:   rongcloud.IntPtr(20),
	})
	if err != nil {
		log.Fatalf("ultra group banned whitelist get error %s", err)
	}
	ugBannedWhitelistRespData, _ := json.Marshal(ugBannedWhitelistResp)
	log.Printf("ultra group whitelist resp data %s", ugBannedWhitelistRespData)

	// 移出超级群全体禁言白名单
	_, err = rc.UltraGroupBannedWhitelistDel(ctx, &rongcloud.UltraGroupBannedWhitelistDelRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		UserIds:    rongcloud.StringPtr("u01,u02"),
	})
	if err != nil {
		log.Fatalf("ultra group banned whitelist del error %s", err)
	}
}
