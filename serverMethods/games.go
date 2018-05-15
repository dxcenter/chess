package serverMethods

import (
	"database/sql"
	"github.com/dxcenter/chess/serverMethods/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type gamesParams struct {
}

func Games(c *gin.Context) {
	var json gamesParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	me := helpers.GetMe(c)

	games, err := me.VisibleGamesScope().Select()
	switch err {
	case sql.ErrNoRows:
	default:
		panic(err)
	}

	c.JSON(200, gin.H{
		"Games": games,
	})
}
