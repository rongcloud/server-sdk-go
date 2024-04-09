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

	// 设置用户免打扰时段
	_, err := rc.UserBlockPushPeriodSet(ctx, &rongcloud.UserBlockPushPeriodSetRequest{
		UserId:    rongcloud.StringPtr("u01"),
		StartTime: rongcloud.StringPtr("22:00:00"),
		Period:    rongcloud.IntPtr(60 * 9), // 22:00:00 - 07:00:00 时段免打扰
		Level:     rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("user block push period set error %s", err)
	}

	// 查询用户免打扰时段
	userBlockPushPeriodGetResp, err := rc.UserBlockPushPeriodGet(ctx, &rongcloud.UserBlockPushPeriodGetRequest{
		UserId: rongcloud.StringPtr("u01"),
	})
	if err != nil {
		log.Fatalf("user block push period get error %s", err)
	}
	userBlockPushPeriodGetRespData, _ := json.Marshal(userBlockPushPeriodGetResp)
	log.Printf("user block push period get response data: %s", userBlockPushPeriodGetRespData)

	// 删除用户免打扰时段
	_, err = rc.UserBlockPushPeriodDelete(ctx, &rongcloud.UserBlockPushPeriodDeleteRequest{
		UserId: rongcloud.StringPtr("u01"),
	})
	if err != nil {
		log.Fatalf("user block push period delete error %s", err)
	}
}
