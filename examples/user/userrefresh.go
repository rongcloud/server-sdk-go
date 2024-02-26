package main

import (
	"context"
	"log"
	"os"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	// 修改用户信息
	_, err := rc.UserRefresh(ctx, &rongcloud.UserRefreshRequest{
		UserId:      rongcloud.StringPtr("u01"),
		Name:        rongcloud.StringPtr("u01"),
		PortraitUri: nil,
	})
	if err != nil {
		log.Fatalf("user refresh error %s", err)
	}
}
