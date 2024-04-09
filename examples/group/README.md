## 群组管理

### 方法列表
| 方法名                           | 说明          | 调用示例                                 |
|-------------------------------|-------------|--------------------------------------|
| GroupCreate                   | 创建群组        | [group.go](./group.go)               |
| GroupJoin                     | 加入群组        | [group.go](./group.go)               |
| GroupUserQuery                | 查询群组成员      | [group.go](./group.go)               |
| GroupDismiss                  | 解散群组        | [group.go](./group.go)               |
| GroupBanAdd                   | 设置群组全体禁言    | [groupban.go](./groupban.go)         |
| GroupBanQuery                 | 查询群组全体禁言    | [groupban.go](./groupban.go)         |
| GroupUserBanWhitelistAdd      | 加入群组全体禁言白名单 | [groupban.go](./groupban.go)         |
| GroupUserBanWhitelistQuery    | 查询群组全体禁言白名单 | [groupban.go](./groupban.go)         |
| GroupUserBanWhitelistRollback | 移除群组全体禁言白名单 | [groupban.go](./groupban.go)         |
| GroupRemarksSet               | 设置群成员推送备注名  | [groupremarks.go](./groupremarks.go) |
| GroupRemarksGet               | 查询群成员推送备注名  | [groupremarks.go](./groupremarks.go) |
| GroupRemarksDel               | 删除群成员推送备注名  | [groupremarks.go](./groupremarks.go) |
| GroupUserGagAdd               | 禁言指定群成员     | [groupusergag.go](./groupusergag.go) |
| GroupUserGagList              | 查询群成员禁言列表   | [groupusergag.go](./groupusergag.go) |
| GroupUserGagRollback          | 取消指定群成员禁言   | [groupusergag.go](./groupusergag.go) |
| UserGroupQuery                | 查询用户所在群组    | [usergroup.go](./usergroup.go)       |
| GroupSync                     | 同步用户所在群组    | [usergroup.go](./usergroup.go)       |
| GroupRefresh                  | 刷新群组信息      | [usergroup.go](./usergroup.go)       |
