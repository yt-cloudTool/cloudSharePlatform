package main

import (
	routes "cloudSharePlatform/routes"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine = gin.Default()

func main() {
	/*
		已使用的路由
	*/
	// 用户注册
	routes.UserAuthorityRegister(r)

	r.Run(":8000")
}
