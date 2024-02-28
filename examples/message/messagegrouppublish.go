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

	// 发送群消息
	_, err := rc.MessageGroupPublish(ctx, &rongcloud.MessageGroupPublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToGroupId:  []string{"g01", "g02"},
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
	})
	if err != nil {
		log.Fatalf("message group publish error %s", err)
	}

	// 发送群定向消息
	_, err = rc.MessageGroupPublish(ctx, &rongcloud.MessageGroupPublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToGroupId:  []string{"g01"}, // 发送群聊定向消息时，仅支持传入一个群组 ID
		ToUserId:   []string{"u02", "u02"},
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
	})
	if err != nil {
		log.Fatalf("message group publish error %s", err)
	}
}
