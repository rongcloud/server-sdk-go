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

	// 查询用户在线状态
	userCheckOnlineResp, err := rc.UserCheckOnline(ctx, &rongcloud.UserCheckOnlineRequest{
		UserId: rongcloud.StringPtr("u01"),
	})
	if err != nil {
		log.Fatalf("user check online error %s", err)
	}
	userCheckOnlineRespData, _ := json.Marshal(userCheckOnlineResp)
	log.Printf("user check online response data: %s", userCheckOnlineRespData)
}
