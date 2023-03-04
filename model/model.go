package model

import (
	"github.com/jmoiron/sqlx"
)

type Model struct {
	Db *sqlx.DB
}

func NewModel(db *sqlx.DB) *Model {
	return &Model{
		Db: db,
	}
}
