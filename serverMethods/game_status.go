package serverMethods

import (
	g "github.com/dxcenter/chess/game"
	"github.com/gin-gonic/gin"
)

func GameStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"GameStatus": g.GetStatus(),
	})
}
