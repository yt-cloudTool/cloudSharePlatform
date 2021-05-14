package routes

import (
	zoom "cloudSharePlatform/zoom"
)

// 获取服务器内存信息
func init_serverInfo() {
	r.GET("/api/serverinfo", zoom.ServerInfo)
}
