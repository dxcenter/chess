package serverMethodsMiddleware

import (
	"bytes"
	"database/sql"
	auth "github.com/dxcenter/chess/auth"
	m "github.com/dxcenter/chess/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type signUpParams struct {
	Login    string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ClosingBuffer struct {
	*bytes.Buffer
}
func (cb *ClosingBuffer) Close() (error) {
	return nil
}

func SignUp(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	c.Request.Body = &ClosingBuffer{bytes.NewBuffer(body)}

	var json signUpParams
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Request.Body = &ClosingBuffer{bytes.NewBuffer(body)}

	if json.Email != "" {
		_, err := m.Email.First(m.EmailF{Address: json.Email})
		switch err {
		case sql.ErrNoRows:
		case nil:
			c.JSON(http.StatusBadRequest, gin.H{"error": "The email is already registered"})
			c.Abort()
			return
		default:
			panic(err)
		}
	}

	userSource := auth.GetInternalDynamicUserSource()
	if userSource == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sign up is disabled (no dynamic user DB)"})
		c.Abort()
		return
	}
	userSourceName := userSource.GetName()

	if userSourceName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sign up is disabled (no dynamic user DB, case 2)"})
		c.Abort()
		return
	}

	/*tx, err = m.Player.StartTransaction()
	if err != nil {
		panic(err)
	}*/

	passwordHash := m.HashPassword(json.Password)
	player := m.NewPlayer()
	player.Nickname = &json.Login
	player.PasswordHash = &passwordHash
	player.Source = &userSourceName
	err = player.Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create user (the nickname is already used?)"})
		c.Abort()
		return
	}

	email := m.NewEmail()
	email.Address = json.Email
	email.PlayerId = player.Id
	err = email.Insert()
	if err != nil {
		player.Delete()
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email is already registered"})
		c.Abort()
		return
	}

	/*err = tx.Commit()
	if err != nil {
		panic(err)
	}*/

	c.Next()
}
