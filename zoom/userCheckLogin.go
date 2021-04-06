package zoom

import (
	gin "github.com/gin-gonic/gin"
)

/*
   header传token
*/
func UserCheckLogin(c *gin.Context) {

	user_id, isExist := c.Get("user_id")
	if isExist == false {
		c.JSON(200, gin.H{"status": 401, "message": "no user_id", "data": ""})
		return
	}

	// 操作articleList表
	c.JSON(200, gin.H{"status": 0, "message": "ok", "data": user_id})
}
