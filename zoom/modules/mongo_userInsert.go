package zoom

import (
	db "cloudSharePlatform/db"

	// gin "github.com/gin-gonic/gin"

	// bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	// options "go.mongodb.org/mongo-driver/mongo/options"
	// utils "cloudSharePlatform/utils"

	mongo "go.mongodb.org/mongo-driver/mongo"
)

func Mongo_userInsert(
	id_ primitive.ObjectID,
	loginname string,
	nickname string,
	password string,
	access int) (*mongo.InsertOneResult, error) {

	// 设置数据
	userDbStoreData := db.MongoUser{
		Id_:       primitive.NewObjectID(),
		LoginName: loginname,
		Nickname:  nickname,
		Password:  password,
		Access:    2, // 默认普通用户权限
	}

	// 添加用户
	dbResult, err := db.MongoInsertOne("cloudshareplatform", "user", userDbStoreData)
	if err != nil {
		return nil, err
	}

	return dbResult, nil
}
