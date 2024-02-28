package main

import (
	"context"
	"log"
	"os"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	txtMsg := &rongcloud.TXTMsg{
		Content: "modify hello world",
	}
	content, err := txtMsg.ToString()
	if err != nil {
		log.Fatalf("txtMsg ToString() error %s", err)
	}
	// 修改超级群消息
	_, err = rc.UltraGroupMsgModify(ctx, &rongcloud.UltraGroupMsgModifyRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		FromUserId: rongcloud.StringPtr("u01"),
		MsgUID:     rongcloud.StringPtr("CDEV-2TTV-IKQ4-3F9N"),
		Content:    rongcloud.StringPtr(content),
	})
	if err != nil {
		log.Fatalf("ultragroup msg modify error %s", err)
	}
}
