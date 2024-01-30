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

	// 设置聊天室全体禁言
	banAddRequest := &rongcloud.ChatroomBanAddRequest{}
	banAddRequest.ChatroomId = rongcloud.StringPtr(grp)
	banAddRequest.Extra = rongcloud.StringPtr("")
	banAddRequest.NeedNotify = rongcloud.BoolPtr(true)
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
	banRollbackRequest := &rongcloud.ChatroomBanRollbackRequest{}
	banRollbackRequest.ChatroomId = rongcloud.StringPtr(grp)
	banRollbackRequest.Extra = rongcloud.StringPtr("")
	banRollbackRequest.NeedNotify = rongcloud.BoolPtr(true)
	_, err = rc.ChatroomBanRollback(ctx, banRollbackRequest)
	if err != nil {
		log.Fatalf("chatroom ban rollback error %s", err)
	}

	// 销毁测试chatroom
	_, err = rc.ChatroomDestroy(ctx, &rongcloud.ChatroomDestroyRequest{ChatroomIds: []string{grp}})
	if err != nil {
		log.Fatalf("chatroom create error %s", err)
	}
}
