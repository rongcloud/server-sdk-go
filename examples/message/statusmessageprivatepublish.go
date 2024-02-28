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

	// 发送单聊状态消息
	_, err := rc.StatusMessagePrivatePublish(ctx, &rongcloud.StatusMessagePrivatePublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToUserIds:  []string{"u02", "u03"},
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
		VerifyBlacklist: nil,
		IsIncludeSender: nil,
	})
	if err != nil {
		log.Fatalf("status message private publish error %s", err)
	}

	// 发送群聊状态消息
	_, err = rc.StatusMessageGroupPublish(ctx, &rongcloud.StatusMessageGroupPublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToGroupId:  []string{"g01", "g02"},
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
	})
	if err != nil {
		log.Fatalf("status message group publish error %s", err)
	}
}
