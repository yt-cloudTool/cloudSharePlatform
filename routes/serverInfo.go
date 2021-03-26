package routes

import (
	zoom "cloudSharePlatform/zoom"

	"github.com/gin-gonic/gin"
)

// 获取服务器内存信息
func ServerInfoMemory(r *gin.Engine) {
	r.GET("/api/serverinfo", zoom.ServerInfo)
}
