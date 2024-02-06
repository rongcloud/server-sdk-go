package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	ctx := context.Background()
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))

	// 发送超级群消息
	txtMsg := &rongcloud.TXTMsg{
		Content: "hello world",
	}
	content, err := txtMsg.ToString()
	if err != nil {
		log.Fatalf("txt msg marshal content error %s", err)
	}
	extraContent := map[string]string{
		"key1": "key2",
	}
	extraContentStr, err := json.Marshal(extraContent)
	if err != nil {
		log.Fatalf("json.Marshal extraContent error %s", err)
	}
	_, err = rc.MessageUltraGroupPublish(ctx, &rongcloud.MessageUltraGroupPublishRequest{
		FromUserId:       rongcloud.StringPtr("u01"),
		ToGroupIds:       []string{"ug01"},
		ObjectName:       rongcloud.StringPtr(txtMsg.ObjectName()),
		Content:          rongcloud.StringPtr(content),
		PushContent:      nil,
		PushData:         nil,
		IsPersisted:      nil,
		IsCounted:        nil,
		IsMentioned:      nil,
		ContentAvailable: nil,
		PushExt:          nil,
		BusChannel:       rongcloud.StringPtr("channel01"),
		Expansion:        rongcloud.BoolPtr(true),
		ExtraContent:     rongcloud.StringPtr(string(extraContentStr)),
	})
	if err != nil {
		log.Fatalf("message ultra group publish error %s", err)
	}
}
