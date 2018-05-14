package serverMethods

import (
	"github.com/gin-gonic/gin"
)

func Whoami(c *gin.Context) {
	c.JSON(200, gin.H{
		"Login": c.GetString("userID"),
	})
}
