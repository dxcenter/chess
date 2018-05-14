package models

//go:generate reform

import (
	"fmt"
	"github.com/notnil/chess"
)

//reform:games
type playersPair struct {
	Id        int `reform:"id,pk"`
	PlayerId0 int `reform:"player_id_0"`
	PlayerId1 int `reform:"player_id_1"`
}

func getPlayersPair(playerId0, playerId1) playersPair {
	playersPair, err := PlayersPairSQL.First("WHERE ( player_id_0 = ? AND player_id_1 = ? ) OR ( player_id_0 = ? AND player_id_1 = ? )", playerId0, playerId1, playerId1, playerId0)
	if err == sql.ErrNoRow {
		playersPair = PlayersPair{
			PlayerId0: playerId0,
			PlayerId1: playerId1,
		}
		err = playersPair.Create()
	}
	if err != nil {
		panic(err)
	}

	return playersPair.Id
}
