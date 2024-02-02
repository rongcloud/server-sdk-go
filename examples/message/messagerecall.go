package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	ctx := context.Background()
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))

	// 撤回消息
	_, err := rc.MessageRecall(ctx, &rongcloud.MessageRecallRequest{
		FromUserId:       rongcloud.StringPtr("u01"),
		ConversationType: rongcloud.IntPtr(1),
		TargetId:         rongcloud.StringPtr("u02"),
		BusChannel:       nil,
		MessageUID:       rongcloud.StringPtr("CDCF-J1T2-UJ64-3F9N"),
		SentTime:         rongcloud.Int64Ptr(time.Now().Unix()),
		IsAdmin:          nil,
		IsDelete:         nil,
		DisablePush:      nil,
		Extra:            nil,
	})
	if err != nil {
		log.Fatalf("message recall error %s", err)
	}
}
