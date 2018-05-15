package serverMethodsMiddleware

import (
	"database/sql"
	m "github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authed(c *gin.Context) {
	nickname := c.GetString("userID");
	if nickname == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": `nickname == ""`})
		c.Abort()
		return
	}
	me, err := m.Player.First(m.PlayerF{Nickname: &nickname})
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}
	if err != nil {
		panic(err)
	}
	c.Set("me", me)
	c.Next()
}
