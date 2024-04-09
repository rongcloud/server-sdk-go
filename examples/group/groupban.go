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

	// 创建群组
	_, err := rc.GroupCreate(ctx, &rongcloud.GroupCreateRequest{
		UserId:    []string{"u01", "u02"},
		GroupId:   rongcloud.StringPtr("grp01"),
		GroupName: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group create error %s", err)
	}

	// 设置群组全体禁言
	_, err = rc.GroupBanAdd(ctx, &rongcloud.GroupBanAddRequest{
		GroupIds: []string{"grp01", "grp02"},
	})
	if err != nil {
		log.Fatalf("group ban add error %s", err)
	}

	// 查询群组全体禁言
	groupBanQueryResp, err := rc.GroupBanQuery(ctx, &rongcloud.GroupBanQueryRequest{
		GroupIds: []string{"grp01", "grp02"},
	})
	if err != nil {
		log.Fatalf("group ban query error %s", err)
	}
	groupBanQueryRespData, _ := json.Marshal(groupBanQueryResp)
	log.Printf("group ban query response data: %s", groupBanQueryRespData)

	// 取消群组全体禁言
	_, err = rc.GroupBanRollback(ctx, &rongcloud.GroupBanRollbackRequest{
		GroupIds: []string{"grp01", "grp02"},
	})
	if err != nil {
		log.Fatalf("group ban rollback error %s", err)
	}

	// 加入群组全体禁言白名单
	_, err = rc.GroupUserBanWhitelistAdd(ctx, &rongcloud.GroupUserBanWhitelistAddRequest{
		UserIds: []string{"u01", "u02"},
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group user ban white list add error %s", err)
	}

	// 查询群组全体禁言白名单
	groupUserBanWhitelistQueryResp, err := rc.GroupUserBanWhitelistQuery(ctx, &rongcloud.GroupUserBanWhitelistQueryRequest{
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group user ban whitelist query error %s", err)
	}
	groupUserBanWhitelistQueryRespData, _ := json.Marshal(groupUserBanWhitelistQueryResp)
	log.Printf("group user ban white list query resp data: %s", groupUserBanWhitelistQueryRespData)

	// 移除群组全体禁言白名单
	_, err = rc.GroupUserBanWhitelistRollback(ctx, &rongcloud.GroupUserBanWhitelistRollbackRequest{
		UserIds: []string{"u01", "u02"},
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group user ban whitelist rollback error %s", err)
	}
}
