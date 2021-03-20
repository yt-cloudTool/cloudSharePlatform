package routes

import (
	zoom "cloudSharePlatform/zoom"

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
