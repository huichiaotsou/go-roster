package model

import "database/sql"

type Model struct {
	// DB
	Db *sql.DB
}

func NewModel() *Model {
	return &Model{}
}
