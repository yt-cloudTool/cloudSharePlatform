package routes

import (
	zoom "cloudSharePlatform/zoom"

	middleware "cloudSharePlatform/zoom/middleware"
)

func init_file() {
	/* *************************************************************************
	                file
	************************************************************************* */
	// 文件上传 POST
	r.POST("/api/file/upload", middleware.HandleTokenMid, zoom.FileUpload)

	// 文件获取 GET
	r.GET("/api/file/download", zoom.FileDownload)

	// 文件删除 POST
	r.POST("/api/file/delete", middleware.HandleTokenMid, zoom.FileDelete)

	// 获取用户临时文件 GET
	r.GET("/api/file/tmpAll", middleware.HandleTokenMid, zoom.FileTmpGetAll)

	/* *************************************************************************
	                fileBox
	************************************************************************* */
	// 生成 fileBox
	r.POST("/api/filebox/create", middleware.HandleTokenMid, zoom.FileBoxCreate)

	// 删除filebox
	r.POST("/api/filebox/delete", middleware.HandleTokenMid, zoom.FileBoxDelete)

	// 文件添加到filebox
	r.POST("/api/filebox/insertInto", middleware.HandleTokenMid, zoom.FileBoxInsertInto)

	// 文件从filebox中删除
	r.POST("/api/filebox/deleteItem", middleware.HandleTokenMid, zoom.FileBoxDeleteItem)

	// 检索
	r.GET("/api/filebox/query", middleware.HandleTokenMid, zoom.FileBoxQuery)
}
