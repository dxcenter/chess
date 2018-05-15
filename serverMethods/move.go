package serverMethods

import (
	m "github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type moveParams struct {
	//GameId int    `json:"game_id"`
	Move   string `json:"move"`
}

func Move(c *gin.Context) {
	var json moveParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//gameId := json.GameId
	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game := m.GetGame(gameId)
	moveError := game.MoveStr(json.Move)
	if moveError != nil {
		c.JSON(200, gin.H{
			"GameStatus": game.GetStatus(),
			"MoveError":  moveError.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"GameStatus": game.GetStatus(),
	})

}
