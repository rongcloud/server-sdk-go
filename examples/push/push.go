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
	platform := []string{"ios", "android"}
	pushNotification := &rongcloud.PushNotification{
		Title:                rongcloud.StringPtr("标题"),
		ForceShowPushContent: nil,
		Alert:                rongcloud.StringPtr("this is a push"),
		IOS: &rongcloud.PushNotificationIOS{
			Title:             nil,
			ContentAvailable:  nil,
			Alert:             rongcloud.StringPtr("override alert"),
			Badge:             nil,
			ThreadId:          rongcloud.StringPtr("223"),
			ApnsCollapseId:    nil,
			Category:          nil,
			RichMediaUri:      nil,
			InterruptionLevel: nil,
			Extras:            map[string]string{"id": "1", "name": "2"},
		},
		Android: &rongcloud.PushNotificationAndroid{
			Alert: rongcloud.StringPtr("override alert"),
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
	txtMsg := &rongcloud.TXTMsg{
		Content: "hello world",
	}

	// 发送全量用户通知
	pushResp, err := rc.Push(ctx, &rongcloud.PushRequest{
		Platform: platform,
		Audience: &rongcloud.PushAudience{
			IsToAll: rongcloud.BoolPtr(true),
		},
		Message:      txtMsg,
		Notification: pushNotification,
	})
	if err != nil {
		log.Fatalf("push response error %s", err)
	}
	pushRespData, _ := json.Marshal(pushResp)
	log.Printf("push response data: %s", pushRespData)

	// 发送应用包名通知
	_, err = rc.Push(ctx, &rongcloud.PushRequest{
		Platform:   platform,
		FromUserid: rongcloud.StringPtr("u01"),
		Audience: &rongcloud.PushAudience{
			PackageName: rongcloud.StringPtr("xxx.rong.xxx"), // 根据应用包名发送通知只需要设置这个参数
			IsToAll:     rongcloud.BoolPtr(false),
		},
		Message:      txtMsg,
		Notification: pushNotification,
	})
	if err != nil {
		log.Fatalf("push response by packageName error %s", err)
	}

	// 发送标签用户通知
	_, err = rc.Push(ctx, &rongcloud.PushRequest{
		Platform:   platform,
		FromUserid: rongcloud.StringPtr("u01"),
		Audience: &rongcloud.PushAudience{
			Tag:     []string{"女", "年轻"},
			TagOr:   []string{"北京", "上海"},
			IsToAll: rongcloud.BoolPtr(false),
		},
		Message:      txtMsg,
		Notification: pushNotification,
	})
	if err != nil {
		log.Fatalf("push response by tag error %s", err)
	}
}
