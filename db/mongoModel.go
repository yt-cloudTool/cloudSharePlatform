package db

import (
    "fmt"
    "context"
    "log"
)

// 测试用
func MongoTst () {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database("cloudshareplatform").Collection("user")
    insertResult, err := collection.InsertOne(context.TODO(), MongoUser{
        
    })
    if err != nil {
    	log.Fatal(err)
    }
    defer func() {
        if err = mongoCli.Disconnect(context.TODO()); err != nil { log.Fatal(err) }
    }()    
    
    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

