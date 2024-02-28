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

	// 设置用户级推送备注名
	_, err := rc.UserRemarksSet(ctx, &rongcloud.UserRemarksSetRequest{
		UserId: rongcloud.StringPtr("u01"),
		Remarks: []*rongcloud.UserRemark{
			{
				Id:     "u02",
				Remark: "alias-u02",
			}, {
				Id:     "u03",
				Remark: "alias-u03",
			},
		},
	})
	if err != nil {
		log.Fatalf("user remarks set error %s", err)
	}

	// 查询用户级推送备注名
	userRemarksGetResp, err := rc.UserRemarksGet(ctx, &rongcloud.UserRemarksGetRequest{
		UserId: rongcloud.StringPtr("u01"),
		Page:   rongcloud.IntPtr(1),
		Size:   rongcloud.IntPtr(50),
	})
	if err != nil {
		log.Fatalf("user remarks get error %s", err)
	}
	userRemarksGetRespData, _ := json.Marshal(userRemarksGetResp)
	log.Printf("user remarks get response data: %s", userRemarksGetRespData)

	// 删除用户级推送备注名
	_, err = rc.UserRemarksDel(ctx, &rongcloud.UserRemarksDelRequest{
		UserId:   rongcloud.StringPtr("u01"),
		TargetId: rongcloud.StringPtr("u02"),
	})
	if err != nil {
		log.Fatalf("user remarks del error %s", err)
	}
}
