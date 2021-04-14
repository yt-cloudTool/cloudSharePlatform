package routes

import (
	zoom "cloudSharePlatform/zoom"

	middleware "cloudSharePlatform/zoom/middleware"

	"github.com/gin-gonic/gin"
)

// 文件上传 POST
func FileUpload(r *gin.Engine) {
	r.POST("/api/fileupload", middleware.HandleTokenMid, zoom.FileUpload)
}

// 文件获取 GET
func FileDownload(r *gin.Engine) {
	// r.GET("/api/filedownload", zoom.FileDownload)
}
