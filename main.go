package main

import (
    "net/http"
    // "log"
    // "fmt"
	routes "cloudSharePlatform/routes"
    // db "cloudSharePlatform/db"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine = gin.Default()

func main() {
    // ==========================================================================
    // 允许跨域
    r.Use(func(context *gin.Context) {
        method := context.Request.Method
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "false")
		// context.Set("content-type", "application/json")  // 设置返回格式是json
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "{}")
		}
		//处理请求
		context.Next()
    })
    // ==========================================================================

    /*
		已使用的路由
	*/
	routes.UserAuthorityRegister(r)     // 用户注册
	routes.UserAuthorityLogin(r)        // 用户登录
	r.Run(":8000")
}
