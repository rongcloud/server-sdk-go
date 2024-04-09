package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	pushUserNotification := &rongcloud.PushUserNotification{
		Title:       rongcloud.StringPtr("标题"),
		PushContent: rongcloud.StringPtr("this is a push"),
		IOS: &rongcloud.PushUserNotificationIOS{
			ThreadId:       rongcloud.StringPtr("223"),
			ApnsCollapseId: rongcloud.StringPtr("111"),
			Extras: map[string]string{
				"id":   "1",
				"name": "2",
			},
		},
		Android: &rongcloud.PushAndroid{
			Honor: &rongcloud.PushAndroidHonor{
				Importance: rongcloud.StringPtr("NORMAL"),
				Image:      rongcloud.StringPtr("https://example.com/image.png"),
			},
			HW: &rongcloud.PushAndroidHW{
				ChannelId:  rongcloud.StringPtr("NotificationKanong"),
				Importance: rongcloud.StringPtr("NORMAL"),
				Image:      rongcloud.StringPtr("https://example.com/image.png"),
			},
			Oppo: &rongcloud.PushAndroidOppo{
				ChannelId: rongcloud.StringPtr("rc_notification_id"),
			},
			Vivo: &rongcloud.PushAndroidVivo{
				Classification: rongcloud.StringPtr("0"),
			},
			Fcm: &rongcloud.PushAndroidFcm{
				ChannelId: rongcloud.StringPtr("rc_notification_id"),
				ImageUrl:  rongcloud.StringPtr("https://example.com/image.png"),
			},
			Extras: map[string]string{
				"id":   "1",
				"name": "2",
			},
		},
	}

	// 发送指定用户不落地通知
	pushUserResp, err := rc.PushUser(ctx, &rongcloud.PushUserRequest{
		UserIds:      []string{"u01", "u02"},
		Notification: pushUserNotification,
	})
	if err != nil {
		log.Fatalf("push user response error %s", err)
	}
	pushUserRespData, _ := json.Marshal(pushUserResp)
	log.Printf("push user resp data: %s", pushUserRespData)
}
