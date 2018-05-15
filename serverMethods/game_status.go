package serverMethods

import (
	m "github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*type gameStatusParams struct {
	GameId int `json:"game_id"`
}*/

func GameStatus(c *gin.Context) {
	/*var json gameStatusParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gameId := json.GameId
	*/

	gameId, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"GameStatus": m.GetGame(gameId).GetStatus(),
	})
}
