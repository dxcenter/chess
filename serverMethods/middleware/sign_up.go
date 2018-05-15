package serverMethodsMiddleware

import (
	"database/sql"
	auth "github.com/dxcenter/chess/auth"
	m "github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpParams struct {
	 Login    string `json:"username"`
	 Password string `json:"password"`
	 Email    string `json:"email"`
}

func SignUp(c *gin.Context) {
	var json signUpParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Email != "" {
		_, err := m.Email.First(m.EmailF{Address: json.Email})
		switch err {
		case sql.ErrNoRows:
		case nil:
			c.JSON(http.StatusBadRequest, gin.H{"error": "The email is already registered"})
			return
		default:
			panic(err)
		}
	}

	userSource := auth.GetInternalDynamicUserSource()
	if userSource == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sign up is disabled (no dynamic user DB)"})
		return
	}

	tx, err := m.Player.StartTransaction()
	if err != nil {
		panic(err)
	}

	passwordHash := m.HashPassword(json.Password)
	userSourceName := userSource.GetName()
	player := m.NewPlayer()
	player.Nickname = &json.Login
	player.PasswordHash = &passwordHash
	player.Source = &userSourceName
	err = player.DB(tx).Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create user (the nickname is already used?)"})
		return
	}

	email := m.NewEmail()
	email.Address = json.Email
	email.PlayerId = player.Id
	err = email.DB(tx).Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email is already registered"})
		return
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	c.Next()
}
