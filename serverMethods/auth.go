package serverMethods

import (
	cfg "github.com/dxcenter/chess/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type LoginPass struct {
	Login    string `json:"login"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Auth(c *gin.Context) {
	var json LoginPass
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range cfg.Get().Users {
		if strings.ToLower(json.Login) == strings.ToLower(user.Login) && json.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
}
