package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	ctx := context.Background()
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))

	messagePrivate(ctx, rc)
}

func messagePrivate(ctx context.Context, rc *rongcloud.RongCloud) {
	// 发送单聊文本消息
	txtMsg := &rongcloud.TXTMsg{
		Content: "hello world",
	}
	requestId := uuid.New().String()
	rongcloud.AddHttpRequestId(ctx, requestId)
	_, err := rc.MessagePrivatePublish(ctx, &rongcloud.MessagePrivatePublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToUserId:   rongcloud.StringPtr("u02"),
		RCMsg:      txtMsg,
	})
	if err != nil {
		log.Fatalf("msg private publish error: %s", err)
	}
	// 发送单聊语音消息
	hqvsMsg := &rongcloud.HQVCMsg{
		RemoteUrl: "http://example.com/1",
		Duration:  2,
	}
	_, err = rc.MessagePrivatePublish(ctx, &rongcloud.MessagePrivatePublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToUserId:   rongcloud.StringPtr("u02"),
		RCMsg:      hqvsMsg,
	})
	if err != nil {
		log.Fatalf("hqvs msg private publish error: %s", err)
	}
	// 发送单聊图片消息
	// 更多信息请参考 https://doc.rongcloud.cn/imserver/server/v1/message/objectname#%E5%9B%BE%E7%89%87%E6%B6%88%E6%81%AF
	imgMsg := &rongcloud.ImgMsg{
		Content:  "/9j/4AAQSkZJRgABAgAAZABkAAD",
		ImageURI: "http://p1.cdn.com/fds78ruhi.jpg",
		Extra:    "",
	}
	_, err = rc.MessagePrivatePublish(ctx, &rongcloud.MessagePrivatePublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToUserId:   rongcloud.StringPtr("u02"),
		RCMsg:      imgMsg,
	})
	if err != nil {
		log.Printf("img msg private publish error: %s", err)
	}

	// 发送单聊自定义消息
	customMsg := &CustomJsonMsg{
		Field1: "custom json msg",
		Field2: 1,
		Field3: true,
	}
	_, err = rc.MessagePrivatePublish(ctx, &rongcloud.MessagePrivatePublishRequest{
		FromUserId: rongcloud.StringPtr("u01"),
		ToUserId:   rongcloud.StringPtr("u02"),
		RCMsg:      customMsg,
	})
	if err != nil {
		log.Fatalf("custom msg private publish error: %s", err)
	}
}

// 自定义消息类型需开发者自行实现rongcloud.RCMsg接口
type CustomJsonMsg struct {
	Field1 string
	Field2 int
	Field3 bool
}

func (m *CustomJsonMsg) ObjectName() string {
	return "RC:CustomMsg"
}

func (m *CustomJsonMsg) ToString() (string, error) {
	b, err := json.Marshal(m)
	return string(b), err
}
