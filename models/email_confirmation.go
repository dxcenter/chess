package models

//go:generate reform

//reform:email_confirmations
type emailConfirmation struct {
	Id       int    `reform:"id,pk"`
	Code     string `reform:"code"`
	EmailId  int    `reform:"email_id"`
	PlayerId int    `reform:"player_id"`
}
