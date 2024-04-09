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

	// 添加用户到黑名单
	_, err := rc.UserBlacklistAdd(ctx, &rongcloud.UserBlacklistAddRequest{
		UserId:      rongcloud.StringPtr("u01"),
		BlackUserId: []string{"u02", "u03"},
	})
	if err != nil {
		log.Fatalf("user black list add error %s", err)
	}

	// 获取黑名单用户列表
	userBlackListQueryResp, err := rc.UserBlacklistQuery(ctx, &rongcloud.UserBlacklistQueryRequest{
		UserId: rongcloud.StringPtr("u01"),
	})
	if err != nil {
		log.Fatalf("user blacklist query error %s", err)
	}
	userBlackListQueryRespData, _ := json.Marshal(userBlackListQueryResp)
	log.Printf("user blacklist query response data: %s", userBlackListQueryRespData)

	//  移除黑名单中用户
	_, err = rc.UserBlacklistRemove(ctx, &rongcloud.UserBlacklistRemoveRequest{
		UserId:      rongcloud.StringPtr("u01"),
		BlackUserId: []string{"u02", "u03"},
	})
	if err != nil {
		log.Fatalf("user blacklist remove error %s", err)
	}
}
