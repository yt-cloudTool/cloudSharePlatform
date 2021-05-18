package zoom

import (
	db "cloudSharePlatform/db"
	// utils "cloudSharePlatform/utils"

	// json "encoding/json"
	// "fmt"
	// "mime/multipart"
	// "strconv"
	// "strings"

	// "time"

	// "strconv"
	gin "github.com/gin-gonic/gin"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	// mongo "go.mongodb.org/mongo-driver/mongo"
)

/*
   header传token
   params:
        box_id: string filebox id
*/
func FileBoxQuery(c *gin.Context) {

	user_id, isExist := c.Get("user_id")

	/*
	   如果存在user_id则检索此用户的文章
	   如果不存在user_id则返回空
	*/
	if isExist == false {
		c.JSON(200, gin.H{"status": 401, "message": "no user_id", "data": ""})
		return
	}

	// -------------------------------------------------------------------------
	// 参数
	params_boxId, _ := c.GetQuery("box_id")
	if params_boxId == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: box_id", "data": ""})
		return
	}
	// 生成OBJECT_ID
	params_boxObjId, err := primitive.ObjectIDFromHex(params_boxId)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "params not format err: box_id", "data": ""})
		return
	}
	// -------------------------------------------------------------------------
	// 成功id数组
	dbResult, err := db.MongoFindAll("cloudshareplatform", "filebox", bson.M{
		"$and": []bson.M{
			bson.M{"_id": params_boxObjId},
			bson.M{"user_id": user_id.(primitive.ObjectID)},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "db find from filbox err", "data": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": dbResult})
}
