package routes

import (
	zoom "cloudSharePlatform/zoom"

	middleware "cloudSharePlatform/zoom/middleware"
)

func init_article() {

	// 新增文章 POST
	r.POST("/api/articleinsert", middleware.HandleTokenMid, zoom.ArticleInsert)

	// 获取桌面数据 GET
	r.GET("/api/articlelist", middleware.HandleTokenMid, zoom.ArticleListGet)
}
