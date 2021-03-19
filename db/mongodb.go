package db

import (
	"context"
	"fmt"
	"log"

	_ "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoURI string = "mongodb://localhost:27017"

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

// =============================================================================
//    操作方法
// =============================================================================
// insertone
func MongoInsertOne (dbName string, collName string, doc interface{}) (*mongo.InsertOneResult, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    insertResult, err := collection.InsertOne(context.TODO(), doc)
    if err != nil {
    	return nil, err
    }
    defer func() {
        if err = mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
    
    return insertResult, nil
}

// insertmany
func MongoInsertMany (dbName string, collName string, doc []interface{}) (*mongo.InsertManyResult, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    insertResult, err := collection.InsertMany(context.TODO(), doc)
    if err != nil {
    	return nil, err
    }
    defer func() {
        if err = mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("Inserted a single document: ", insertResult.InsertedIDs)
    
    return insertResult, nil
}

// updateone
func MongoUpdateOne (dbName string, collName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
    var mongoCli = MongodbInit()
    collection := mongoCli.Database(dbName).Collection(collName)
    updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
    	return nil, err
    }
    defer func() {
        if err = mongoCli.Disconnect(context.TODO()); err != nil {
            log.Fatalln(err)
        }
    }()  
    
    fmt.Println("Inserted a single document: ", updateResult.InsertedIDs)
    
    return updateResult, nil
}