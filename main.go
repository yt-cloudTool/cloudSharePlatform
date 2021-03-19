package main

import (
    "fmt"
	routes "cloudSharePlatform/routes"
    db "cloudSharePlatform/db"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine = gin.Default()

func main() {
    /*
        数据库测试
    */
    result, err := db.MongoInsertOne("cloudshareplatform", "user", db.MongoUser{})
    if err != nil {
        fmt.Println("err ==========>", err)
    }
    fmt.Println("insertOne result ===============>", result.InsertedID)
    
    
	/*
		已使用的路由
	*/
	// 用户注册
	routes.UserAuthorityRegister(r)

	r.Run(":8000")
}
