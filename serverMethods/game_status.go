package serverMethods

import (
	m "github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type gameStatusParams struct {
	GameId int `json:"game_id"`
}

func GameStatus(c *gin.Context) {
	var json gameStatusParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"GameStatus": m.GetGame(json.GameId).GetStatus(),
	})
}
