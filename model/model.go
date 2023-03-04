package model

import (
	"database/sql"
)

type Model struct {
	Db *sql.DB
}

func NewModel(db *sql.DB) *Model {
	return &Model{
		Db: db,
	}
}
