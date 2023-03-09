package middleware

import (
	"github.com/huichiaotsou/go-roster/model"
	log "github.com/sirupsen/logrus"
)

type Middleware struct {
	Db     *model.Database
	Logger *log.Logger
}

func New(db *model.Database, logger *log.Logger) *Middleware {
	return &Middleware{
		Db:     db,
		Logger: logger,
	}
}
