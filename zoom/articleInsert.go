package zoom

import (
	db "cloudSharePlatform/db"
	// utils "cloudSharePlatform/utils"
	// json "encoding/json"
	"fmt"
	"time"

	// "strconv"

	gin "github.com/gin-gonic/gin"
	// bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	// primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   header传token
    params:
        label string
        type string (type=="article"时内容为富文本)
        img string (如果type==-1则使用此项, 但是type字段也要存储)
        content string 内容
        fileboxid string  文件夹id (每个文章可以关联一个或多个文件夹) 格式为json字符串"[fileboxid1, fileboxid2...]"
*/
func ArticleInsert(c *gin.Context) {

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
	param_label := c.PostForm("label")
	// label参数必须
	if param_label == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough", "data": ""})
		return
	}
	param_type := c.PostForm("type")
	if param_type == "" {
		param_type = "normal" // 默认为 normal
	}
	param_img := c.PostForm("img")
	param_content := c.PostForm("content")
	param_fileboxid := c.PostForm("fileboxid") // json数组字符串
	// 存储转换为数组的fileboxid array
	// var param_fileboxidArr []string
	// if param_fileboxid != "" {
	// 	var tmpFileboxidArr []string
	// 	err := json.Unmarshal([]byte(param_fileboxid), &tmpFileboxidArr)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{"status": 0, "message": "fileboxid param err", "data": ""})
	// 		return
	// 	}
	// 	param_fileboxidArr = tmpFileboxidArr
	// }string
	// -------------------------------------------------------------------------

	// 插入
	mongoArticle := db.MongoArticle{
		Id_:         primitive.NewObjectID(),
		User_id:     user_id.(primitive.ObjectID),
		Type:        param_type,                  // 类型 (判断图标类型)
		Img:         param_img,                   // 图标 (如果type==-1则用此字段显示图标)
		Label:       param_label,                 // 标签名称 (可重复)
		Content:     param_content,               // 内容
		Fileboxid:   param_fileboxid,             // 字符串
		Create_time: time.Now().UnixNano() / 1e6, // 毫秒时间戳
	}

	dbResult, err := db.MongoInsertOne("cloudshareplatform", "article", mongoArticle)
	if err != nil {
		c.JSON(500, gin.H{"status": -1, "message": "db find err", "data": ""})
	}
	fmt.Println("dbResult ===========>", dbResult)

	// 操作articleList表
	c.JSON(200, gin.H{"status": 0, "message": "ok", "data": dbResult})
}
