package models

//go:generate reform

//reform:emails
type email struct {
	Id       int    `reform:"id,pk"`
	Address  string `reform:"address" sql:"unique_index" sql_size:"255"`
	PlayerId int    `reform:"player_id" sql:"index"`
}

func NewEmail() *email {
	return &email{}
}
