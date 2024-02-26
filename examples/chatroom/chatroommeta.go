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
