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

	//  封禁用户
	_, err := rc.UserBlock(ctx, &rongcloud.UserBlockRequest{
		UserId: []string{"u01", "u02"},
		Minute: rongcloud.IntPtr(43200),
	})
	if err != nil {
		log.Fatalf("user block error %s", err)
	}

	// 获取封禁用户列表
	userBlockQueryResp, err := rc.UserBlockQuery(ctx, &rongcloud.UserBlockQueryRequest{
		Page: rongcloud.IntPtr(1),
		Size: rongcloud.IntPtr(50),
	})
	if err != nil {
		log.Fatalf("user block error %s", err)
	}
	userBlockQueryRespData, _ := json.Marshal(userBlockQueryResp)
	log.Printf("user block query response data: %s", userBlockQueryRespData)

	// 解除封禁
	_, err = rc.UserUnBlock(ctx, &rongcloud.UserUnBlockRequest{
		UserId: []string{"u01", "u02"},
	})
	if err != nil {
		log.Fatalf("user unblock error %s", err)
	}
}
