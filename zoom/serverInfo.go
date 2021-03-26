package zoom

import (
	// "fmt"
	"runtime"

	gin "github.com/gin-gonic/gin"
	// bson "go.mongodb.org/mongo-driver/bson"
)

// 获取服务器内存 CPU GET
func ServerInfo(c *gin.Context) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	total := mem.Sys

	c.JSON(200, gin.H{"status": 1, "message": "ok", "data": gin.H{
		"memTotal": total,
		"numCPU":   runtime.NumCPU(),
		"mem":      mem,
	}})
}
