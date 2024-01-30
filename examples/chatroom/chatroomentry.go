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

	chatroomEntry(ctx, rc)
}

func chatroomEntry(ctx context.Context, rc *rongcloud.RongCloud) {
	grp := "grp1"
	_, err := rc.ChatroomCreateNew(ctx, &rongcloud.ChatroomCreateNewRequest{
		ChatroomId: rongcloud.StringPtr(grp),
	})
	if err != nil {
		log.Fatalf("chatroom create error: %s", err)
	}

	// 设置聊天室属性（KV）
	_, err = rc.ChatroomEntrySet(ctx, &rongcloud.ChatroomEntrySetRequest{
		ChatroomId: rongcloud.StringPtr(grp),
		UserId:     rongcloud.StringPtr("u01"),
		Key:        rongcloud.StringPtr("key1"),
		Value:      rongcloud.StringPtr("val1"),
		AutoDelete: rongcloud.IntPtr(1),
		RCMsg: &rongcloud.ChrmKVNotiMsg{
			Type:  1,
			Key:   "key1",
			Value: "val1",
			Extra: "extra info",
		},
	})
	if err != nil {
		log.Fatalf("chatroom entry set err %s", err)
	}

	// 批量设置聊天室属性（KV）
	_, err = rc.ChatroomEntryBatchSet(ctx, &rongcloud.ChatroomEntryBatchSetRequest{
		ChatroomId:   rongcloud.StringPtr(grp),
		AutoDelete:   rongcloud.IntPtr(1),
		EntryOwnerId: rongcloud.StringPtr("u01"),
		EntryInfo:    map[string]string{"k1": "k2"},
	})
	if err != nil {
		log.Fatalf("chatroom entry batch set err %s", err)
	}

	entryQueryResp, err := rc.ChatroomEntryQuery(ctx, &rongcloud.ChatroomEntryQueryRequest{
		ChatroomId: rongcloud.StringPtr(grp),
		Keys:       []string{"key1"},
	})
	if err != nil {
		log.Fatalf("chatroom entry query error %s", err)
	}
	entryQueryRespData, err := json.Marshal(entryQueryResp)
	log.Printf("chatroom entry query resp: %s", entryQueryRespData)

	// 删除聊天室属性（KV)
	_, err = rc.ChatroomEntryRemove(ctx, &rongcloud.ChatroomEntryRemoveRequest{
		ChatroomId: rongcloud.StringPtr(grp),
		UserId:     rongcloud.StringPtr("u01"),
		Key:        rongcloud.StringPtr("key1"),
	})
	if err != nil {
		log.Fatalf("chatroom entry remove %s", err)
	}

	// 销毁测试聊天室
	_, err = rc.ChatroomDestroy(ctx, &rongcloud.ChatroomDestroyRequest{
		ChatroomIds: []string{grp},
	})
	if err != nil {
		log.Fatalf("chatroom destroy error: %s", err)
	}
}
