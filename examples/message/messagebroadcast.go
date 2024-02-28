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

	// 发送全量用户落地通知
	_, err := rc.MessageBroadcast(ctx, &rongcloud.MessageBroadcastRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
		PushContent:      nil,
		PushData:         nil,
		ContentAvailable: nil,
		PushExt:          nil,
	})
	if err != nil {
		log.Fatalf("message broadcast error %s", err)
	}

	// 发送全体聊天室广播消息
	_, err = rc.MessageChatroomBroadcast(ctx, &rongcloud.MessageChatroomBroadcastRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
		IsIncludeSender: nil,
	})
	if err != nil {
		log.Fatalf("message chatroom broadcast error %s", err)
	}

	// 发送在线用户广播
	_, err = rc.MessageOnlineBroadcast(ctx, &rongcloud.MessageOnlineBroadcastRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		RCMsg: &rongcloud.TXTMsg{
			Content: "hello world",
		},
	})
	if err != nil {
		log.Fatalf("message online broadcast error %s", err)
	}
}
