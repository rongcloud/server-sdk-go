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

	// 设置超级群全体禁言
	_, err := rc.UltraGroupGlobalBannedSet(ctx, &rongcloud.UltraGroupGlobalBannedSetRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		Status:     rongcloud.BoolPtr(true),
	})
	if err != nil {
		log.Fatalf("ultra group global banned set error %s", err)
	}

	// 查询超级群全体禁言
	globalBannedGetResp, err := rc.UltraGroupGlobalBannedGet(ctx, &rongcloud.UltraGroupGlobalBannedGetRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
	})
	if err != nil {
		log.Fatalf("ultra group global banned get error %s", err)
	}
	globalBannedGetRespData, _ := json.Marshal(globalBannedGetResp)
	log.Printf("ultra group global banned get resp data: %s", globalBannedGetRespData)
}
