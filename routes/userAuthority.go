package routes

import (
	zoom "cloudSharePlatform/zoom"
	middleware "cloudSharePlatform/zoom/middleware"

	"github.com/gin-gonic/gin"
)

// 注册
func UserAuthorityRegister(r *gin.Engine) {
	r.POST("/api/register", zoom.UserRegister)
}

// 登录
func UserAuthorityLogin(r *gin.Engine) {
	r.POST("/api/login", zoom.UserLogin)
}

// 检查登录
func UserAuthorityCheckLogin(r *gin.Engine) {
	r.POST("/api/checklogin", middleware.HandleTokenMid, zoom.UserCheckLogin)
}
