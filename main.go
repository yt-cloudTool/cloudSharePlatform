package main

import (
    // "fmt"
	routes "cloudSharePlatform/routes"
    // db "cloudSharePlatform/db"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine = gin.Default()

func main() {
	/*
		已使用的路由
	*/
	routes.UserAuthorityRegister(r)     // 用户注册
	routes.UserAuthorityLogin(r)        // 用户登录

	r.Run(":8000")
}
