package zoom

import (
	// db "cloudSharePlatform/db"

	gin "github.com/gin-gonic/gin"

	// bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	// options "go.mongodb.org/mongo-driver/mongo/options"
	utils "cloudSharePlatform/utils"
	modules "cloudSharePlatform/zoom/modules"
)

func UserRegister(c *gin.Context) {
	loginname := c.PostForm("loginname")
	password := c.PostForm("password")

	// 判断参数
	if loginname == "" || password == "" {
		c.JSON(401, gin.H{"status": 0, "message": "post data not enough", "data": ""})
		return
	}

	// 生成密码
	generatedPwd, err := utils.BcryptGenerate(password)
	if err != nil {
		c.JSON(401, gin.H{"status": -1, "message": "hash password generate err", "data": ""})
		return
	}

	// 添加用户
	dbResult, err := modules.Mongo_userInsert(
		primitive.NewObjectID(),
		loginname,
		loginname,
		generatedPwd,
		2)
	if err != nil {
		c.JSON(401, gin.H{"status": -2, "message": "login name err", "data": ""})
		return
	}
	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": dbResult})
	return
}
