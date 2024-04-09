package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	rc := rongcloud.NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	ctx := context.Background()

	// 注销用户
	_, err := rc.UserDeactivate(ctx, &rongcloud.UserDeactivateRequest{
		UserId: rongcloud.StringPtr(strings.Join([]string{"u01", "u02"}, ",")),
	})
	if err != nil {
		log.Fatalf("user deactivate error %s", err)
	}

	// 查询已注销用户
	userDeactivateQueryResp, err := rc.UserDeactivateQuery(ctx, &rongcloud.UserDeactivateQueryRequest{
		PageNo:   rongcloud.IntPtr(1),
		PageSize: rongcloud.IntPtr(20),
	})
	if err != nil {
		log.Fatalf("user deactivate query error %s", err)
	}
	userDeactivateQueryRespData, _ := json.Marshal(userDeactivateQueryResp)
	log.Printf("user deactivate query response data: %s", userDeactivateQueryRespData)

	// 重新激活用户 ID
	_, err = rc.UserReactivate(ctx, &rongcloud.UserReactivateRequest{
		UserId: rongcloud.StringPtr(strings.Join([]string{"u01", "u02"}, ",")),
	})
	if err != nil {
		log.Fatalf("use reactivate error %s", err)
	}
}
