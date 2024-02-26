server-sdk-go
=============

Rong Cloud Server SDK in Go.

# 版本说明

- 为方便开发者的接入使用，更好的对融云 Server SDK 进行维护管理，融云 Server SDK v3 统一规范了命名及调用方式，结构更加清晰。老版本的 Server SDK 已经切入 v1 v2.0.1 分支，仍然可以使用，但不会再做新的功能更新。
- 如果您是新接入的开发者，建议您使用 Server SDK v3 版本。 对已集成使用老版本 Server SDK 的开发者，不能直接升级使用，强烈建议您重新测试后使用

# API文档
- [官方文档](https://doc.rongcloud.cn/imserver/server/v1/overview)

## 如何使用

```go
package main

import (
	"fmt"
    
    "github.com/rongcloud/server-sdk-go/v4/rongcloud"
)

func main() {
	rc := rongcloud.NewRongCloud("appKey", "appSecret")
    
    // 发送单聊文本消息
    txtMsg := &rongcloud.TXTMsg{
        Content: "hello world",
    }
    requestId := uuid.New().String()
    rongcloud.AddHttpRequestId(ctx, requestId)
    _, err := rc.MessagePrivatePublish(ctx, &rongcloud.MessagePrivatePublishRequest{
        FromUserId: rongcloud.StringPtr("u01"),
        ToUserId:   rongcloud.StringPtr("u02"),
        RCMsg:      txtMsg,
    })
	
	fmt.Println(err)
}
```

> 更多代码示例请参考 [examples](./examples) 

### http 参数优化

- http连接相关的性能优化
- `sdk.WithMaxIdleConnsPerHost` : 每个域名最大活跃连接数，默认 100
- `sdk.WithTimeout` : 连接超时设置，默认 10 秒；最小单位为秒， `sdk.WithTimeout(30)` 表示设置为30秒
- `sdk.WithKeepAlive` : 连接保活时间，默认 30 秒；最小单位为秒， `sdk.WithKeepAlive(30)` 表示设置保活时间为30秒
- `rc.SetHttpTransport` : 手动设置 http client
- `rc.GetHttpTransport` : 获得当前全局 http client

```go
package main

import "fmt"
import "time"
import "net"
import "net/http"
import "github.com/rongcloud/server-sdk-go/v4/rongcloud"

func main() {
	// 方法1： 创建对象时设置
	rc := sdk.NewRongCloud("appKey",
		"appSecret",
		// 每个域名最大活跃连接数
		sdk.WithMaxIdleConnsPerHost(100),
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
* [用户信息](./examples/user)
* [聊天室](./examples/chatroom)
* [群组](./examples/group)
* [超级群](./examples/ultragroup)
* [会话](./examples/conversation)
* [推送](./examples/push)
* [敏感词](./examples/sensitive)