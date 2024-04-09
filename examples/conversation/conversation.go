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

	// 会话置顶
	_, err := rc.ConversationTopSet(ctx, &rongcloud.ConversationTopSetRequest{
		UserId:           rongcloud.StringPtr("u01"),
		ConversationType: rongcloud.IntPtr(1),
		TargetId:         rongcloud.StringPtr("u02"),
		SetTop:           rongcloud.BoolPtr(true),
	})
	if err != nil {
		log.Fatalf("conversation top set error %s", err)
	}

	// 设置指定会话免打扰
	_, err = rc.ConversationNotificationSet(ctx, &rongcloud.ConversationNotificationSetRequest{
		ConversationType: rongcloud.IntPtr(1),
		RequestId:        rongcloud.StringPtr("u01"),
		TargetId:         rongcloud.StringPtr("u02"),
		IsMuted:          nil,
		BusChannel:       nil,
		UnpushLevel:      rongcloud.IntPtr(-1),
	})
	if err != nil {
		log.Fatalf("conversation notification set error %s", err)
	}

	// 查询指定会话免打扰
	notificationGetResp, err := rc.ConversationNotificationGet(ctx, &rongcloud.ConversationNotificationGetRequest{
		ConversationType: rongcloud.IntPtr(1),
		RequestId:        rongcloud.StringPtr("u01"),
		TargetId:         rongcloud.StringPtr("u02"),
		BusChannel:       nil,
	})
	if err != nil {
		log.Fatalf("conversation notification get error %s", err)
	}
	notificationGetRespData, _ := json.Marshal(notificationGetResp)
	log.Printf("conversation notification get resp data: %s", notificationGetRespData)

	// 设置指定会话类型免打扰
	_, err = rc.ConversationTypeNotificationSet(ctx, &rongcloud.ConversationTypeNotificationSetRequest{
		ConversationType: rongcloud.IntPtr(1),
		RequestId:        rongcloud.StringPtr("u01"),
		UnpushLevel:      rongcloud.IntPtr(-1),
	})
	if err != nil {
		log.Fatalf("conversation notification type set error %s", err)
	}

	// 查询指定会话类型免打扰
	typeNotificationGetResp, err := rc.ConversationTypeNotificationGet(ctx, &rongcloud.ConversationTypeNotificationGetRequest{
		ConversationType: rongcloud.IntPtr(1),
		RequestId:        rongcloud.StringPtr("u01"),
	})
	if err != nil {
		log.Fatalf("conversation type notification get error %s", err)
	}
	typeNotificationGetRespData, _ := json.Marshal(typeNotificationGetResp)
	log.Printf("conversation type notification get resp data: %s", typeNotificationGetRespData)
}
