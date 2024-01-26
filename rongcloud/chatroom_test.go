package rongcloud

import (
	"context"
	"encoding/json"
	"os"
	"testing"
)

func testNewRongCloud() *RongCloud {
	return NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
}

func TestRongCloud_ChatroomCreateNew(t *testing.T) {
	ctx := context.TODO()
	rc := testNewRongCloud()
	rc.Setting.DisableCodeCheck = true
	resp, err := rc.ChatroomCreateNew(ctx, &ChatroomCreateNewRequest{
		ChatroomId: String("grp1"),
	})
	if err != nil {
		t.Fatalf("chatroom create err %s", err)
	}
	b, _ := json.Marshal(resp)
	t.Logf("chat room create resp: %s", b)
	t.Logf("http response get %+v", resp.GetHttpResponse())

	destroySetResp, err := rc.ChatroomDestroySet(ctx, &ChatroomDestroySetRequest{
		ChatroomId:  String("grp1"),
		DestroyTime: Int(60),
		DestroyType: Int(1),
	})
	if err != nil {
		t.Fatalf("chatroom destroy set err %s", err)
	}
	destroySetData, _ := json.Marshal(destroySetResp)
	t.Logf("chatroom destroy set resp: %s", destroySetData)

	chatroomGetResp, err := rc.ChatroomGet(ctx, &ChatroomGetRequest{ChatroomId: String("grp1")})
	if err != nil {
		t.Fatalf("chatroom get err %s", err)
	}
	chatroomGetData, _ := json.Marshal(chatroomGetResp)
	t.Logf("chatroom get resp: %s", chatroomGetData)

	chatroomEntrySetResp, err := rc.ChatroomEntrySet(ctx, &ChatroomEntrySetRequest{
		ChatroomId: String("grp1"),
		UserId:     String("user01"),
		Key:        String("key1"),
		Value:      String("val1"),
		AutoDelete: Int(1),
		RCMsg: &ChrmKVNotiMsg{
			Type:  1,
			Key:   "key1",
			Value: "val1",
			Extra: "extra info",
		},
	})
	if err != nil {
		t.Fatalf("chatroom entry set err %s", err)
	}
	chatroomEntrySetData, _ := json.Marshal(chatroomEntrySetResp)
	t.Logf("chatroom entry set resp: %s", chatroomEntrySetData)
	entryBatchSetResp, err := rc.ChatroomEntryBatchSet(ctx, &ChatroomEntryBatchSetRequest{
		ChatroomId:   String("grp1"),
		AutoDelete:   Int(1),
		EntryOwnerId: String("u01"),
		EntryInfo:    map[string]string{"k1": "k2"},
	})
	if err != nil {
		t.Fatalf("chatroom entry batch set err %s", err)
	}
	entryBatchSetData, _ := json.Marshal(entryBatchSetResp)
	t.Logf("chat room entry batch set resp: %s", entryBatchSetData)

	// destroy chatroom
	destroyChatroomResp, err := rc.ChatroomDestroy(ctx, &ChatroomDestroyRequest{
		ChatroomIds: []string{"grp1", "grp2"},
	})
	if err != nil {
		t.Fatalf("chatroom destroy err %s", err)
	}
	destroyChatRoomData, _ := json.Marshal(destroyChatroomResp)
	t.Logf("chatroom destroy: %s", destroyChatRoomData)
}

func TestRongCloud_ChatroomUserExist(t *testing.T) {
	rc := testNewRongCloud()
	resp, err := rc.ChatroomUserExist(context.TODO(), &ChatroomUserExistRequest{})
	if err != nil {
		t.Fatalf("chatroom user exit err %s", err)
	}
	b, _ := json.Marshal(resp)
	t.Logf("chatroom user exist resp: %s", b)
	t.Logf("http resposne get resp: %+v", resp.GetHttpResponse())
}

func TestRongCloud_ChatroomUserQuery(t *testing.T) {
	rc := testNewRongCloud()
	ctx := context.TODO()
	resp, err := rc.ChatroomUserQuery(ctx, &ChatroomUserQueryRequest{
		ChatroomId: String("grp1"),
		Count:      Int(200),
		Order:      Int(1),
	})
	if err != nil {
		t.Fatalf("chatroom user query err %s", err)
	}
	data, _ := json.Marshal(resp)
	t.Logf("chat room user query: %s", data)
}

func TestRongCloud_ChatroomUsersExist(t *testing.T) {
	rc := testNewRongCloud()
	ctx := context.TODO()
	resp, err := rc.ChatroomUsersExist(ctx, &ChatroomUsersExistRequest{
		ChatroomId: String("grp1"),
		UserIds:    []string{"u01", "u02"},
	})
	if err != nil {
		t.Fatalf("chatroom users exist err %s", err)
	}
	respData, _ := json.Marshal(resp)
	t.Logf("chatroom users exist resp: %s", respData)
}

func TestRongCloud_ChatroomUserBlockAdd(t *testing.T) {
	rc := testNewRongCloud()
	ctx := context.TODO()
	resp, err := rc.ChatroomUserBlockAdd(ctx, &ChatroomUserBlockAddRequest{
		UserIds:    []string{"u01", "u02"},
		ChatroomId: String("grp1"),
		Minute:     Int(10),
		Extra:      String(""),
		NeedNotify: Bool(true),
	})
	if err != nil {
		t.Fatalf("chatroom user block add err %s", err)
	}
	d, _ := json.Marshal(resp)
	t.Logf("chatroom user block data: %s", d)

	rollbackResp, err := rc.ChatroomUserBlockRollback(ctx, &ChatroomUserBlockRollbackRequest{
		UserIds:    []string{"u01", "u02"},
		ChatroomId: String("grp1"),
		Extra:      String(""),
		NeedNotify: Bool(true),
	})
	if err != nil {
		t.Fatalf("chatroom user block rollback err %s", err)
	}
	rollbackData, _ := json.Marshal(rollbackResp)
	t.Logf("rollback data: %s", rollbackData)
	userBlockListResp, err := rc.ChatroomUserBlockList(ctx, &ChatroomUserBlockListRequest{ChatroomId: String("grp1")})
	if err != nil {
		t.Fatalf("chatroom user block list err: %s", err)
	}
	userBlockListData, _ := json.Marshal(userBlockListResp)
	t.Logf("user block list data: %s", userBlockListData)

	userBanAddResp, err := rc.ChatroomUserBanAdd(ctx, &ChatroomUserBanAddRequest{
		UserIds:    []string{"u01", "u02"},
		Minute:     Int(10),
		Extra:      String(""),
		NeedNotify: Bool(true),
	})
	if err != nil {
		t.Fatalf("user ban add err %s", err)
	}
	userBanAddData, _ := json.Marshal(userBanAddResp)
	t.Logf("user ban add data: %s", userBanAddData)
	userBanRemoveResp, err := rc.ChatroomUserBanRemove(ctx, &ChatroomUserBanRemoveRequest{
		UserIds:    []string{"u01", "u02"},
		Extra:      String(""),
		NeedNotify: Bool(true),
	})
	if err != nil {
		t.Fatalf("user ban remove data: %s", err)
	}
	userBanRemoveData, _ := json.Marshal(userBanRemoveResp)
	t.Logf("user ban remove data: %s", userBanRemoveData)

	userBanQueryResp, err := rc.ChatroomUserBanQuery(ctx, nil)
	if err != nil {
		t.Fatalf("user ban query err %s", err)
	}
	userBanQueryData, _ := json.Marshal(userBanQueryResp)
	t.Logf("user ban query data: %s", userBanQueryData)
}

func TestRongCloud_ChatroomUserGagAdd(t *testing.T) {
	rc := testNewRongCloud()
	ctx := context.TODO()
	userGagAddResp, err := rc.ChatroomUserGagAdd(ctx, &ChatroomUserGagAddRequest{
		UserIds:    []string{"u01", "u02"},
		ChatroomId: String("grp1"),
		Minute:     Int(10),
		Extra:      String(""),
		NeedNotify: Bool(true),
	})
	if err != nil {
		t.Fatalf("chatroom user gag add err %s", err)
	}
	userGagAddData, _ := json.Marshal(userGagAddResp)
	t.Logf("chatroom user gag add data: %s", userGagAddData)
	userGagListResp, err := rc.ChatroomUserGagList(ctx, &ChatroomUserGagListRequest{
		ChatroomId: String("grp1"),
	})
	if err != nil {
		t.Fatalf("chatroom user gag list err %s", err)
	}
	userGagListData, _ := json.Marshal(userGagListResp)
	t.Logf("chatroom user gag list data: %s", userGagListData)
	userGagRollbackResp, err := rc.ChatroomUserGagRollback(ctx, &ChatroomUserGagRollbackRequest{
		UserIds:    []string{"u01", "u02"},
		ChatroomId: String("grp1"),
		Extra:      String(""),
		NeedNotify: Bool(true),
	})
	if err != nil {
		t.Fatalf("chatroom user gag rollback err %s", err)
	}
	userGagRollbackData, _ := json.Marshal(userGagRollbackResp)
	t.Logf("chatroom user gag rollback data: %s", userGagRollbackData)
}

func TestRongCloud_ChatroomMessagePriorityRemove(t *testing.T) {
	rc := testNewRongCloud()
	ctx := context.TODO()
	// first list message priority
	query1Resp, err := rc.ChatroomMessagePriorityQuery(ctx, nil)
	if err != nil {
		t.Fatalf("chatroom message priority query err: %s", err)
	}
	query1Data, _ := json.Marshal(query1Resp)
	t.Logf("chatroom query response data %s", query1Data)

	// add message priority
	vcMsg := "RC:VcMsg"
	imgMsg := "RC:ImgMsg"
	addResp, err := rc.ChatroomMessagePriorityAdd(ctx, &ChatroomMessagePriorityAddRequest{
		ObjectNames: []string{vcMsg, imgMsg},
	})
	if err != nil {
		t.Fatalf("chatroom message priority add err %s", err)
	}
	addData, _ := json.Marshal(addResp)
	t.Logf("chatroom message priority add resp: %s", addData)

	// query after add
	query2Resp, err := rc.ChatroomMessagePriorityQuery(ctx, nil)
	if err != nil {
		t.Fatalf("chatroom message priority query err: %s", err)
	}
	query2Data, _ := json.Marshal(query2Resp)
	t.Logf("chatroom query response data %s", query2Data)

	// remove message priority
	removeResp, err := rc.ChatroomMessagePriorityRemove(ctx, &ChatroomMessagePriorityRemoveRequest{
		ObjectNames: []string{vcMsg, imgMsg},
	})
	if err != nil {
		t.Fatalf("chatroom message priority remove err %s", err)
	}
	removeData, _ := json.Marshal(removeResp)
	t.Logf("chatroom message priority remove err: %s", removeData)

	// query again after remove
	query3Resp, err := rc.ChatroomMessagePriorityQuery(ctx, nil)
	if err != nil {
		t.Fatalf("chatroom message priority query err: %s", err)
	}
	query3Data, _ := json.Marshal(query3Resp)
	t.Logf("chatroom query response data %s", query3Data)
}
