package zoom

import (
	db "cloudSharePlatform/db"
	// utils "cloudSharePlatform/utils"

	// json "encoding/json"
	// "fmt"
	// "mime/multipart"
	"strconv"
	"strings"

	// "time"

	// "strconv"

	gin "github.com/gin-gonic/gin"
	// bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   header传token
   params:
        box_name: string filebox名称
        is_tmp: int 是否是临时文件 0:否 1:是
        is_pub: int 是否是公开 0:否 1:是
        files: string 文件id字符串用逗号隔开
*/
func FileBoxCreate(c *gin.Context) {

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
	params_boxName := c.PostForm("box_name")
	if params_boxName == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: box_name", "data": ""})
		return
	}
	param_isTmp := c.PostForm("is_tmp")
	if param_isTmp == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: is_tmp", "data": ""})
		return
	}
	param_isTmp_int, err := strconv.Atoi(param_isTmp)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "params is_tmp wrong format", "data": ""})
		return
	}
	// isTmp参数必须
	param_isPub := c.PostForm("is_pub")
	if param_isPub == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: is_pub", "data": ""})
		return
	}
	param_isPub_int, err := strconv.Atoi(param_isPub)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "params is_pub wrong format", "data": ""})
		return
	}
	// 文件id字符串
	param_files := c.PostForm("files")
	if param_files == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: files", "data": ""})
		return
	}
	// 文件id数组
	fileId_array := strings.Split(param_files, ",")

	// -------------------------------------------------------------------------
	// 将数据存到filebox表
	MongoFileBox := db.MongoFileBox{
		Id_:     primitive.NewObjectID(),
		User_id: user_id.(primitive.ObjectID),
		IsTmp:   param_isTmp_int,
		IsPub:   param_isPub_int,
		BoxName: params_boxName,
		Files:   fileId_array,
	}

	dbResult, err := db.MongoInsertOne("cloudshareplatform", "filebox", MongoFileBox)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "db insert to filbox err", "data": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": dbResult})
}
