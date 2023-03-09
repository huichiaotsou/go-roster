package middleware

import (
	"github.com/huichiaotsou/go-roster/model"
)

type Middleware struct {
	Db *model.Database
}

func New(db *model.Database) *Middleware {
	return &Middleware{
		Db: db,
	}
}
