package zoom

import (
	db "cloudSharePlatform/db"
	utils "cloudSharePlatform/utils"
	"fmt"

	gin "github.com/gin-gonic/gin"
	bson "go.mongodb.org/mongo-driver/bson"
)

func UserLogin(c *gin.Context) {
	fmt.Println("params ==================>", c.Request)
	loginname := c.PostForm("loginname")
	password := c.PostForm("password")

	if loginname == "" || password == "" {
		c.JSON(401, gin.H{"status": 0, "message": "params not enough", "data": ""})
		return
	}

	// 用获取用户密码
	dbResult, err := db.MongoFindOne("cloudshareplatform", "user", bson.M{"loginname": loginname})
	if err != nil {
		c.JSON(401, gin.H{"status": -1, "message": "login name err", "data": ""})
		return
	}

	// 获取数据库查询结果
	var dbResultData db.MongoUser
	err = dbResult.Decode(&dbResultData)
	if err != nil {
		c.JSON(401, gin.H{"status": -2, "message": "db decode err", "data": ""})
		return
	}

	// 判断密码正确性
	compareResult := utils.BcryptCompare(dbResultData.Password, password)
	if compareResult == true {
		// 生成token
		tokenStr, err := utils.JwtCreate(loginname)
		if err != nil {
			c.JSON(401, gin.H{"status": -4, "message": "create token err", "data": ""})
			return
		}
		c.JSON(200, gin.H{
			"status":  1,
			"message": "ok",
			"data": gin.H{
				"token":     tokenStr,
				"loginname": dbResultData.LoginName,
				"nickname":  dbResultData.Nickname,
				"access":    dbResultData.Access,
			},
		})
		return
	} else {
		c.JSON(401, gin.H{"status": -3, "message": "password compare err", "data": ""})
		return
	}
}
