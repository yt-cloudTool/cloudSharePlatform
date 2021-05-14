package routes

import (
	zoom "cloudSharePlatform/zoom"

	middleware "cloudSharePlatform/zoom/middleware"
)

func init_file() {
	// 文件上传 POST
	r.POST("/api/fileupload", middleware.HandleTokenMid, zoom.FileUpload)

	// 文件获取 GET
	r.GET("/api/filedownload", zoom.FileDownload)

	// 获取用户临时文件
	r.POST("/api/fileGetTmp", middleware.HandleTokenMid, zoom.FileGetTmp)

	// 生成 fileBox
	r.POST("/api/fileboxcreate", middleware.HandleTokenMid, zoom.FileBoxCreate)

	// 删除filebox
	r.POST("/api/fileboxdelete", middleware.HandleTokenMid, zoom.FileBoxDelete)

	// 文件添加到filebox
	r.POST("/api/fileboxinsertinto", middleware.HandleTokenMid, zoom.FileBoxInsertInto)

	// 文件从filebox中删除
	r.POST("/api/fileboxdeleteitem", middleware.HandleTokenMid, zoom.FileBoxDeleteItem)

	// 检索
	r.GET("/api/fileboxquery", middleware.HandleTokenMid, zoom.FileBoxQuery)
}
