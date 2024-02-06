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

	testMsgUID := "CDCV-7C6J-P604-3F9O"
	// 设置单群聊消息扩展
	extraKeyVal := map[string]string{
		"key1": "val1",
	}
	_, err := rc.MessageExpansionSet(ctx, &rongcloud.MessageExpansionSetRequest{
		MsgUID:           rongcloud.StringPtr(testMsgUID),
		UserId:           rongcloud.StringPtr("u01"),
		ConversationType: rongcloud.StringPtr(rongcloud.ConversationTypePrivate),
		TargetId:         rongcloud.StringPtr("u02"),
		ExtraKeyVal:      extraKeyVal,
		IsSyncSender:     rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("message expansion set error %s", err)
	}
	log.Printf("message expansion set success")

	// 获取单群聊消息扩展
	expansionResp, err := rc.MessageExpansionQuery(ctx, &rongcloud.MessageExpansionQueryRequest{
		MsgUID: rongcloud.StringPtr(testMsgUID),
		PageNo: rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("message expansion query error %s", err)
	}
	for k, value := range expansionResp.ExtraContent {
		log.Printf("expansion query key: %s, value: %+v", k, value)
	}

	// 删除单群聊消息扩展
	_, err = rc.MessageExpansionDelete(ctx, &rongcloud.MessageExpansionDeleteRequest{
		MsgUID:           rongcloud.StringPtr(testMsgUID),
		UserId:           rongcloud.StringPtr("u01"),
		ConversationType: rongcloud.StringPtr(rongcloud.ConversationTypePrivate),
		TargetId:         rongcloud.StringPtr("u02"),
		ExtraKey:         []string{"key1", "key2"},
		IsSyncSender:     rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("message expansion delete error %s", err)
	}
	log.Printf("message expansion delete success")
}
