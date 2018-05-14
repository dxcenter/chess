package models

//go:generate reform

//reform:emails
type email struct {
	Id       int    `reform:"id,pk"`
	Address  string `reform:"address,unique"`
	PlayerId int    `reform:"player_id"`
}