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

	chatroomMetaOp(ctx, rc)
}

func chatroomMetaOp(ctx context.Context, rc *rongcloud.RongCloud) {
	// just for test
	rc.Setting.DisableCodeCheck = false

	// create chatroom
	resp, err := rc.ChatroomCreateNew(ctx, &rongcloud.ChatroomCreateNewRequest{
		ChatroomId: rongcloud.StringPtr("grp1"),
	})
	if err != nil {
		log.Fatalf("chatroom create err %s", err)
	}
	b, _ := json.Marshal(resp)
	log.Printf("chat room create resp: %s", b)
	log.Printf("http response get %+v", resp.GetHttpResponse())

	// chatroom destroy settings
	destroySetResp, err := rc.ChatroomDestroySet(ctx, &rongcloud.ChatroomDestroySetRequest{
		ChatroomId:  rongcloud.StringPtr("grp1"),
		DestroyTime: rongcloud.IntPtr(60),
		DestroyType: rongcloud.IntPtr(1),
	})
	if err != nil {
		log.Fatalf("chatroom destroy set err %s", err)
	}
	destroySetData, _ := json.Marshal(destroySetResp)
	log.Printf("chatroom destroy set resp: %s", destroySetData)

	// get a chatroom information
	chatroomGetResp, err := rc.ChatroomGet(ctx, &rongcloud.ChatroomGetRequest{ChatroomId: rongcloud.StringPtr("grp1")})
	if err != nil {
		log.Fatalf("chatroom get err %s", err)
	}
	chatroomGetData, _ := json.Marshal(chatroomGetResp)
	log.Printf("chatroom get resp: %s", chatroomGetData)

	// 设置聊天室属性（KV）
	chatroomEntrySetResp, err := rc.ChatroomEntrySet(ctx, &rongcloud.ChatroomEntrySetRequest{
		ChatroomId: rongcloud.StringPtr("grp1"),
		UserId:     rongcloud.StringPtr("user01"),
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
	chatroomEntrySetData, _ := json.Marshal(chatroomEntrySetResp)
	log.Printf("chatroom entry set resp: %s", chatroomEntrySetData)

	// 批量设置聊天室属性（KV）
	entryBatchSetResp, err := rc.ChatroomEntryBatchSet(ctx, &rongcloud.ChatroomEntryBatchSetRequest{
		ChatroomId:   rongcloud.StringPtr("grp1"),
		AutoDelete:   rongcloud.IntPtr(1),
		EntryOwnerId: rongcloud.StringPtr("u01"),
		EntryInfo:    map[string]string{"k1": "k2"},
	})
	if err != nil {
		log.Fatalf("chatroom entry batch set err %s", err)
	}
	entryBatchSetData, _ := json.Marshal(entryBatchSetResp)
	log.Printf("chat room entry batch set resp: %s", entryBatchSetData)

	// destroy chatroom 销毁聊天室
	destroyChatroomResp, err := rc.ChatroomDestroy(ctx, &rongcloud.ChatroomDestroyRequest{
		ChatroomIds: []string{"grp1", "grp2"},
	})
	if err != nil {
		log.Fatalf("chatroom destroy err %s", err)
	}
	destroyChatRoomData, _ := json.Marshal(destroyChatroomResp)
	log.Printf("chatroom destroy: %s", destroyChatRoomData)
}
