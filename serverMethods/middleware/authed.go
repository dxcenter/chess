package serverMethodsMiddleware

import (
	"github.com/gin-gonic/gin"
)

func Authed(c *gin.Context) {
	c.Next()
}
