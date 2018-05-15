package models

//go:generate reform

import (
	"database/sql"
)

//reform:players_pairs
type playersPair struct {
	Id        int `reform:"id,pk"`
	PlayerId0 int `reform:"player_id_0" sql:"index"`
	PlayerId1 int `reform:"player_id_1" sql:"index"`
}

func getPlayersPair(playerId0, playerId1 int) playersPair {
	pair, err := PlayersPair.First("( player_id_0 = ? AND player_id_1 = ? ) OR ( player_id_0 = ? AND player_id_1 = ? )", playerId0, playerId1, playerId1, playerId0)
	if err == sql.ErrNoRows {
		pair = playersPair{
			PlayerId0: playerId0,
			PlayerId1: playerId1,
		}
		err = pair.Create()
	}
	if err != nil {
		panic(err)
	}

	return pair
}
