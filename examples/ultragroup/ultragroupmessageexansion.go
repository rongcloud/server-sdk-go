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

	testMsgUID := "CDF1-21C2-A73B-DE0J"
	// 设置超级群消息扩展
	extraKeyVal := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	extraKeyValStr, err := json.Marshal(extraKeyVal)
	if err != nil {
		log.Fatalf("json marshal extraKeyVal error %s", err)
	}
	_, err = rc.UltraGroupMessageExpansionSet(ctx, &rongcloud.UltraGroupMessageExpansionSetRequest{
		MsgUID:      rongcloud.StringPtr(testMsgUID),
		UserId:      rongcloud.StringPtr("u01"),
		BusChannel:  rongcloud.StringPtr("channel01"),
		GroupId:     rongcloud.StringPtr("ug01"),
		ExtraKeyVal: rongcloud.StringPtr(string(extraKeyValStr)),
	})
	if err != nil {
		log.Fatalf("ultra group message expansion set error %s", err)
	}

	// 获取超级群消息扩展
	messageExpansionQueryResp, err := rc.UltraGroupMessageExpansionQuery(ctx, &rongcloud.UltraGroupMessageExpansionQueryRequest{
		MsgUID:     rongcloud.StringPtr(testMsgUID),
		BusChannel: rongcloud.StringPtr("channel01"),
		GroupId:    rongcloud.StringPtr("ug01"),
		PageNo:     rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("ultragroup message expansion query error %s", err)
	}
	messageExpansionQueryRespData, _ := json.Marshal(messageExpansionQueryResp)
	log.Printf("ultra group message expansion query data %s", messageExpansionQueryRespData)

	// 删除超级群消息扩展
	extraKeys, err := json.Marshal([]string{"k1", "k2", "key1"})
	_, err = rc.UltraGroupMessageExpansionDelete(ctx, &rongcloud.UltraGroupMessageExpansionDeleteRequest{
		MsgUID:     rongcloud.StringPtr(testMsgUID),
		UserId:     rongcloud.StringPtr("u01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		GroupId:    rongcloud.StringPtr("ug01"),
		ExtraKey:   rongcloud.StringPtr(string(extraKeys)),
	})
	if err != nil {
		log.Fatalf("ultra group message expansion delete error %s", err)
	}
}
