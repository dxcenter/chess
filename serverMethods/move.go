package serverMethods

import (
	g "github.com/dxcenter/chess/game"
	"github.com/gin-gonic/gin"
)

func Move(c *gin.Context) {
	move := c.Param("move")
	moveError := g.MoveStr(move)
	c.JSON(200, gin.H{
		"GameStatus": g.GetStatus(),
		"MoveError":  moveError,
	})
}
