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
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	// 搜索超级群消息
	yesterday := time.Now().Add(-time.Hour * 24)
	now := time.Now()
	hisMsgResp, err := rc.UltraGroupHismsgQuery(ctx, &rongcloud.UltraGroupHismsgQueryRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		StartTime:  rongcloud.Int64Ptr(yesterday.UnixMilli()),
		EndTime:    rongcloud.Int64Ptr(now.UnixMilli()),
		//FromUserId: rongcloud.StringPtr("u01"),
		PageSize: rongcloud.IntPtr(20),
	})
	if err != nil {
		log.Fatalf("ultra group his msg query error %s", err)
	}
	hisMsgData, _ := json.Marshal(hisMsgResp)
	log.Printf("ultra group his msg query data %s", hisMsgData)

	// 搜索超级群消息上下文
	hisIdMsg, err := rc.UltraGroupHismsgMsgIdQuery(ctx, &rongcloud.UltraGroupHisMsgMsgIdQueryRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		MsgUID:     rongcloud.StringPtr("CDEV-2TTV-IKQ4-3F9N"),
		PrevNum:    nil,
	})
	if err != nil {
		log.Fatalf("ultra group his msg msgId query error %s", err)
	}
	hisMsgIdData, _ := json.Marshal(hisIdMsg)
	log.Printf("ultra group his msg %s", hisMsgIdData)
}
