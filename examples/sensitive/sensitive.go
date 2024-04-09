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

	//  添加消息敏感词
	_, err := rc.SensitiveWordAdd(ctx, &rongcloud.SensitiveWordAddRequest{
		Word:        rongcloud.StringPtr("money"),
		ReplaceWord: rongcloud.StringPtr("***"),
	})
	if err != nil {
		log.Fatalf("sensitive word add error %s", err)
	}

	// 查询消息敏感词
	listSensitiveWordResp, err := rc.SensitiveWordList(ctx, &rongcloud.SensitiveWordListRequest{
		Type: rongcloud.StringPtr("2"),
	})
	if err != nil {
		log.Fatalf("sensitive word list error %s", err)
	}
	listSensitiveWordRespData, _ := json.Marshal(listSensitiveWordResp)
	log.Printf("sensitive word list response data: %s", listSensitiveWordRespData)

	// 批量移除消息敏感词
	_, err = rc.SensitiveWordBatchDelete(ctx, &rongcloud.SensitiveWordBatchDeleteRequest{
		Words: []string{"money"},
	})
	if err != nil {
		log.Fatalf("sensitive word batch delete error %s", err)
	}
}
