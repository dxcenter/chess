package serverMethods

import (
	"database/sql"
	"github.com/dxcenter/chess/serverMethods/helpers"
	"github.com/gin-gonic/gin"
	//"net/http"
	"strconv"
)

/*type playersPairsParams struct {
	ActiveOnly bool `json:"active_only,omitempty"`
	MyOnly bool `json:"my_only,omitempty"`
}*/

func PlayersPairs(c *gin.Context) {
	/*var json playersPairsParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/

	activeOnly, _ := strconv.ParseBool(c.Param("activeOnly"))
	myOnly, _ := strconv.ParseBool(c.Param("myOnly"))

	me := helpers.GetMe(c)
	scope := me.VisiblePlayersPairsScope()

	if activeOnly {
		// TODO: do something here
	}
	if myOnly {
		myId := me.GetPlayerId()
		scope = scope.Where("player_id_0 = ? OR player_id_1 = ?", myId, myId)
	}

	playersPairs, err := scope.Select()
	switch err {
	case nil, sql.ErrNoRows:
	default:
		panic(err)
	}

	c.JSON(200, gin.H{
		"PlayersPairs": playersPairs,
	})
}
