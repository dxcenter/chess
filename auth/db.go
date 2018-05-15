package auth

import (
	"database/sql"
	cfg "github.com/dxcenter/chess/config"
	db "github.com/dxcenter/chess/db"
	m "github.com/dxcenter/chess/models"
	"github.com/xaionaro/reform"
)

func init() {
}

type DbUserSource struct {
	db *reform.DB
}

func NewDbUserSource(userSourceRaw cfg.UserSource) (result DbUserSource) {
	dbBlockName := userSourceRaw.Data.ToDbUserSourceData()
	result.db = db.GetDB(string(dbBlockName))
	return
}

func (s DbUserSource) SignIn(login, password string) (string, bool) {
	passwordHash := m.HashPassword(password)
	_, err := m.Player.DB(s.db).First(m.PlayerF{Nickname: &login, PasswordHash: &passwordHash})
	if err == sql.ErrNoRows {
		return "", false
	}
	if err != nil {
		panic(err)
		return "", false
	}
	return login, true
}
