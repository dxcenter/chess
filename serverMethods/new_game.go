package serverMethods

import (
	g "github.com/dxcenter/chess/game"
	"github.com/gin-gonic/gin"
)

func NewGame(c *gin.Context) {
	g.NewGame()

	c.JSON(200, gin.H{
		"GameStatus": g.GetStatus(),
	})
}
