package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	ctx := context.Background()
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))

	now := time.Now().Format("2006010215")

	// 获取历史消息日志
	messageHistory, err := rc.MessageHistory(ctx, &rongcloud.MessageHistoryRequest{
		Date: rongcloud.StringPtr(now),
	})
	if err != nil {
		log.Fatalf("message history error %s", err)
	}
	messageHistoryData, _ := json.Marshal(messageHistory)
	log.Printf("message history %s", messageHistoryData)

	// 删除历史消息日志
	_, err = rc.MessageHistoryDelete(ctx, &rongcloud.MessageHistoryDeleteRequest{
		Date: rongcloud.StringPtr(now),
	})
	if err != nil {
		log.Fatalf("message history delete error %s", err)
	}
}
