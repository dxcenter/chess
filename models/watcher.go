package models

//go:generate reform

//reform:watchers
type watcher struct {
	Id       int `reform:"id,pk"`
	PlayerId int `reform:"player_id" sql:"index"`
	GameId   int `reform:"game_id" sql:"index"`
}
