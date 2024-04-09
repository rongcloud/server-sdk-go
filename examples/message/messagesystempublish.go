package main

import (
	"context"
	"log"
	"os"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	ctx := context.Background()
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))

	// 发送系统通知普通消息
	_, err := rc.MessageSystemPublish(ctx, &rongcloud.MessageSystemPublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToUserId:   []string{"u02", "u03"},
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
	})
	if err != nil {
		log.Fatalf("message system publish error %s", err)
	}
}
