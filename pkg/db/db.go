package db

import (
	"embed"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

var migrationsFS embed.FS

func NewPgDb(
	log *logrus.Logger,
	address string,
	user string,
	password string,
	database string,
) *pg.DB {
	log.Info(fmt.Sprintf(
		"Connect to db: %s address: %s user: %s",
		database,
		address,
		user,
	))
	db := pg.Connect(&pg.Options{
		Addr:     address,
		User:     user,
		Password: password,
		Database: database,
	})

	return db
}
