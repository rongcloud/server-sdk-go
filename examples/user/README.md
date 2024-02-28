## 用户管理

### 方法列表

| 方法名                       | 说明         | 调用示例                                               |
|---------------------------|------------|----------------------------------------------------|
| UserBlacklistAdd          | 添加用户到黑名单   | [userblacklist.go](./userblacklist.go)             |
| UserBlacklistQuery        | 获取黑名单用户列表  | [userblacklist.go](./userblacklist.go)             |
| UserBlacklistRemove       | 移除黑名单中用户   | [userblacklist.go](./userblacklist.go)             |
| UserBlock                 | 封禁用户       | [userblock.go](./userblock.go)                     |
| UserBlockQuery            | 获取封禁用户列表   | [userblock.go](./userblock.go)                     |
| UserUnBlock               | 解除封禁       | [userblock.go](./userblock.go)                     |
| UserBlockPushPeriodSet    | 设置用户免打扰时段  | [userblockpushperiod.go](./userblockpushperiod.go) |
| UserBlockPushPeriodGet    | 查询用户免打扰时段  | [userblockpushperiod.go](./userblockpushperiod.go) |
| UserBlockPushPeriodDelete | 删除用户免打扰时段  | [userblockpushperiod.go](./userblockpushperiod.go) |
| UserChatFBSet             | 设置用户单聊禁言   | [userchatfb.go](./userchatfb.go)                   |
| UserChatFBQueryList       | 查询单聊禁言用户列表 | [userchatfb.go](./userchatfb.go)                   |
| UserCheckOnline           | 查询用户在线状态   | [usercheckonline.go](./usercheckonline.go)         |
| UserDeactivate            | 注销用户       | [userdeactivate.go](./userdeactivate.go)           |
| UserDeactivateQuery       | 查询已注销用户    | [userdeactivate.go](./userdeactivate.go)           |
| UserReactivate            | 重新激活用户 ID  | [userdeactivate.go](./userdeactivate.go)           |
| UserRefresh               | 修改用户信息     | [userrefresh.go](./userrefresh.go)                 |
| UserRemarksSet            | 设置用户级推送备注名 | [userremarks.go](./userremarks.go)                 |
| UserRemarksGet            | 查询用户级推送备注名 | [userremarks.go](./userremarks.go)                 |
| UserRemarksDel            | 删除用户级推送备注名 | [userremarks.go](./userremarks.go)                 |
| UserTagSet                | 设置用户标签     | [usertag.go](./usertag.go)                         |
| UserTagBatchSet           | 批量设置用户标签   | [usertag.go](./usertag.go)                         |
| UserTagsGet               | 获取用户标签     | [usertag.go](./usertag.go)                         |
| UserGetToken              | 注册用户       | [usertoken.go](./usertoken.go)                     |
| UserTokenExpire           | 作废token    | [usertoken.go](./usertoken.go)                     |
| UserWhitelistAdd          | 添加用户到白名单   | [userwhitelist.go](./userwhitelist.go)             |
| UserWhitelistQuery        | 移除白名单中用户   | [userwhitelist.go](./userwhitelist.go)             |
| UserWhitelistRemove       | 移除白名单中用户   | [userwhitelist.go](./userwhitelist.go)             |
