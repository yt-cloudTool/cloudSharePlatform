package db

// 用户表
type MongoUser struct {
    LoginName string `bson:"loginname"` // 登录名(邮箱 手机号 自定义登录名等 默认同昵称 全局唯一)
    Nickname string `bson:"nickname"`   // 昵称 (仅显示用)
    Password string `bson:"password"`   // 加密后的密码
    Access   int `bson:"access"`        // 权限 0超级用户 1普通管理员 2普通用户 3冻结
}

// 用户配置信息表 (待定)
type MongoUserCfg struct {
    Id string `bson"_id"`             // 用户_id
    Functions interface{} `bson:"functions"` // 用户所能使用的菜单中的功能
}