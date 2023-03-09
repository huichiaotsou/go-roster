package model

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Sql *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Sql: db,
	}
}
