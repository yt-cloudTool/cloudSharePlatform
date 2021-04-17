package db

import (
	"context"
	"fmt"
	"log"

	_ "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// =============================================================================
// 模型
// =============================================================================
// 创建唯一键索引
func MongoCollectionUniqueIndexModel(dbName string, colName string, field string) {
	indexOption := options.Index()
	indexOption.SetBackground(true)
	indexOption.SetUnique(true)

	mod := mongo.IndexModel{
		Keys:    bsonx.Doc{{field, bsonx.Int32(1)}},
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
func MongoInsertOne(dbName string, collName string, doc interface{}) (*mongo.InsertOneResult, error) {
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
func MongoInsertMany(dbName string, collName string, doc []interface{}) (*mongo.InsertManyResult, error) {
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
func MongoUpdateOne(dbName string, collName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
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
func MongoUpdateMany(dbName string, collName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
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
func MongoFindOne(dbName string, collName string, filter interface{}) (*mongo.SingleResult, error) {
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

// find (dbname collname filter page size) 如果page和size都0则表示不进行分页搜索
func MongoFind(dbName string, collName string, filter interface{}, page int64, size int64) ([]map[string]interface{}, error) {
	var mongoCli = MongodbInit()

	// options
	findOptions := new(options.FindOptions)
	if page != 0 && size != 0 {
		var limit int64 = size
		var skip int64 = (page - 1) * size

		if page > 0 && size > 0 {
			findOptions.SetSkip(skip)
			findOptions.SetLimit(limit)
		}
	}

	collection := mongoCli.Database(dbName).Collection(collName)
	result, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := mongoCli.Disconnect(context.TODO()); err != nil {
			log.Fatalln(err)
		}
	}()

	fmt.Println("Find document: ", result)

	var resultData []map[string]interface{}
	err = result.All(context.TODO(), &resultData)

	return resultData, nil
}

// =============================================================================
//    事务
// =============================================================================
// transaction
type MongoTransacStruct struct {
	DbName  string
	ColName string

	ctx        context.Context
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	session    mongo.Session
	errSlice   []error

	// 记录这次操作的结果(按类型分为五种)
	InsertOneResult  *mongo.InsertOneResult
	InsertManyResult *mongo.InsertManyResult
	UpdateResult     *mongo.UpdateResult
	FindOneResult    *mongo.SingleResult
	FindResult       *mongo.Cursor

	// 操作结果历史
	ResultRecord []interface{}
}

func (self *MongoTransacStruct) MongoTransaction() *MongoTransacStruct {
	var err error
	// cient
	self.client = MongodbInit()
	// database
	self.database = self.client.Database(self.DbName)
	// ctx
	self.ctx = context.Background()
	// collection
	self.collection = self.database.Collection(self.ColName)
	// session
	self.session, err = self.database.Client().StartSession()
	if err != nil {
		self.errSlice = append(self.errSlice, err)
		return self
	}
	// 事务开始
	err = self.session.StartTransaction()
	if err != nil {
		self.errSlice = append(self.errSlice, err)
		return self
	}
	return self
}

func (self *MongoTransacStruct) MongoTransactionAbort() *MongoTransacStruct {
	self.session.AbortTransaction(self.ctx)
	return self
}

func (self *MongoTransacStruct) MongoTransactionCommit() *MongoTransacStruct {
	self.session.CommitTransaction(self.ctx)
	return self
}

func (self *MongoTransacStruct) MongoTransactionEnd() *MongoTransacStruct {
	self.session.EndSession(self.ctx)
	self.database.Client().Disconnect(self.ctx)
	return self
}

func (self *MongoTransacStruct) MongoTransactionInsertOne(doc interface{}, opts ...*options.InsertOneOptions) *MongoTransacStruct {
	result, err := self.collection.InsertOne(self.ctx, doc, opts...)
	if err != nil {
		self.errSlice = append(self.errSlice, err)
		return self
	}
	self.InsertOneResult = result
	self.ResultRecord = append(self.ResultRecord, result)
	return self
}

func (self *MongoTransacStruct) MongoTransactionInsertMoney(doc []interface{}, opts ...*options.InsertManyOptions) *MongoTransacStruct {
	result, err := self.collection.InsertMany(self.ctx, doc, opts...)
	if err != nil {
		self.errSlice = append(self.errSlice, err)
		return self
	}
	self.InsertManyResult = result
	self.ResultRecord = append(self.ResultRecord, result)
	return self
}

func (self *MongoTransacStruct) MongoTransactionUpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) *MongoTransacStruct {
	result, err := self.collection.UpdateOne(self.ctx, filter, update, opts...)
	if err != nil {
		self.errSlice = append(self.errSlice, err)
		return self
	}
	self.UpdateResult = result
	self.ResultRecord = append(self.ResultRecord, result)
	return self
}

func (self *MongoTransacStruct) MongoTransactionUpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) *MongoTransacStruct {
	result, err := self.collection.UpdateMany(self.ctx, filter, update, opts...)
	if err != nil {
		self.errSlice = append(self.errSlice, err)
		return self
	}
	self.UpdateResult = result
	self.ResultRecord = append(self.ResultRecord, result)
	return self
}

func (self *MongoTransacStruct) MongoTransactionFindOne(filter interface{}, opts ...*options.FindOneOptions) *MongoTransacStruct {
	result := self.collection.FindOne(self.ctx, filter, opts...)
	self.FindOneResult = result
	self.ResultRecord = append(self.ResultRecord, result)
	return self
}

func (self *MongoTransacStruct) MongoTransactionFind(filter interface{}, page int64, size int64) *MongoTransacStruct {
	// options
	findOptions := new(options.FindOptions)
	if page != 0 && size != 0 {
		var limit int64 = size
		var skip int64 = (page - 1) * size

		if page > 0 && size > 0 {
			findOptions.SetSkip(skip)
			findOptions.SetLimit(limit)
		}
	}
	result, err := self.collection.Find(self.ctx, filter, findOptions)
	if err != nil {
		self.errSlice = append(self.errSlice, err)
		return self
	}
	self.FindResult = result
	self.ResultRecord = append(self.ResultRecord, result)
	return self
}
