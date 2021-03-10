package zoom

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}
