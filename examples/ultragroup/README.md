## 超级群管理

### 方法列表
| 方法名                              | 说明            | 调用示例                                                                   |
|----------------------------------|---------------|------------------------------------------------------------------------|
| UltraGroupCreate                 | 创建超级群         | [ultragroup.go](./ultragroup.go)                                       |
| UltraGroupJoin                   | 加入超级群         | [ultragroup.go](./ultragroup.go)                                       |
| UltraGroupMemberExist            | 查询用户是否为群成员    | [ultragroup.go](./ultragroup.go)                                       |
| UltraGroupQuit                   | 退出超级群         | [ultragroup.go](./ultragroup.go)                                       |
| UltraGroupRefresh                | 刷新超级群信息       | [ultragroup.go](./ultragroup.go)                                       |
| UltraGroupDis                    | 解散超级群         | [ultragroup.go](./ultragroup.go)                                       |
| UltraGroupBannedWhitelistAdd     | 加入超级群全体禁言白名单  | [ultragroupbannedwhitelist.go](./ultragroupbannedwhitelist.go)         |
| UltraGroupBannedWhitelistGet     | 查询群组全体禁言白名单   | [ultragroupbannedwhitelist.go](./ultragroupbannedwhitelist.go)         |
| UltraGroupBannedWhitelistDel     | 移出超级群全体禁言白名单  | [ultragroupbannedwhitelist.go](./ultragroupbannedwhitelist.go)         |
| UltraGroupChannelCreate          | 创建频道          | [ultragroupchannel.go](./ultragroupchannel.go)                         |
| UltraGroupChannelGet             | 查询频道列表        | [ultragroupchannel.go](./ultragroupchannel.go)                         |
| UltraGroupChannelTypeChange      | 变更频道类型        | [ultragroupchannel.go](./ultragroupchannel.go)                         |
| UltraGroupChannelDel             | 删除频道          | [ultragroupchannel.go](./ultragroupchannel.go)                         |
| UltraGroupChannelPrivateUsersAdd | 添加私有频道成员      | [ultragroupchannelprivateusers.go](./ultragroupchannelprivateusers.go) |
| UltraGroupChannelPrivateUsersGet | 查询私有频道成员列表    | [ultragroupchannelprivateusers.go](./ultragroupchannelprivateusers.go) |
| UltraGroupChannelPrivateUsersDel | 删除私有频道成员      | [ultragroupchannelprivateusers.go](./ultragroupchannelprivateusers.go) |
| UltraGroupGlobalBannedSet        | 设置超级群全体禁言     | [ultragroupglobalban.go](./ultragroupglobalban.go)                     |
| UltraGroupGlobalBannedGet        | 查询超级群全体禁言     | [ultragroupglobalban.go](./ultragroupglobalban.go)                     |
| UltraGroupHismsgQuery            | 搜索超级群消息       | [ultragrouphismsg.go](./ultragrouphismsg.go)                           |
| UltraGroupHismsgMsgIdQuery       | 搜索超级群消息上下文    | [ultragrouphismsg.go](./ultragrouphismsg.go)                           |
| UltraGroupMessageExpansionSet    | 设置超级群消息扩展     | [ultragroupmessageexansion.go](./ultragroupmessageexansion.go)         |
| UltraGroupMessageExpansionQuery  | 获取超级群消息扩展     | [ultragroupmessageexansion.go](./ultragroupmessageexansion.go)         |
| UltraGroupMessageExpansionDelete | 删除超级群消息扩展     | [ultragroupmessageexansion.go](./ultragroupmessageexansion.go)         |
| UltraGroupNotDisturbSet          | 设置群/频道默认免打扰   | [ultragroupnotdisturb.go](./ultragroupnotdisturb.go)                   |
| UltraGroupNotDisturbGet          | 查询默认免打扰配置     | [ultragroupnotdisturb.go](./ultragroupnotdisturb.go)                   |
| UltraGroupUserBannedAdd          | 禁言指定超级群成员     | [ultragroupuserbanned.go](./ultragroupuserbanned.go)                   |
| UltraGroupUserBannedGet          | 查询超级群成员禁言列表   | [ultragroupuserbanned.go](./ultragroupuserbanned.go)                   |
| UltraGroupUserBannedDel          | 取消指定超级群成员禁言   | [ultragroupuserbanned.go](./ultragroupuserbanned.go)                   |
| UltraGroupUserGroupAdd           | 创建用户组         | [ultragroupusergroup.go](./ultragroupusergroup.go)                     |
| UltraGroupUserGroupDel           | 删除用户组         | [ultragroupusergroup.go](./ultragroupusergroup.go)                     |
| UltraGroupUserGroupQuery         | 查询用户组列表       | [ultragroupusergroup.go](./ultragroupusergroup.go)                     |
| UltraGroupChannelUserGroupBind   | 超级群绑定频道与用户组   | [ultragroupusergroup.go](./ultragroupusergroup.go)                     |
| UltraGroupChannelUserGroupQuery  | 超级群查询频道绑定的用户组 | [ultragroupusergroup.go](./ultragroupusergroup.go)                     |
| UltraGroupUserGroupChannelQuery  | 超级群查询用户组绑定的频道 | [ultragroupusergroup.go](./ultragroupusergroup.go)                     |
| UltraGroupChannelUserGroupUnbind | 超级群解绑频道与用户组   | [ultragroupusergroup.go](./ultragroupusergroup.go)                     |
| UltraGroupUserGroupUserAdd       | 超级群用户组添加用户    | [ultragroupusergroupuser.go](./ultragroupusergroupuser.go)             |
| UltraGroupUserUserGroupQuery     | 超级群查询用户所属用户组  | [ultragroupusergroupuser.go](./ultragroupusergroupuser.go)             |
| UltraGroupUserGroupUserDel       | 超级群用户组移出用户    | [ultragroupusergroupuser.go](./ultragroupusergroupuser.go)             |
