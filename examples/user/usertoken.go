package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	// 注册用户
	getTokenResp, err := rc.UserGetToken(ctx, &rongcloud.UserGetTokenRequest{
		UserId:      rongcloud.StringPtr("uu01"),
		Name:        rongcloud.StringPtr("uu01"),
		PortraitUri: rongcloud.StringPtr("http://a.b.com/a.jpg"),
	})
	if err != nil {
		log.Fatalf("user get token error %s", err)
	}
	getTokenRespData, _ := json.Marshal(getTokenResp)
	log.Printf("use get token response data: %s", getTokenRespData)

	// 作废token
	_, err = rc.UserTokenExpire(ctx, &rongcloud.UserTokenExpireRequest{
		UserId: []string{"uu01"},
		Time:   rongcloud.Int64Ptr(time.Now().Add(time.Hour * 10).UnixMilli()),
	})
	if err != nil {
		log.Fatalf("user token expire error %s", err)
	}
}
