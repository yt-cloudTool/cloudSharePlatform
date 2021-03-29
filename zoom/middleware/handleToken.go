package middleware

import (
	"cloudSharePlatform/utils"

	"github.com/gin-gonic/gin"
)

func HandleTokenMid(c *gin.Context) {
	headerToken := c.GetHeader("token")

	// 解析token
	user_id, err := utils.JwtValidate(headerToken)
	if err != nil {
		c.JSON(401, gin.H{"status": 0, "message": "access err", "data": ""})
		c.Abort()
		return
	}
	c.Set("user_id", user_id)
	c.Next()
}
