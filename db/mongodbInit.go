package db

import (
	"context"
	"fmt"
	"log"

	bson "go.mongodb.org/mongo-driver/bson"
	// primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoURI string = "mongodb://localhost:27017"

type BsonM bson.M

func init() {
	// mongo文档模型初始化
	// ===========================================================================
	// 创建唯一键索引
	MongoCollectionUniqueIndexModel("cloudshareplatform", "user", "loginname")
	MongoCollectionUniqueIndexModel("cloudshareplatform", "menuavailable", "code")

	// /*
	//    mongo fileBox默认临时文档 用于存储临时文件
	//    默认id为-1
	// */
	// // ===========================================================================
	// // 将数据存到filebox表
	//    t_id, _ = primitive.ObjectIDFromHex("-1")
	// MongoFileBox := MongoFileBox{
	// 	Id_:     t_id,
	// 	User_id: user_id.(primitive.ObjectID),
	// 	IsTmp:   param_isTmp_int,
	// 	IsPub:   param_isPub_int,
	// 	BoxName: params_boxName,
	// 	Files:   fileId_array,
	// }

	// dbResult, err := db.MongoInsertOne("cloudshareplatform", "filebox", MongoFileBox)
	// if err != nil {
	// 	c.JSON(500, gin.H{"status": 0, "message": "db insert to filbox err", "data": err.Error()})
	// 	return
	// }
}

func MongodbInit() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
