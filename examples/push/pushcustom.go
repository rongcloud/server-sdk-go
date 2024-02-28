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

	pushNotification := &rongcloud.PushCustomNotification{
		Title: rongcloud.StringPtr("标题"),
		Alert: rongcloud.StringPtr("this is a push"),
		IOS: &rongcloud.PushCustomNotificationIOS{
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
	// 示例 1：按用户标签推送（同时使用 audience.tag 与 audience.tag_or）
	pushResp, err := rc.PushCustom(ctx, &rongcloud.PushCustomRequest{
		Platform: []string{"ios", "android"},
		Audience: &rongcloud.PushCustomAudience{
			Tag:         []string{"女", "年轻"},
			TagOr:       []string{"北京", "上海"},
			PackageName: nil,
			IsToAll:     nil,
			TagItems:    nil,
		},
		Notification: pushNotification,
	})
	if err != nil {
		log.Fatalf("push custom example 1 error %s", err)
	}
	pushRespData, _ := json.Marshal(pushResp)
	log.Printf("push response data: %s", pushRespData)

	// 示例 2：按用户标签推送（使用 audience.tagItems）
	pushResp2, err := rc.PushCustom(ctx, &rongcloud.PushCustomRequest{
		Platform: []string{"ios", "android"},
		Audience: &rongcloud.PushCustomAudience{
			TagItems: []*rongcloud.PushCustomAudienceTagItem{
				{
					Tags:          []string{"guangdong", "hunan"},
					IsNot:         rongcloud.BoolPtr(false),
					TagsOperator:  rongcloud.StringPtr("OR"),
					ItemsOperator: rongcloud.StringPtr("OR"),
				}, {
					Tags:          []string{"20200408"},
					IsNot:         rongcloud.BoolPtr(true),
					TagsOperator:  rongcloud.StringPtr("OR"),
					ItemsOperator: rongcloud.StringPtr("AND"),
				}, {
					Tags:          []string{"male", "female"},
					IsNot:         rongcloud.BoolPtr(false),
					TagsOperator:  rongcloud.StringPtr("OR"),
					ItemsOperator: rongcloud.StringPtr("OR"),
				},
			},
			IsToAll: rongcloud.BoolPtr(false),
		},
		Notification: pushNotification,
	})
	if err != nil {
		log.Fatalf("push custom example 2 error %s", err)
	}
	pushResp2Data, _ := json.Marshal(pushResp2)
	log.Printf("push response example 2 data: %s", pushResp2Data)
}
