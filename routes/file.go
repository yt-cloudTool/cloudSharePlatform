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
	r.GET("/api/filedownload", zoom.FileDownload)
}

// 生成 fileBox
func FileBoxCreate(r *gin.Engine) {
	r.POST("/api/fileboxcreate", middleware.HandleTokenMid, zoom.FileBoxCreate)
}

// 删除filebox
func FileBoxDelete(r *gin.Engine) {
	r.POST("/api/fileboxdelete", middleware.HandleTokenMid, zoom.FileBoxDelete)
}

// 文件添加到filebox
func FileBoxInsertInto(r *gin.Engine) {
	r.POST("/api/fileboxinsertinto", middleware.HandleTokenMid, zoom.FileBoxInsertInto)
}

// 文件从filebox中删除
func FileBoxDeleteItem(r *gin.Engine) {
	r.POST("/api/fileboxdeleteitem", middleware.HandleTokenMid, zoom.FileBoxDeleteItem)
}

// 检索
func FileBoxQuery(r *gin.Engine) {
	r.GET("/api/fileboxquery", middleware.HandleTokenMid, zoom.FileBoxQuery)
}
