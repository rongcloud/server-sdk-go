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

	// 设置群成员推送备注名
	_, err = rc.GroupRemarksSet(ctx, &rongcloud.GroupRemarksSetRequest{
		UserId:  rongcloud.StringPtr("u01"),
		GroupId: rongcloud.StringPtr("grp01"),
		Remark:  rongcloud.StringPtr("u01Remarks"),
	})
	if err != nil {
		log.Fatalf("group remakrs set error %s", err)
	}

	// 查询群成员推送备注名
	remarkGetResp, err := rc.GroupRemarksGet(ctx, &rongcloud.GroupRemarksGetRequest{
		UserId:  rongcloud.StringPtr("u01"),
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group remarks get error %s", err)
	}
	remarkGetRespData, _ := json.Marshal(remarkGetResp)
	log.Printf("group remarks get resp data: %s", remarkGetRespData)

	//  删除群成员推送备注名
	_, err = rc.GroupRemarksDel(ctx, &rongcloud.GroupRemarksDelRequest{
		UserId:  rongcloud.StringPtr("u01"),
		GroupId: rongcloud.StringPtr("grp01"),
	})
	if err != nil {
		log.Fatalf("group remark del error %s", err)
	}
}
