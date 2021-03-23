package zoom

import (
    gin "github.com/gin-gonic/gin"
    db "cloudSharePlatform/db"
    // bson "go.mongodb.org/mongo-driver/bson"
    // options "go.mongodb.org/mongo-driver/mongo/options"
    utils "cloudSharePlatform/utils"
)

func UserRegister(c *gin.Context) {
    loginname := c.PostForm("loginname")
    password  := c.PostForm("password")
    
    // 判断参数
    if loginname == "" || password == "" {
        c.JSON(401, gin.H{ "status": 0, "message": "post data not enough", "data": "" }); return
    }
    
    // 生成密码
    generatedPwd, err := utils.BcryptGenerate(password)
    if err != nil {
        c.JSON(401, gin.H{ "status": -1, "message": "hash password generate err", "data": "" }); return
    }
    
    // 设置数据
    userDbStoreData := db.MongoUser{
        LoginName: loginname,
        Nickname: loginname,
        Password: generatedPwd,
        Access: 2, // 默认普通用户权限
    }
    
    // 添加用户
    dbResult, err := db.MongoInsertOne("cloudshareplatform", "user", userDbStoreData)
    if err != nil {
        c.JSON(401, gin.H{ "status": -2, "message": "login name err", "data": "" }); return
    }
    c.JSON(200, gin.H{ "status": 1, "message": "ok", "data": dbResult }); return
}