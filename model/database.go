package model

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Sqlx *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Sqlx: db,
	}
}

func (db *Database) Close() {
	db.Sqlx.Close()
}
