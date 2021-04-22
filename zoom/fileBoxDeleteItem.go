package zoom

import (
	db "cloudSharePlatform/db"
	// utils "cloudSharePlatform/utils"

	// json "encoding/json"
	// "fmt"
	// "mime/multipart"
	// "strconv"
	"strings"

	// "time"

	// "strconv"

	gin "github.com/gin-gonic/gin"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   header传token
   params:
        box_id: string filebox id
        files: []string 文件id字符串用逗号隔开
*/
func FileBoxDeleteItem(c *gin.Context) {

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
	params_boxId := c.PostForm("box_id")
	if params_boxId == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: box_id", "data": ""})
		return
	}

	params_files := c.PostForm("files")
	if params_files == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: files", "data": ""})
		return
	}

	paramsBoxId_Object, err := primitive.ObjectIDFromHex(params_boxId)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "params format err: box_id", "data": ""})
		return
	}

	// 文件id数组
	fileId_array := strings.Split(params_files, ",")

	// --------------------------------------------------------------------------
	dbResult, err := db.MongoUpdateOne("cloudshareplatform", "filebox", bson.M{
		"$and": bson.A{
			bson.M{"user_id": user_id.(primitive.ObjectID)},
			bson.M{"_id": paramsBoxId_Object},
		},
	}, bson.M{
		"$pull": bson.M{
			"files": bson.M{"$in": fileId_array},
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "db insert to filbox err", "data": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": dbResult})
}
