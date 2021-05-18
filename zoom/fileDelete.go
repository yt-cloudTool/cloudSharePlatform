package zoom

import (
	db "cloudSharePlatform/db"

	gin "github.com/gin-gonic/gin"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   header传token
   params:
        file_id: string 文件id
*/
func FileDelete(c *gin.Context) {

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
	file_id := c.PostForm("file_id")
	if file_id == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: file_id", "data": ""})
		return
	}
	// 生成OBJECT_ID
	params_fileId, err := primitive.ObjectIDFromHex(file_id)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "params not format err: file_id", "data": ""})
		return
	}
	// -------------------------------------------------------------------------
	// 查询文件id对应的文件
	// dbResult, err := db.MongoFindAll("cloudshareplatform", "file", bson.M{
	// 	"$and": []bson.M{
	// 		bson.M{"_id": params_fileId},
	// 		bson.M{"user_id": user_id.(primitive.ObjectID)},
	// 	},
	// })
	// if err != nil {
	// 	c.JSON(500, gin.H{"status": 0, "message": "db find from filb err", "data": err.Error()})
	// 	return
	// }

	var mongoLink = db.MongoLink{DbName: "cloudshareplatform", CollName: "file"}
	var collection = mongoLink.GetCollection()
	var dbRes = collection.FindOne(mongoLink.Ctx, bson.M{
		"$and": []bson.M{
			bson.M{"_id", params_fileId},
			bson.M{"user_id": user_id.(primitive.ObjectID)},
		},
	})

	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": dbRes})
}
