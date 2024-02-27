server-sdk-go
=============

Rong Cloud Server SDK in Go.

# 版本说明
当前版本为v4, 后续版本会保持向前兼容, 但不兼容旧的v3，v2等版本, 旧版本不再更新, 请升级后使用。

# API文档
- [官方文档](https://doc.rongcloud.cn/imserver/server/v1/overview)

## 如何使用

```go
package main

import (
    "context"
    "fmt"

    "github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
    rc := rongcloud.NewRongCloud("appKey", "appSecret")
    // 单聊文本消息
    txtMsg := &rongcloud.TXTMsg{
        Content: "hello world",
    }
    ctx := context.Background()
    
    // 自定义requestId
    requestId := uuid.New().String()
    rongcloud.AddHttpRequestId(ctx, requestId)
    
    // 发送单聊消息
    resp, err := rc.MessagePrivatePublish(ctx, &rongcloud.MessagePrivatePublishRequest{
        FromUserId: rongcloud.StringPtr("u01"),
        ToUserId:   rongcloud.StringPtr("u02"),
        RCMsg:      txtMsg,
    })
    fmt.Println(err)
    
    // 获取http.Response
    httpResp := resp.GetHttpResponse()
    
    // 快捷方法获取x-request-id
    requestIdResp := resp.GetRequestId()
}
```

> 更多代码示例请参考 [examples](./examples/README.md) 

### http 参数优化

- http连接相关的性能优化
- `rongcloud.WithMaxIdleConnsPerHost` : 每个域名最大活跃连接数，默认 100.
- `rongcloud.WithTimeout` : 连接超时设置，默认 10 秒；最小单位为秒， `sdk.WithTimeout(30*time.Second)` 表示设置为30秒.
- `rongcloud.WithKeepAlive` : 连接保活时间，默认 30 秒；最小单位为秒， `sdk.WithKeepAlive(30*time.Second)` 表示设置保活时间为30秒.
- `rongcloud.WithTransport`: 设置 http client transport 参数，优先级大于其他http参数.
- `rc.SetHttpTransport` : 实例上设置 http client transport 参数，优先级大于其他http参数.
- `rc.GetHttpTransport` : 获得当前实例http client transport.

```go
package main

import "time"
import "net"
import "net/http"
import "github.com/rongcloud/server-sdk-go/v4/rongcloud"

func main() {
	// 方法1： 创建对象时设置
	rc := rongcloud.NewRongCloud(
        "appKey",
		"appSecret",
		// 每个域名最大活跃连接数
		rongcloud.WithMaxIdleConnsPerHost(100),
		)
	
	// 方法2： 自定义 http client， 调用 set 方法设置
	dialer := &net.Dialer{
        Timeout:   10 * time.Second,
        KeepAlive: 30 * time.Second,
    }
    globalTransport := &http.Transport{
        DialContext:         dialer.DialContext,
        MaxIdleConnsPerHost: 100,
    }
    rc.SetHttpTransport(globalTransport)
	
}
```

### GO SDK 功能支持的版本清单

* [消息发送](./examples/message/README.md)
* [用户管理](./examples/user/README.md)
* [聊天室](./examples/chatroom/README.md)
* [群组](./examples/group/README.md)
* [超级群](./examples/ultragroup/README.md)
* [会话](./examples/conversation/README.md)
* [推送](./examples/push/README.md)
* [敏感词](./examples/sensitive/README.md)