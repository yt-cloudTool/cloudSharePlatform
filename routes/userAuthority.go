package routes

import (
	zoom "cloudSharePlatform/zoom"
	middleware "cloudSharePlatform/zoom/middleware"
	// "github.com/gin-gonic/gin"
)

func init_userAuthority() {
	// 注册
	r.POST("/api/user/register", zoom.UserRegister)

	// 登录
	r.POST("/api/user/login", zoom.UserLogin)

	// 检查登录
	r.POST("/api/user/checkLogin", middleware.HandleTokenMid, zoom.UserCheckLogin)
}
