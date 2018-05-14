package serverMethods

import (
	"fmt"
	g "github.com/dxcenter/chess/game"
	"github.com/gin-gonic/gin"
	"net/http"
)

type moveParams struct {
	Move string `json:"move"`
}

func Move(c *gin.Context) {
	var json moveParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	moveError := g.MoveStr(json.Move)
	if moveError != nil {
		c.JSON(200, gin.H{
			"GameStatus": g.GetStatus(),
			"MoveError":  moveError.Error(),
		})
	}
	fmt.Println("move", json.Move, moveError)
	c.JSON(200, gin.H{
		"GameStatus": g.GetStatus(),
	})
}
