package zoom

import (
	// "net/http"

	gin "github.com/gin-gonic/gin"
    db "cloudSharePlatform/db"
    bson "go.mongodb.org/mongo-driver/bson"
    utils "cloudSharePlatform/utils"
)

func UserLogin (c *gin.Context) {
    loginname := c.PostForm("loginname")
    password  := c.PostForm("password")
    
    // 用获取用户密码
    dbResult, err := db.MongoFindOne("cloudshareplatform", "user", bson.M{ "loginname": loginname })
    if err != nil {
        c.JSON(401, gin.H{ "status": 0, "message": "login name err", "data": "" }); return
    }
    
    // 获取数据库查询结果
    var dbResultData db.MongoUser
    err = dbResult.Decode(&dbResultData)
    if err != nil {
       c.JSON(401, gin.H{ "status": -1, "message": "db decode err", "data": "" }); return
    }
    
    // 判断密码正确性
    compareResult := utils.BcryptCompare(dbResultData.Password, password)
    if compareResult == true {
        c.JSON(200, gin.H{ "status": 1, "message": "ok", "data": dbResultData }); return
    } else {
        c.JSON(401, gin.H{ "status": -2, "message": "password compare err", "data": "" }); return
    }
}    