package db

import (
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   无逻辑删除
*/

// 用户表
type MongoUser struct {
	Id_       primitive.ObjectID `bson:"_id"`
	LoginName string             `bson:"loginname"` // 登录名(邮箱 手机号 自定义登录名等 默认同昵称 全局唯一)
	Nickname  string             `bson:"nickname"`  // 昵称 (仅显示用)
	Password  string             `bson:"password"`  // 加密后的密码
	Access    int                `bson:"access"`    // 权限 0超级用户 1普通管理员 2普通用户 3冻结
	Status    int                `bson:"status"`    // 用户状态 0正常 1冻结
}

// 用户配置信息表 (待定)
type MongoUserCfg struct {
	Id_       primitive.ObjectID `bson"_id"`
	User_id   primitive.ObjectID `bson"user_id"`    // 用户id
	Functions interface{}        `bson:"functions"` // 用户所能使用的菜单中的功能
}

// 可用功能表 (侧边栏菜单功能)
type MongoMenuAvailable struct {
	Id_          primitive.ObjectID `bson:"_id"`
	Code         int                `bson:"code"`         // 代号
	Permissions1 int                `bson:"permissions1"` // 所需权限 (用户类型)
	Permissions2 int                `bson:"permissions2"` // 所需二级权限 (用户状态,如有token则判断此项)
}

// 文章表 (桌面图标)
/*
   type: normal article note important filebox
*/
type MongoArticle struct {
	Id_         primitive.ObjectID `bson:"_id"`
	User_id     primitive.ObjectID `bson:"user_id"`     // 用户id
	Type        string             `bson:"type"`        // 类型 (判断图标类型)
	Img         string             `bson:"img"`         // 图标 (如果type==-1则用此字段显示图标)
	Label       string             `bson:"label"`       // 标签名称 (可重复)
	Content     string             `bson:"content"`     // 内容 (可重复)
	Fileboxid   string             `bson:"fileboxid"`   // 字符串数组
	Create_time int64              `bson:"create_time"` // 创建时间
}

// 文件表
type MongoFile struct {
	Id_           primitive.ObjectID `bson:"_id"`
	User_id       primitive.ObjectID `bson:"user_id"`       // 用户iduser
	IsTmp         int                `bson:"is_tmp"`        // 是否是临时文件
	IsPub         int                `bson:"is_pub"`        // 是否是公开文件
	FileName      string             `bson:"filename"`      // 文件名
	StoreFileName string             `bson:"storefilename"` // 存储的文件名
	Size          int64              `bson:"size"`          // 文件大小
}

// 文件box表
type MongoFileBox struct {
	Id_     primitive.ObjectID `bson:"_id"`
	User_id primitive.ObjectID `bson:"user_id"`  // 用户iduser
	IsTmp   int                `bson:"is_tmp"`   // 是否是临时文件box
	IsPub   int                `bson:"is_pub"`   // 是否是公开文件box
	BoxName string             `bson:"box_name"` // 文件box名
	Files   []string           `bson:"files`     // 文件id数组
}
