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

	// 添加用户到白名单
	_, err := rc.UserWhitelistAdd(ctx, &rongcloud.UserWhitelistAddRequest{
		UserId:      rongcloud.StringPtr("u01"),
		WhiteUserId: []string{"u02", "u03"},
	})
	if err != nil {
		log.Fatalf("user whitelist add error %s", err)
	}

	// 查询白名单中用户列表
	whitelistQueryResp, err := rc.UserWhitelistQuery(ctx, &rongcloud.UserWhitelistQueryRequest{
		UserId: rongcloud.StringPtr("u01"),
	})
	if err != nil {
		log.Fatalf("user whitelist query error %s", err)
	}
	whitelistQueryRespData, _ := json.Marshal(whitelistQueryResp)
	log.Printf("user whitelist query response data: %s", whitelistQueryRespData)

	// 移除白名单中用户
	_, err = rc.UserWhitelistRemove(ctx, &rongcloud.UserWhitelistRemoveRequest{
		UserId:      rongcloud.StringPtr("u01"),
		WhiteUserId: []string{"u02", "u03"},
	})
	if err != nil {
		log.Fatalf("user whitelist remove error %s", err)
	}
}
