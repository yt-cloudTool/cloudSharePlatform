package routes

import (
	zoom "cloudSharePlatform/zoom"

	middleware "cloudSharePlatform/zoom/middleware"

	"github.com/gin-gonic/gin"
)

// 新增文章 POST
func ArticleInsert(r *gin.Engine) {
	r.POST("/api/articleinsert", middleware.HandleTokenMid, zoom.ArticleInsert)
}

// 获取桌面数据 GET
func ArticleListGet(r *gin.Engine) {
	r.GET("/api/articlelist", middleware.HandleTokenMid, zoom.ArticleListGet)
}
