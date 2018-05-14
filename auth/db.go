package auth

import (
	"database/sql"
	cfg "github.com/dxcenter/chess/config"
	db "github.com/dxcenter/chess/db"
	m "github.com/dxcenter/chess/models"
	"github.com/xaionaro/reform"
	"strconv"
)

func init() {
}

type DbUserSource struct {
	db *reform.DB
}

func NewDbUserSource(userSourceRaw cfg.UserSource) (result DbUserSource) {
	d := userSourceRaw.Data.ToDbUserSourceData()
	result.db = db.InitDB(db.InitDBParams{
		Driver:   d.Driver,
		Protocol: d.Protocol,
		Host:     d.Host,
		Port:     d.Port,
		Db:       d.Db,
		User:     d.User,
		Password: d.Password,
		Path:     d.Path,
	})
	return
}

func (s DbUserSource) SignIn(login, password string) (string, bool) {
	passwordHash := m.HashPassword(password)
	player, err := m.Player.DB(s.db).First(m.PlayerF{Nickname: &login, PasswordHash: &passwordHash})
	if err == sql.ErrNoRows {
		return "", false
	}
	if err != nil {
		panic(err)
		return "", false
	}
	return strconv.Itoa(player.GetPlayerId()), true
}
