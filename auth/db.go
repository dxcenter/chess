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
	UserSourceBase
	db *reform.DB
}

func NewDbUserSource(userSourceRaw cfg.UserSource) (result *DbUserSource) {
	result = &DbUserSource{}
	dbBlockName := userSourceRaw.Data.ToDbUserSourceData()
	result.db = db.GetDB(string(dbBlockName))
	return
}

func (s DbUserSource) SignIn(login, password string) (string, bool) {
	player, err := m.Player.DB(s.db).First(m.PlayerF{Nickname: &login})
	if err == sql.ErrNoRows {
		return "", false
	}
	if err != nil {
		panic(err)
		return "", false
	}
	return login, player.CheckPassword([]byte(password))
}

func GetInternalDynamicUserSource() *DbUserSource {
	myInternalDb := db.GetDB(cfg.Get().MyDb)

	for _, userSourceI := range userSources {
		userSource, ok := userSourceI.(*DbUserSource)
		if !ok {
			continue
		}
		if userSource.db == myInternalDb {
			return userSource
		}
	}

	return nil
}
