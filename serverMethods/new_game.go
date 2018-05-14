package serverMethods

import (
	"github.com/dxcenter/chess/serverMethods/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newGameParams struct {
	InvitedPlayerId int `json:"invitied_player_id,omitempty"`
}

func NewGame(c *gin.Context) {
	var json newGameParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	me := helpers.GetMe(c)
	game := me.NewGame(json.InvitedPlayerId)

	c.JSON(200, gin.H{
		"GameStatus": game.GetStatus(),
	})
}
