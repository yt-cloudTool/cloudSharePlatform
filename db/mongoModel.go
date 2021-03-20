package db

import (
    "fmt"
    "context"
    "log"
    
    _ "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
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

// =============================================================================
//    操作方法
// =============================================================================
// insertone
func MongoInsertOne (dbName string, collName string, doc interface{}) (*mongo.InsertOneResult, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    result, err := collection.InsertOne(context.TODO(), doc)
    if err != nil {
    	return nil, err
    }
    defer func() {
        if err = mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("InsertOne a single document: ", result.InsertedID)
    
    return result, nil
}

// insertmany
func MongoInsertMany (dbName string, collName string, doc []interface{}) (*mongo.InsertManyResult, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    result, err := collection.InsertMany(context.TODO(), doc)
    if err != nil {
    	return nil, err
    }
    defer func() {
        if err = mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("InsertMany document: ", result.InsertedIDs)
    
    return result, nil
}

// updateone
func MongoUpdateOne (dbName string, collName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    result, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
    	return nil, err
    }
    defer func() {
        if err = mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("UpdateOne document: ", result.UpsertedID)
    
    return result, nil
}

// updatemany
func MongoUpdateMany (dbName string, collName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    result, err := collection.UpdateMany(context.TODO(), filter, update)
    if err != nil {
    	return nil, err
    }
    defer func() {
        if err = mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("UpdateMany document: ", result.UpsertedID)
    
    return result, nil
}

// findone
func MongoFindOne (dbName string, collName string, filter interface{}) (*mongo.SingleResult, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    result := collection.FindOne(context.TODO(), filter)
    defer func() {
        if err := mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("FindOne document: ", result)
    
    return result, nil
}

// findone
func MongoFind (dbName string, collName string, filter interface{}) (*mongo.Cursor, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    result, err := collection.Find(context.TODO(), filter)
    if err != nil {
    	return nil, err
    }
    defer func() {
        if err := mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("Find document: ", result)
    
    return result, nil
}