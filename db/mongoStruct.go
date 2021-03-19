package db

// 用户表结构
type MongoUser struct {
	Nickname string `bson:"nickname"`        // 昵称
    Password string `bson:"password"`        // 加密后的密码
    Access   string `bson:"access"`          // 权限 0超级用户 1普通管理员 2普通用户 3冻结
}