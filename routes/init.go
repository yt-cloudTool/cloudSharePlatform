package routes

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func SetRoutesObj(routes *gin.Engine) {
	r = routes
}

func ExeRoutesObj() {
	init_userAuthority() // 用户权限相关
	init_serverInfo()    // 服务器信息
	init_file()          // 文件相关
	init_article()       // 文章相关
}
