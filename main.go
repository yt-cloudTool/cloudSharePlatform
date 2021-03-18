package main

import (
	routes "cloudSharePlatform/routes"
    db "cloudSharePlatform/db"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine = gin.Default()

func main() {
    /*
        数据库初始化
    */
    db.MongodbInit()
    
	/*
		已使用的路由
	*/
	// 用户注册
	routes.UserAuthorityRegister(r)

	r.Run(":8000")
}
