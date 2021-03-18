package db

import (
    "context"
    "fmt"
    "log"

    // "go.mongodb.org/mongo-driver/bson"
    mongo "go.mongodb.org/mongo-driver/mongo"
    options "go.mongodb.org/mongo-driver/mongo/options"
)

var mongodbClient *mongo.Client

func GetMongoCli () *mongo.Client {
    return mongodbClient
}

func MongodbInit () {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    
    // Connect to MongoDB
    MongodbClient, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    
    // Check the connection
    err = MongodbClient.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Connected to MongoDB!")
    
    mongodbTst()
}

func mongodbTst () {
    collection := GetMongoCli().Database("test").Collection("test")
    
    type Student struct {
    	Id int
    	Name string
    }
    
    insertResult, err := collection.InsertOne(context.TODO(), Student{12, "小红"})
    if err != nil {
    	log.Fatal(err)
    }
    
    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}