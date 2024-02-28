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

	// 禁言指定群成员
	_, err := rc.GroupUserGagAdd(ctx, &rongcloud.GroupUserGagAddRequest{
		UserId:  []string{"u01", "u02"},
		GroupId: rongcloud.StringPtr("grp01"),
		Minute:  rongcloud.IntPtr(30),
	})
	if err != nil {
		log.Fatalf("group user gaga add error %s", err)
	}

	// 查询群成员禁言列表
	groupUserGagListResp, err := rc.GroupUserGagList(ctx, &rongcloud.GroupUserGagListRequest{
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group user gag list error %s", err)
	}
	groupUserGagListRespData, _ := json.Marshal(groupUserGagListResp)
	log.Printf("group user gag list resp data: %s", groupUserGagListRespData)

	// 取消指定群成员禁言
	_, err = rc.GroupUserGagRollback(ctx, &rongcloud.GroupUserGagRollbackRequest{
		UserId:  []string{"u01", "u02"},
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group user gag rollback error %s", err)
	}
}
