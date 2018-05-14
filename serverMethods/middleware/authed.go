package serverMethodsMiddleware

import (
	m "github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
)

func Authed(c *gin.Context) {
	me, err := m.Player.First(c.GetInt("playerId"))
	if err != nil {
		panic(err)
	}
	c.Set("me", me)
	c.Next()
}
