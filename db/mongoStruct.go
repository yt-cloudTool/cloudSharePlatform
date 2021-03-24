package db

// 用户表
type MongoUser struct {
	LoginName string `bson:"loginname"` // 登录名(邮箱 手机号 自定义登录名等 默认同昵称 全局唯一)
	Nickname  string `bson:"nickname"`  // 昵称 (仅显示用)
	Password  string `bson:"password"`  // 加密后的密码
	Access    int    `bson:"access"`    // 权限 0超级用户 1普通管理员 2普通用户 3冻结
	Status    int    `bson:"status"`    // 用户状态 0正常 1冻结
}

// 用户配置信息表 (待定)
type MongoUserCfg struct {
	Id        string      `bson"_id"`        // 用户_id
	Functions interface{} `bson:"functions"` // 用户所能使用的菜单中的功能
}

// 可用功能表 (侧边栏菜单功能)
type MongoAvailMenuItem struct {
	Code         int `bson:"code"`         // 代号
	Permissions1 int `bson:"permissions1"` // 所需权限 (用户类型)
	Permissions2 int `bson:"permissions2"` // 所需二级权限 (用户状态)
}
