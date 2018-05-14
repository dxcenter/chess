package helpers

import (
	m "github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) m.PlayerI {
	me, ok := c.Get("me")
	if !ok {
		panic(`"me" is not set :(`)
	}
	return me.(m.PlayerI)
}
