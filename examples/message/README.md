## 消息发送

### 方法列表
| 方法名                           | 说明         | 调用示例                                                               |
|-------------------------------|------------|--------------------------------------------------------------------|
| MessagePrivatePublish         | 发送单聊普通消息   | [messageprivate.go](./messageprivate.go)                           |
| MessageBroadcast              | 发送全量用户落地通知 | [messagebroadcast.go](./messagebroadcast.go)                       |
| MessageChatroomPublish        | 发送聊天室消息    | [messagechatroompublish.go](./messagechatroompublish.go)           |
| MessageExpansionSet           | 设置单群聊消息扩展  | [messageexpansion.go](./messageexpansion.go)                       |
| MessageExpansionQuery         | 获取单群聊消息扩展  | [messageexpansion.go](./messageexpansion.go)                       |
| MessageExpansionDelete        | 删除单群聊消息扩展  | [messageexpansion.go](./messageexpansion.go)                       |
| MessageGroupPublish           | 发送群消息      | [messagegrouppublish.go](./messagegrouppublish.go)                 |
| MessageHistory                | 获取历史消息日志   | [messagehistory.go](./messagehistory.go)                           |
| MessageHistoryDelete          | 删除历史消息日志   | [messagehistory.go](./messagehistory.go)                           |
| MessagePrivatePublishTemplate | 发送单聊模板消息   | [messageprivatetemplate.go](./messageprivatetemplate.go)           |
| MessageRecall                 | 撤回消息       | [messagerecall.go](./messagerecall.go)                             |
| MessageSystemPublish          | 发送系统通知普通消息 | [messagesystempublish.go](./messagesystempublish.go)               |
| MessageUltraGroupPublish      | 发送超级群消息    | [messageultragroup.go](./messageultragroup.go)                     |
| StatusMessagePrivatePublish   | 发送单聊状态消息   | [statusmessageprivatepublish.go](./statusmessageprivatepublish.go) |