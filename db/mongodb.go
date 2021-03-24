package db

import (
	"context"
	"fmt"
	"log"

	bson "go.mongodb.org/mongo-driver/bson"
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
