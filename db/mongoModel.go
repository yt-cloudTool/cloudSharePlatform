package db

import (
    "fmt"
    "context"
    "log"
    
    _ "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/x/bsonx"
    "go.mongodb.org/mongo-driver/mongo/options"
)


// =============================================================================
// 模型
// =============================================================================
// 创建唯一键索引
func MongoCollectionUniqueIndexModel (dbName string, colName string, field string) {
    indexOption := options.Index()
    indexOption.SetBackground(true)
    indexOption.SetUnique(true)
    
    mod := mongo.IndexModel{
        Keys: bsonx.Doc{{field, bsonx.Int32(1)}},
        Options: indexOption,
    }
    
    // 创建索引
    collection := MongodbInit().Database(dbName).Collection(colName)
    index, err := collection.Indexes().CreateOne(context.TODO(), mod)
    if err != nil {
        log.Fatalln("create index err ===========>", err)
    }
    fmt.Println("index is ==========>", index)
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