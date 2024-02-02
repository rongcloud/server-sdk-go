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

	_, err := rc.MessageChatroomPublish(ctx, &rongcloud.MessageChatroomPublishRequest{
		FromUserId:   rongcloud.StringPtr("u01"),
		ToChatroomId: []string{"chrm01", "chrm02"},
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
	})
	if err != nil {
		log.Fatalf("golang ")
	}
}
