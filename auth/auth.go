package auth

import (
	"fmt"
	cfg "github.com/dxcenter/chess/config"
)

func init() {
	cfg.AddReloadHook(Rehash)
}

var userSources UserSources

type UserSources []UserSourceI
type UserSourceI interface {
	SignIn(login, password string) (loginFixed string, success bool)
}

func (userSources UserSources) SignIn(login, password string) (string, bool) {
	for _, userSource := range userSources {
		loginFixed, success := userSource.SignIn(login, password)
		if success {
			return loginFixed, success
		}
	}
	return "", false
}

func Rehash() {
	userSourcesRaw := cfg.Get().UserSources

	for _, userSourceRaw := range userSourcesRaw {
		switch userSourceRaw.Type {
		case "internal":
			userSources = append(userSources, NewInternalUserSource(userSourceRaw))
		case "db":
			userSources = append(userSources, NewDbUserSource(userSourceRaw))
		}
	}
	fmt.Println(userSources)
}

func SignIn(login, password string) (loginFixed string, success bool) {
	return userSources.SignIn(login, password)
}
