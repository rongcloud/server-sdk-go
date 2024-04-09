package main

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	requestId := uuid.New().String()
	// 自定义requestId
	rongcloud.AddHttpRequestId(ctx, requestId)

	resp, err := rc.UserGetToken(ctx, &rongcloud.UserGetTokenRequest{
		UserId:      rongcloud.StringPtr("u01"),
		Name:        rongcloud.StringPtr("u01"),
		PortraitUri: nil,
	})
	if err != nil {
		log.Fatalf("user get token error %s", err)
	}
	// 获取http.Response
	httpResp := resp.GetHttpResponse()
	// 快捷方法获取x-request-id

	requestIdResp := resp.GetRequestId()
	log.Printf("http response %+v, requestId: %s, requestIdResp: %s", httpResp, requestId, requestIdResp)
}
