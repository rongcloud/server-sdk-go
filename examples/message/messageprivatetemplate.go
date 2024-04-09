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

	// 发送单聊模板消息
	pushContent := "Dear {username}, please check your score."
	_, err := rc.MessagePrivatePublishTemplate(ctx, &rongcloud.MessagePrivatePublishTemplateRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		RCMsg: &rongcloud.TXTMsg{
			Content: "Dear {username}, your score is {score}",
		},
		MessageTemplate: []*rongcloud.MessageTemplate{
			{
				ToUserId: "u02",
				Value: map[string]string{
					"username": "u02",
					"score":    "1",
				},
				PushContent: pushContent,
				PushData:    "",
			}, {
				ToUserId: "u03",
				Value: map[string]string{
					"username": "u03",
					"score":    "2",
				},
				PushContent: pushContent,
				PushData:    "",
			},
		},
	})
	if err != nil {
		log.Fatalf("message private publish template error %s", err)
	}

	// 发送系统通知模板消息
	_, err = rc.MessageSystemPublishTemplate(ctx, &rongcloud.MessageSystemPublishTemplateRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		RCMsg: &rongcloud.TXTMsg{
			Content: "Dear {username}, your score is {score}",
		},
		MessageTemplate: []*rongcloud.MessageTemplate{
			{
				ToUserId: "u02",
				Value: map[string]string{
					"username": "u02",
					"score":    "1",
				},
				PushContent: pushContent,
				PushData:    "",
			}, {
				ToUserId: "u03",
				Value: map[string]string{
					"username": "u03",
					"score":    "2",
				},
				PushContent: pushContent,
				PushData:    "",
			},
		},
	})
	if err != nil {
		log.Fatalf("message system publish template error %s", err)
	}
}
