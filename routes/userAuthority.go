package routes

import (
	zoom "cloudSharePlatform/zoom"

	"github.com/gin-gonic/gin"
)

func UserAuthorityRegister(r *gin.Engine) {
	r.POST("/api/register", zoom.UserRegister)
}
