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
	// 创建超级群
	_, err := rc.UltraGroupCreate(ctx, &rongcloud.UltraGroupCreateRequest{
		UserId:    rongcloud.StringPtr("u01"),
		GroupId:   rongcloud.StringPtr("ug01"),
		GroupName: rongcloud.StringPtr("ug01"),
	})
	if err != nil {
		log.Fatalf("ultra group create error %s", err)
	}

	// 创建频道
	_, err = rc.UltraGroupChannelCreate(ctx, &rongcloud.UltraGroupChannelCreateRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		Type:       rongcloud.IntPtr(0),
	})
	if err != nil {
		log.Fatalf("ultra group channel create error %s", err)
	}

	// 创建用户组
	_, err = rc.UltraGroupUserGroupAdd(ctx, &rongcloud.UltraGroupUserGroupAddRequest{
		GroupId: rongcloud.StringPtr("ug01"),
		UserGroups: []*rongcloud.UltraGroupUserGroup{
			{
				UserGroupId: "usg01",
			}, {
				UserGroupId: "usg02",
			},
		},
	})
	if err != nil {
		log.Fatalf("ultra group user group add error %s", err)
	}

	// 查询用户组列表
	userGroupQueryResp, err := rc.UltraGroupUserGroupQuery(ctx, &rongcloud.UltraGroupUserGroupQueryRequest{
		GroupId:  rongcloud.StringPtr("ug01"),
		Page:     rongcloud.IntPtr(1),
		PageSize: rongcloud.IntPtr(10),
	})
	if err != nil {
		log.Fatalf("ultra group user group query error %s", err)
	}
	userGroupQueryRespData, _ := json.Marshal(userGroupQueryResp)
	log.Printf("ultra group user query resp data: %s", userGroupQueryRespData)

	// 超级群绑定频道与用户组
	_, err = rc.UltraGroupChannelUserGroupBind(ctx, &rongcloud.UltraGroupChannelUserGroupBindRequest{
		GroupId:      rongcloud.StringPtr("ug01"),
		BusChannel:   rongcloud.StringPtr("channel01"),
		UserGroupIds: rongcloud.StringPtr("usg01,usg02"),
	})
	if err != nil {
		log.Fatalf("ultra group channel userGroup bind error %s", err)
	}

	// 超级群查询频道绑定的用户组
	channelUserGroupQueryResp, err := rc.UltraGroupChannelUserGroupQuery(ctx, &rongcloud.UltraGroupChannelUserGroupQueryRequest{
		GroupId:    rongcloud.StringPtr("ug01"),
		BusChannel: rongcloud.StringPtr("channel01"),
		Page:       rongcloud.IntPtr(1),
		PageSize:   rongcloud.IntPtr(10),
	})
	if err != nil {
		log.Fatalf("ultra group channel userGroup query error %s", err)
	}
	channelUserGroupQueryRespData, _ := json.Marshal(channelUserGroupQueryResp)
	log.Printf("ultra group channel userGroup query data %s", channelUserGroupQueryRespData)

	// 超级群查询用户组绑定的频道
	userGroupChannelQueryResp, err := rc.UltraGroupUserGroupChannelQuery(ctx, &rongcloud.UltraGroupUserGroupChannelQueryRequest{
		GroupId:     rongcloud.StringPtr("ug01"),
		UserGroupId: rongcloud.StringPtr("usg01"),
		Page:        rongcloud.IntPtr(1),
		PageSize:    rongcloud.IntPtr(10),
	})
	if err != nil {
		log.Fatalf("ultra group userGroup channel query error %s", err)
	}
	userGroupChannelQueryRespData, _ := json.Marshal(userGroupChannelQueryResp)
	log.Printf("ultra group userGroup channel query data %s", userGroupChannelQueryRespData)

	// 超级群解绑频道与用户组
	_, err = rc.UltraGroupChannelUserGroupUnbind(ctx, &rongcloud.UltraGroupChannelUserGroupUnbindRequest{
		GroupId:      rongcloud.StringPtr("ug01"),
		BusChannel:   rongcloud.StringPtr("channel01"),
		UserGroupIds: rongcloud.StringPtr("usg01,usg02"),
	})
	if err != nil {
		log.Fatalf("ultra group channel userGroup unbind error %s", err)
	}

	// 删除用户组
	_, err = rc.UltraGroupUserGroupDel(ctx, &rongcloud.UltraGroupUserGroupDelRequest{
		GroupId:      rongcloud.StringPtr("ug01"),
		UserGroupIds: rongcloud.StringPtr("usg01,usg02"),
	})
	if err != nil {
		log.Fatalf("ultra group user group del error %s", err)
	}

	// 解散超级群
	_, err = rc.UltraGroupDis(ctx, &rongcloud.UltraGroupDisRequest{GroupId: rongcloud.StringPtr("ug01")})
	if err != nil {
		log.Fatalf("ultra group dis error %s", err)
	}
}
