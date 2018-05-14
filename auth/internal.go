package auth

import (
	cfg "github.com/dxcenter/chess/config"
	"strings"
)

func init() {
}

type InternalUserSource struct {
	Users cfg.Users
}

func NewInternalUserSource(userSourceRaw cfg.UserSource) (result InternalUserSource) {
	result.Users = userSourceRaw.Data.ToInternalUserSourceData().Users
	return
}

func (s InternalUserSource) SignIn(login, password string) (string, bool) {
	login = strings.ToLower(login)
	for _, user := range s.Users {
		if login == strings.ToLower(user.Login) && password == user.Password {
			return login, true
		}
	}

	return login, false
}
