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

	// 设置用户标签
	_, err := rc.UserTagSet(ctx, &rongcloud.UserTagSetRequest{
		UserId: rongcloud.StringPtr("u01"),
		Tags:   []string{"tag1", "tag2"},
	})
	if err != nil {
		log.Fatalf("user tag set error %s", err)
	}

	// 批量设置用户标签
	_, err = rc.UserTagBatchSet(ctx, &rongcloud.UserTagBatchSetRequest{
		UserIds: []string{"u01", "u02"},
		Tags:    []string{"tag3", "tag4"},
	})
	if err != nil {
		log.Fatalf("user tag batch set error %s", err)
	}

	// 获取用户标签
	userTagsGetResp, err := rc.UserTagsGet(ctx, &rongcloud.UserTagsGetRequest{
		UserIds: []string{"u01", "u02"},
	})
	if err != nil {
		log.Fatalf("user tags get error %s", err)
	}
	userTagsGetRespData, _ := json.Marshal(userTagsGetResp)
	log.Printf("user tags get response data: %s", userTagsGetRespData)
}
