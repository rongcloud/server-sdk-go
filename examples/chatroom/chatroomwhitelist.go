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

	chatroomWhitelist(ctx, rc)

	chatroomUserWhitelist(ctx, rc)
}

func chatroomWhitelist(ctx context.Context, rc *rongcloud.RongCloud) {
	//  加入聊天室消息白名单
	_, err := rc.ChatroomWhitelistAdd(ctx, &rongcloud.ChatroomWhitelistAddRequest{
		ObjectNames: []string{"RC:VcMsg", "RC:ImgTextMsg"}})
	if err != nil {
		log.Fatalf("chatroom white list add error %s", err)
	}

	// 查询聊天室消息白名单
	whitelistQueryResp, err := rc.ChatroomWhitelistQuery(ctx, nil)
	if err != nil {
		log.Fatalf("chatroom white list query error %s", err)
	}
	whitelistQueryData, _ := json.Marshal(whitelistQueryResp)
	log.Printf("chatroom white list query data: %s", whitelistQueryData)

	// 移出聊天室消息白名单
	_, err = rc.ChatroomWhitelistRemove(ctx, &rongcloud.ChatroomWhitelistRemoveRequest{
		ObjectNames: []string{"RC:VcMsg", "RC:ImgTextMsg"},
	})
	if err != nil {
		log.Fatalf("chatroom white list remove error: %s", err)
	}
}

func chatroomUserWhitelist(ctx context.Context, rc *rongcloud.RongCloud) {
	// 创建测试chatroom
	grp := "grp1"
	_, err := rc.ChatroomCreateNew(ctx, &rongcloud.ChatroomCreateNewRequest{
		ChatroomId: rongcloud.StringPtr(grp),
	})
	if err != nil {
		log.Fatalf("chatroom create error %s", err)
	}

	// 加入聊天室用户白名单
	_, err = rc.ChatroomUserWhitelistAdd(ctx, &rongcloud.ChatroomUserWhitelistAddRequest{
		ChatroomId: rongcloud.StringPtr(grp),
		UserIds:    []string{"u01", "u02"},
	})
	if err != nil {
		log.Fatalf("chatroom user whitelist add error %s", err)
	}

	// 查询聊天室用户白名单
	queryUserWhitelistResp, err := rc.ChatroomUserWhitelistQuery(ctx, &rongcloud.ChatroomUserWhitelistQueryRequest{ChatroomId: rongcloud.StringPtr(grp)})
	if err != nil {
		log.Fatalf("chatroom user whitelist query error %s", err)
	}
	queryUserWhitelistRespData, _ := json.Marshal(queryUserWhitelistResp)
	log.Printf("query user whitelist resp data: %s", queryUserWhitelistRespData)

	// 移出聊天室用户白名单
	_, err = rc.ChatroomUserWhitelistRemove(ctx, &rongcloud.ChatroomUserWhitelistRemoveRequest{
		ChatroomId: rongcloud.StringPtr(grp),
		UserIds:    []string{"u01", "u02"},
	})
	if err != nil {
		log.Printf("chatroom user whitlist ")
	}

	// 销毁测试chatroom
	_, err = rc.ChatroomDestroy(ctx, &rongcloud.ChatroomDestroyRequest{ChatroomIds: []string{grp}})
	if err != nil {
		log.Fatalf("chatroom destroy error %s", err)
	}
}
