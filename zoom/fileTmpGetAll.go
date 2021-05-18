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
func FileTmpGetAll(c *gin.Context) {

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
	// 成功id数组
	dbResult, err := db.MongoFindAll("cloudshareplatform", "file", bson.M{
		"$and": []bson.M{
			bson.M{"is_tmp": 1},
			bson.M{"user_id": user_id.(primitive.ObjectID)},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "db find from file err", "data": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": dbResult})
}
