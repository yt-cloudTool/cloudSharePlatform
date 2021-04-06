package zoom

import (
	db "cloudSharePlatform/db"
	// utils "cloudSharePlatform/utils"
	"fmt"

	"strconv"

	gin "github.com/gin-gonic/gin"
	bson "go.mongodb.org/mongo-driver/bson"
	// primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   header传token
    params:
        label string
        type string
        time string
        page int
        size int

*/
func ArticleListGet(c *gin.Context) {

	user_id, isExist := c.Get("user_id")
	/*
	   如果存在user_id则检索此用户的文章
	   如果不存在user_id则返回空
	*/
	if isExist == false {
		c.JSON(200, gin.H{"status": -100, "message": "no user_id", "data": ""})
		return
	}

	// -------------------------------------------------------------------------
	// 参数
	param_page, _ := c.GetQuery("page") // 搜索条件 当前页码
	param_size, _ := c.GetQuery("size") // 搜索条件 每页显示数量
	// 两种参数必须
	if param_page == "" || param_size == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough", "data": ""})
		return
	}
	// 页码与数量字段转int
	param_page_int, err := strconv.ParseInt(param_page, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "page type err", "data": ""})
		return
	}
	param_size_int, err := strconv.ParseInt(param_size, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "size type err", "data": ""})
		return
	}
	param_label, _ := c.GetQuery("label") // 搜索条件 文章名称(可选)
	param_type, _ := c.GetQuery("type")   // 搜索条件 文章类型(可选)
	param_time, _ := c.GetQuery("time")   // 搜索条件 时间排序(可选)
	// -------------------------------------------------------------------------

	fmt.Println("params ===============>", user_id, param_page, param_size, param_page_int, param_size_int, param_label, param_type, param_time)

	// 查询
	dbResult, err := db.MongoFind("cloudshareplatform", "article", bson.M{
		"user_id": user_id,
		"$and": bson.A{
			// bson.M{
			// 	"$or": bson.A{
			// 		bson.M{"type": param_type},
			// 	},
			// },
			// bson.M{
			// 	"$or": bson.A{
			bson.M{"label": bson.M{"$regex": param_label}},
			// },
			// },
		},
	}, param_page_int, param_size_int)
	if err != nil {
		c.JSON(500, gin.H{"status": -1, "message": "db find err", "data": ""})
		return
	}

	// var dbResultData db.MongoArticle
	// err = dbResult.Decode(&dbResultData)
	// if err != nil {
	// 	c.JSON(500, gin.H{"status": -2, "message": "dbResult Decode err", "data": ""})
	// 	return
	// }
	// fmt.Println("dbResult ===========>", dbResult, param_time)

	// dbResult.All()

	// 操作articleList表
	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": dbResult})
}
