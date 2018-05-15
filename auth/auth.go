package auth

import (
	cfg "github.com/dxcenter/chess/config"
)

func init() {
	cfg.AddReloadHook(Rehash)
}

var userSources UserSources

type UserSourceBase struct {
	Name string
}
func (u UserSourceBase) GetName() string {
	return u.Name
}
func (u UserSourceBase) SetName(newName string) {
	u.Name = newName
}

type UserSources []UserSourceI
type UserSourceI interface {
	GetName() string
	SetName(string)
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

	for name, userSourceRaw := range userSourcesRaw {
		var newUserSource UserSourceI

		switch userSourceRaw.Type {
		case "internal":
			newUserSource = NewInternalUserSource(userSourceRaw)
		case "db":
			newUserSource = NewDbUserSource(userSourceRaw)
		}

		newUserSource.SetName(name)
		userSources = append(userSources, newUserSource)
	}
}

func SignIn(login, password string) (loginFixed string, success bool) {
	return userSources.SignIn(login, password)
}
