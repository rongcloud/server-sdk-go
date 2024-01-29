package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

const (
	testChatroomId = "grp1"
)

func main() {
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	// create test chatroom
	_, err := rc.ChatroomCreateNew(ctx, &rongcloud.ChatroomCreateNewRequest{
		ChatroomId: rongcloud.StringPtr(testChatroomId),
	})
	if err != nil {
		log.Fatalf("chatroom create new error %s", err)
	}

	chatroomUser(ctx, rc)

	chatroomUserBlock(ctx, rc)

	chatroomMessagePriority(ctx, rc)

	chatroomKeepAlive(ctx, rc)

	// destroy test chatroom
	_, err = rc.ChatroomDestroy(ctx, &rongcloud.ChatroomDestroyRequest{
		ChatroomIds: []string{testChatroomId},
	})
	if err != nil {
		log.Fatalf("chatroom destroy error %s", err)
	}
}

func chatroomUser(ctx context.Context, rc *rongcloud.RongCloud) {
	// 查询用户是否在聊天室中
	userExistResp, err := rc.ChatroomUserExist(ctx,
		&rongcloud.ChatroomUserExistRequest{
			ChatroomId: rongcloud.StringPtr(testChatroomId),
			UserId:     rongcloud.StringPtr("u01"),
		},
	)
	if err != nil {
		log.Fatalf("chatroom user exit err %s", err)
	}
	b, _ := json.Marshal(userExistResp)
	log.Printf("chatroom user exist resp data: %s", b)

	// 批量查询用户是否在聊天室中
	usersExistResp, err := rc.ChatroomUsersExist(ctx, &rongcloud.ChatroomUsersExistRequest{
		ChatroomId: rongcloud.StringPtr(testChatroomId),
		UserIds:    []string{"u01", "u02"},
	})
	if err != nil {
		log.Fatalf("chatroom users exist err %s", err)
	}
	usersExistRespData, _ := json.Marshal(usersExistResp)
	log.Printf("chatroom users exist resp: %s", usersExistRespData)

	// 获取聊天室成员
	chatroomUserQueryResp, err := rc.ChatroomUserQuery(ctx, &rongcloud.ChatroomUserQueryRequest{
		ChatroomId: rongcloud.StringPtr(testChatroomId),
		Count:      rongcloud.IntPtr(200),
		Order:      rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("chatroom user query err %s", err)
	}
	data, _ := json.Marshal(chatroomUserQueryResp)
	log.Printf("chatroom user query resp data: %s", data)
}

func chatroomUserBlock(ctx context.Context, rc *rongcloud.RongCloud) {
	// 封禁聊天室用户
	resp, err := rc.ChatroomUserBlockAdd(ctx, &rongcloud.ChatroomUserBlockAddRequest{
		UserIds:    []string{"u01", "u02"},
		ChatroomId: rongcloud.StringPtr(testChatroomId),
		Minute:     rongcloud.IntPtr(10),
		Extra:      rongcloud.StringPtr(""),
		NeedNotify: rongcloud.BoolPtr(true),
	})
	if err != nil {
		log.Fatalf("chatroom user block add err %s", err)
	}
	d, _ := json.Marshal(resp)
	log.Printf("chatroom user block data: %s", d)

	// 查询聊天室封禁用户
	userBlockListResp, err := rc.ChatroomUserBlockList(ctx, &rongcloud.ChatroomUserBlockListRequest{
		ChatroomId: rongcloud.StringPtr(testChatroomId)})
	if err != nil {
		log.Fatalf("chatroom user block list err: %s", err)
	}
	userBlockListData, _ := json.Marshal(userBlockListResp)
	log.Printf("chatroom user block list data: %s", userBlockListData)

	// 解除封禁聊天室用户
	rollbackResp, err := rc.ChatroomUserBlockRollback(ctx, &rongcloud.ChatroomUserBlockRollbackRequest{
		UserIds:    []string{"u01", "u02"},
		ChatroomId: rongcloud.StringPtr(testChatroomId),
		Extra:      rongcloud.StringPtr(""),
		NeedNotify: rongcloud.BoolPtr(true),
	})
	if err != nil {
		log.Fatalf("chatroom user block rollback err %s", err)
	}
	rollbackData, _ := json.Marshal(rollbackResp)
	log.Printf("chatroom user rollback data: %s", rollbackData)

	// 全局禁言用户
	userBanAddResp, err := rc.ChatroomUserBanAdd(ctx, &rongcloud.ChatroomUserBanAddRequest{
		UserIds:    []string{"u01", "u02"},
		Minute:     rongcloud.IntPtr(10),
		Extra:      rongcloud.StringPtr(""),
		NeedNotify: rongcloud.BoolPtr(true),
	})
	if err != nil {
		log.Fatalf("chatroom user ban add err %s", err)
	}
	userBanAddData, _ := json.Marshal(userBanAddResp)
	log.Printf("chatroom user ban add data: %s", userBanAddData)

	// 查询全局禁言用户列表
	userBanQueryResp, err := rc.ChatroomUserBanQuery(ctx, nil)
	if err != nil {
		log.Fatalf("chatroom user ban query err %s", err)
	}
	userBanQueryData, _ := json.Marshal(userBanQueryResp)
	log.Printf("chatroom user ban query data: %s", userBanQueryData)

	// 取消全局禁言用户
	userBanRemoveResp, err := rc.ChatroomUserBanRemove(ctx, &rongcloud.ChatroomUserBanRemoveRequest{
		UserIds:    []string{"u01", "u02"},
		Extra:      rongcloud.StringPtr(""),
		NeedNotify: rongcloud.BoolPtr(true),
	})
	if err != nil {
		log.Fatalf("chatroom user ban remove data: %s", err)
	}
	userBanRemoveData, _ := json.Marshal(userBanRemoveResp)
	log.Printf("chatroom user ban remove data: %s", userBanRemoveData)

	// 禁言指定聊天室用户
	userGagAddResp, err := rc.ChatroomUserGagAdd(ctx, &rongcloud.ChatroomUserGagAddRequest{
		UserIds:    []string{"u01", "u02"},
		ChatroomId: rongcloud.StringPtr(testChatroomId),
		Minute:     rongcloud.IntPtr(10),
		Extra:      rongcloud.StringPtr(""),
		NeedNotify: rongcloud.BoolPtr(true),
	})
	if err != nil {
		log.Fatalf("chatroom user gag add err %s", err)
	}
	userGagAddData, _ := json.Marshal(userGagAddResp)
	log.Printf("chatroom user gag add data: %s", userGagAddData)

	// 查询聊天室禁言用户列表
	userGagListResp, err := rc.ChatroomUserGagList(ctx, &rongcloud.ChatroomUserGagListRequest{
		ChatroomId: rongcloud.StringPtr(testChatroomId),
	})
	if err != nil {
		log.Fatalf("chatroom user gag list err %s", err)
	}
	userGagListData, _ := json.Marshal(userGagListResp)
	log.Printf("chatroom user gag list data: %s", userGagListData)

	//  取消禁言指定聊天室用户
	userGagRollbackResp, err := rc.ChatroomUserGagRollback(ctx, &rongcloud.ChatroomUserGagRollbackRequest{
		UserIds:    []string{"u01", "u02"},
		ChatroomId: rongcloud.StringPtr(testChatroomId),
		Extra:      rongcloud.StringPtr(""),
		NeedNotify: rongcloud.BoolPtr(true),
	})
	if err != nil {
		log.Fatalf("chatroom user gag rollback err %s", err)
	}
	userGagRollbackData, _ := json.Marshal(userGagRollbackResp)
	log.Printf("chatroom user gag rollback data: %s", userGagRollbackData)
}

func chatroomMessagePriority(ctx context.Context, rc *rongcloud.RongCloud) {
	// add message priority
	vcMsg := "RC:VcMsg"
	imgMsg := "RC:ImgMsg"
	addResp, err := rc.ChatroomMessagePriorityAdd(ctx, &rongcloud.ChatroomMessagePriorityAddRequest{
		ObjectNames: []string{vcMsg, imgMsg},
	})
	if err != nil {
		log.Fatalf("chatroom message priority add err %s", err)
	}
	addData, _ := json.Marshal(addResp)
	log.Printf("chatroom message priority add resp: %s", addData)

	// query after add
	query2Resp, err := rc.ChatroomMessagePriorityQuery(ctx, nil)
	if err != nil {
		log.Fatalf("chatroom message priority query err: %s", err)
	}
	query2Data, _ := json.Marshal(query2Resp)
	log.Printf("chatroom query response data %s", query2Data)

	// remove message priority
	removeResp, err := rc.ChatroomMessagePriorityRemove(ctx, &rongcloud.ChatroomMessagePriorityRemoveRequest{
		ObjectNames: []string{vcMsg, imgMsg},
	})
	if err != nil {
		log.Fatalf("chatroom message priority remove err %s", err)
	}
	removeData, _ := json.Marshal(removeResp)
	log.Printf("chatroom message priority remove resp: %s", removeData)
}

func chatroomKeepAlive(ctx context.Context, rc *rongcloud.RongCloud) {
	addResp, err := rc.ChatroomKeepaliveAdd(ctx, &rongcloud.ChatroomKeepaliveAddRequest{ChatroomId: rongcloud.StringPtr("grp1")})
	if err != nil {
		log.Fatalf("chatroom keepalive add resp err %s", err)
	}
	addData, _ := json.Marshal(addResp)
	log.Printf("chatroom keepalive add resp: %s", addData)
	queryResp, err := rc.ChatroomKeepaliveQuery(ctx, nil)
	if err != nil {
		log.Fatalf("chatroom keepalive query resp err %s", err)
	}
	queryData, _ := json.Marshal(queryResp)
	log.Printf("chatroom keepalive query resp: %s", queryData)
	removeResp, err := rc.ChatroomKeepaliveRemove(ctx, &rongcloud.ChatroomKeepaliveRemoveRequest{ChatroomId: rongcloud.StringPtr("grp1")})
	if err != nil {
		log.Fatalf("chatroom keepalive remove err %s", err)
	}
	removeData, _ := json.Marshal(removeResp)
	log.Printf("chatroom keepalive remove resp %s", removeData)
}
