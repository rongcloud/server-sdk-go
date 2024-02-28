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
	const personType = "PERSON"

	// 设置用户单聊禁言
	_, err := rc.UserChatFBSet(ctx, &rongcloud.UserChatFBSetRequest{
		UserId: []string{"u02", "u01"},
		State:  rongcloud.IntPtr(1),
		Type:   rongcloud.StringPtr(personType),
	})
	if err != nil {
		log.Fatalf("user chat fb set error %s", err)
	}

	// 查询单聊禁言用户列表
	userChatFBQueryListResp, err := rc.UserChatFBQueryList(ctx, &rongcloud.UserChatFBQueryListRequest{
		Num:    rongcloud.IntPtr(50),
		Offset: rongcloud.IntPtr(0),
		Type:   rongcloud.StringPtr(personType),
	})
	if err != nil {
		log.Fatalf("user chat fb query list error %s", err)
	}
	userChatFBQueryListRespData, _ := json.Marshal(userChatFBQueryListResp)
	log.Printf("user chat fb query list response data: %s", userChatFBQueryListRespData)
}
