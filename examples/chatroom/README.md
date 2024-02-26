## 聊天室

### 方法列表
| 方法名                              | 说明            | 调用示例                                           |
|----------------------------------|---------------|------------------------------------------------|
| ChatroomBanAdd                   | 设置聊天室全体禁言     | [chatroomban.go](./chatroomban.go)             |
| ChatroomBanQuery                 | 查询聊天室全体禁言列表   | [chatroomban.go](./chatroomban.go)             |
| ChatroomBanRollback              | 取消聊天室全体禁言     | [chatroomban.go](./chatroomban.go)             |
| ChatroomUserBanWhitelistAdd      | 加入聊天室全体禁言白名单  | [chatroomban.go](./chatroomban.go)             |
| ChatroomUserBanWhitelistQuery    | 查询聊天室全体禁言白名单  | [chatroomban.go](./chatroomban.go)             |
| ChatroomUserBanWhitelistRollback | 移出聊天室全体禁言白名单  | [chatroomban.go](./chatroomban.go)             |
| ChatroomEntrySet                 | 设置聊天室属性（KV）   | [chatroomentry.go.go](./chatroomentry.go.go)   |
| ChatroomEntryBatchSet            | 批量设置聊天室属性（KV） | [chatroomentry.go](./chatroomentry.go)         |
| ChatroomEntryRemove              | 删除聊天室属性（KV)   | [chatroomentry.go](./chatroomentry.go)         |
| ChatroomDestroySet               | 设置聊天室销毁类型     | [chatroommeta.go.go](./chatroommeta.go.go)     |
| ChatroomGet                      | 查询聊天室信息       | [chatroommeta.go](./chatroommeta.go)           |
| ChatroomDestroy                  | 销毁聊天室         | [chatroommeta.go](./chatroommeta.go)           |
| ChatroomUserExist                | 查询用户是否在聊天室中   | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUsersExist               | 批量查询用户是否在聊天室中 | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserQuery                | 获取聊天室成员       | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserBlockAdd             | 封禁聊天室用户       | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserBlockList            | 查询聊天室封禁用户     | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserBlockRollback        | 解除封禁聊天室用户     | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserBanAdd               | 全局禁言用户        | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserBanQuery             | 查询全局禁言用户列表    | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserBanRemove            | 取消全局禁言用户      | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserGagAdd               | 禁言指定聊天室用户     | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserGagList              | 查询聊天室禁言用户列表   | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomUserGagRollback          | 取消禁言指定聊天室用户   | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomMessagePriorityAdd       | 添加低级别消息类型     | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomMessagePriorityQuery     | 查询低级别消息类型     | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomMessagePriorityRemove    | 移除低级别消息类型     | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomKeepaliveAdd             | 保活聊天室         | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomKeepaliveQuery           | 查询保活聊天室       | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomKeepaliveRemove          | 取消保活聊天室       | [chatroomuser.go](./chatroomuser.go)           |
| ChatroomWhitelistAdd             | 加入聊天室消息白名单    | [chatroomwhitelist.go](./chatroomwhitelist.go) |
| ChatroomWhitelistQuery           | 查询聊天室消息白名单    | [chatroomwhitelist.go](./chatroomwhitelist.go) |
| ChatroomWhitelistRemove          | 移出聊天室消息白名单    | [chatroomwhitelist.go](./chatroomwhitelist.go) |
| ChatroomUserWhitelistAdd         | 加入聊天室用户白名单    | [chatroomwhitelist.go](./chatroomwhitelist.go) |
| ChatroomUserWhitelistQuery       | 查询聊天室用户白名单    | [chatroomwhitelist.go](./chatroomwhitelist.go) |
| ChatroomUserWhitelistRemove      | 移出聊天室用户白名单    | [chatroomwhitelist.go](./chatroomwhitelist.go) |
