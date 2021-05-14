package routes

import (
	zoom "cloudSharePlatform/zoom"
	middleware "cloudSharePlatform/zoom/middleware"
	// "github.com/gin-gonic/gin"
)

func init_userAuthority() {
	// 注册
	r.POST("/api/register", zoom.UserRegister)

	// 登录
	r.POST("/api/login", zoom.UserLogin)

	// 检查登录
	r.POST("/api/checklogin", middleware.HandleTokenMid, zoom.UserCheckLogin)
}
