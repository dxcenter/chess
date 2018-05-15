package serverMethods

import (
	"database/sql"
	"github.com/dxcenter/chess/serverMethods/helpers"
	"github.com/gin-gonic/gin"
	//"net/http"
)

/*type playersParams struct {
}*/

func Players(c *gin.Context) {
	/*var json playersParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/

	me := helpers.GetMe(c)

	players, err := me.VisiblePlayersScope().Select()
	switch err {
	case nil, sql.ErrNoRows:
	default:
		panic(err)
	}

	if players == nil {
		panic(`players == nil`)
	}

	c.JSON(200, gin.H{
		"Players": players,
	})
}
