package rongcloud

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMessagePrivatePublishRequest_MarshalJson(t *testing.T) {
	txtMsg := &TXTMsg{
		Content: "hello,world",
		User: MsgUserInfo{
			ID:       "userinfo_id",
			Name:     "userinfo_name",
			Icon:     "userinfo_icon",
			Portrait: "userinfo_portrait",
			Extra:    "userinfo_extra",
		},
		Extra: "extra",
	}
	b, err := json.Marshal(txtMsg)
	if err != nil {
		t.Fatalf("txt msg fail: %s", err)
	}
	t.Logf("txt msg %s", b)
}

func TestRongCloud_MessagePrivatePublish(t *testing.T) {
	// test txt msg
	ctx := context.TODO()
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	txtMsg := &TXTMsg{
		Content: "hello world",
	}
	resp, err := rc.MessagePrivatePublish(ctx, &MessagePrivatePublishRequest{
		FromUserId: String("u01"),
		ToUserId:   String("u02"),
		RCMsg:      txtMsg,
	})
	if err != nil {
		t.Fatalf("msg private publish err: %s", err)
	}
	if resp.Code != 200 {
		t.Fatalf("msg private publish non 200 code")
	}
	httpResponse := resp.GetHttpResponse()
	t.Logf("message private publish resp: %+v, raw http request %+v response: %+v", resp, httpResponse.Request, httpResponse)

	// test HQVSMsg
	customRequestId := uuid.New().String()
	ctx = AddHttpRequestId(ctx, customRequestId)
	hqResp, err := rc.MessagePrivatePublish(ctx, &MessagePrivatePublishRequest{
		FromUserId: String("u01"),
		ToUserId:   String("u02"),
		RCMsg: &HQVCMsg{
			RemoteUrl: "1",
			Duration:  2,
		},
	})
	if err != nil {
		t.Fatalf("hqvc msg private err: %s", err)
	}
	if resp.Code != 200 {
		t.Fatalf("hqvc msg private publish non 200 code")
	}
	hqHttpResp := hqResp.GetHttpResponse()
	httpRequestId := hqResp.GetRequestId()
	t.Logf("msg private publish resp: %+v, raw http request: %+v, response: %+v", hqResp, hqHttpResp.Request, hqHttpResp)
	assert.Equal(t, customRequestId, httpRequestId, "customRequestId should equal httpResponse requestId")
}
