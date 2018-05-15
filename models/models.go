package models

//go:generate reform

import (
	"github.com/xaionaro/reform"
)

type modelI interface {
	SetDefaultDB(*reform.DB) error
	Table() reform.Table
}

func List() []modelI {
	return []modelI{
		&email{},
		&emailConfirmation{},
		&game{},
		&player{},
		&playersPair{},
		&watcher{},
	}
}

func Init(db *reform.DB) {
	for _, model := range List() {
		model.SetDefaultDB(db)
		model.Table().CreateTableIfNotExists(db)
	}
}
