package zoom

import (
	config "cloudSharePlatform/config"
	db "cloudSharePlatform/db"
	utils "cloudSharePlatform/utils"

	// "fmt"
	// "io"

	// "strconv"

	gin "github.com/gin-gonic/gin"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   params:
       fileid: string 文件id
*/
func FileDownload(c *gin.Context) {
	// -------------------------------------------------------------------------
	// 1
	// 获取文件id
	fileid, _ := c.GetQuery("fileid")
	if fileid == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough", "data": ""})
		return
	}
	// -------------------------------------------------------------------------
	// 2
	// 判断文件是否是公开的 如果不是公开的则校验token
	var objectFileId, _ = primitive.ObjectIDFromHex(fileid)
	dbResult, err := db.MongoFindOne("cloudshareplatform", "file", bson.M{"_id": objectFileId})
	if err != nil {
		c.JSON(500, gin.H{"status": -1, "message": "fileid not found err", "data": ""})
		return
	}
	// 获取数据库查询结果
	var dbResultData db.MongoFile
	err = dbResult.Decode(&dbResultData)
	if err != nil {
		c.JSON(500, gin.H{"status": -2, "message": err.Error(), "data": ""})
		return
	}
	// 判断是否是公开文件
	// 公开的文件
	if dbResultData.IsPub == 1 {
		// 执行下载
		FileDownloadControl(c, dbResultData)
	} else {
		// 非公开
		// -------------------------------------------------------------------------
		// 3
		// 解析token
		// 获取请求头token
		headerToken := c.GetHeader("token")
		user_id, err := utils.JwtValidate(headerToken)
		if err != nil {
			c.JSON(401, gin.H{"status": 0, "message": "access err1", "data": ""})
			return
		}
		// 如果不是此用户的文件
		if dbResultData.User_id.Hex() != user_id.Hex() {
			c.JSON(401, gin.H{"status": 0, "message": "access err2", "data": ""})
			return
		} else {
			// 如果是此用户的文件则执行下载
			FileDownloadControl(c, dbResultData)
		}
	}
}

// 文件下载方法
func FileDownloadControl(c *gin.Context, dbResultData db.MongoFile) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+dbResultData.FileName)
	c.Header("Content-Transfer-Encoding", "binary")

	c.File(config.GetFileStorePath() + dbResultData.User_id.Hex() + "/" + dbResultData.StoreFileName)
	// c.JSON(200, gin.H{"status": 1, "message": "ok", "data": ""})

}
