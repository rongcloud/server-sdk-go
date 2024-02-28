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

	chatroomBan(ctx, rc)
}

func chatroomBan(ctx context.Context, rc *rongcloud.RongCloud) {
	grp := "grp1"
	// 创建测试chatroom
	_, err := rc.ChatroomCreateNew(ctx, &rongcloud.ChatroomCreateNewRequest{ChatroomId: rongcloud.StringPtr(grp)})
	if err != nil {
		log.Fatalf("chatroom create error %s", err)
	}

	banReq := rongcloud.ChatroomBanRequest{
		ChatroomId: rongcloud.StringPtr(grp),
		Extra:      rongcloud.StringPtr(""),
		NeedNotify: rongcloud.BoolPtr(true),
	}

	// 设置聊天室全体禁言
	banAddRequest := &rongcloud.ChatroomBanAddRequest{
		ChatroomBanRequest: banReq,
	}
	_, err = rc.ChatroomBanAdd(ctx, banAddRequest)
	if err != nil {
		log.Fatalf("chatroom ban add error %s", err)
	}

	// 查询聊天室全体禁言列表
	banQueryResp, err := rc.ChatroomBanQuery(ctx, &rongcloud.ChatroomBanQueryRequest{
		Size: rongcloud.IntPtr(50),
		Page: rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("chatroom ban query error %s", err)
	}
	banQueryRespData, _ := json.Marshal(banQueryResp)
	log.Printf("chatroom ban query resp data: %s", banQueryRespData)

	banCheckResp, err := rc.ChatroomBanCheck(ctx, &rongcloud.ChatroomBanCheckRequest{ChatroomId: rongcloud.StringPtr(grp)})
	if err != nil {
		log.Fatalf("chatroom ban check error %s", err)
	}
	banCheckRespData, _ := json.Marshal(banCheckResp)
	log.Printf("chatroom ban check resp data: %s", banCheckRespData)

	// 取消聊天室全体禁言
	banRollbackRequest := &rongcloud.ChatroomBanRollbackRequest{
		ChatroomBanRequest: banReq,
	}
	_, err = rc.ChatroomBanRollback(ctx, banRollbackRequest)
	if err != nil {
		log.Fatalf("chatroom ban rollback error %s", err)
	}

	// 加入聊天室全体禁言白名单
	userBanWhiteListReq := rongcloud.ChatroomUserBanWhitelistRequest{
		ChatroomId: rongcloud.StringPtr(grp),
		UserIds:    []string{"u01", "u02"},
		Extra:      rongcloud.StringPtr(""),
		NeedNotify: rongcloud.BoolPtr(true),
	}
	userBanWhitelistAddReq := &rongcloud.ChatroomUserBanWhitelistAddRequest{
		ChatroomUserBanWhitelistRequest: userBanWhiteListReq,
	}
	_, err = rc.ChatroomUserBanWhitelistAdd(ctx, userBanWhitelistAddReq)
	if err != nil {
		log.Fatalf("chatroom user ban whitelist add error %s", err)
	}

	// 查询聊天室全体禁言白名单
	userBanWhitelistQueryReq := &rongcloud.ChatroomUserBanWhitelistQueryRequest{
		ChatroomId: rongcloud.StringPtr(grp),
	}
	userBanWhitelistQueryResp, err := rc.ChatroomUserBanWhitelistQuery(ctx, userBanWhitelistQueryReq)
	if err != nil {
		log.Fatalf("chatroom user ban whitelist query error %s", err)
	}
	userBanWhitelistQueryRespData, _ := json.Marshal(userBanWhitelistQueryResp)
	log.Printf("chatroom user ban whitelist query resp data: %s", userBanWhitelistQueryRespData)

	// 移出聊天室全体禁言白名单
	userBanWhitelistRollbackReq := &rongcloud.ChatroomUserBanWhitelistRollbackRequest{
		ChatroomUserBanWhitelistRequest: userBanWhiteListReq,
	}
	_, err = rc.ChatroomUserBanWhitelistRollback(ctx, userBanWhitelistRollbackReq)
	if err != nil {
		log.Fatalf("chatroom user ban whitelist rollback error %s", err)
	}

	// 销毁测试chatroom
	_, err = rc.ChatroomDestroy(ctx, &rongcloud.ChatroomDestroyRequest{ChatroomIds: []string{grp}})
	if err != nil {
		log.Fatalf("chatroom create error %s", err)
	}
}
