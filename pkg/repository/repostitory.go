package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type S struct {
	log *logrus.Logger
	db  *pg.DB
}

func New(log *logrus.Logger, db *pg.DB) *S {
	return &S{
		log: log,
		db:  db,
	}
}
