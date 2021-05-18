package routes

import (
	zoom "cloudSharePlatform/zoom"

	middleware "cloudSharePlatform/zoom/middleware"
)

func init_article() {

	// 新增文章 POST
	r.POST("/api/article/insert", middleware.HandleTokenMid, zoom.ArticleInsert)

	// 获取桌面数据 GET
	r.GET("/api/article/list", middleware.HandleTokenMid, zoom.ArticleListGet)
}
