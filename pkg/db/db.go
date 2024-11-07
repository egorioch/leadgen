package db

import (
	"context"
	"embed"
	"fmt"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var migrationsFS embed.FS

type dbLogger struct {
	log *logrus.Logger
}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	if d.log.GetLevel() >= logrus.DebugLevel {
		queryText, _ := q.FormattedQuery()
		d.log.WithField("service", "DB").WithField("query", string(queryText)).Info("DB REQUEST")
	}

	return nil
}

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

	db.AddQueryHook(dbLogger{log})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("SKIP_MIGRATIONS") == "Y" {
		return db
	}
	migrations.Run(db, "init")

	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrationsFromFilesystem(http.FS(migrationsFS), "migrations")
	if err != nil {
		panic(err.Error())
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		panic(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
	fmt.Printf("migrations done\n")

	return db

}

func RunMigrations(db *pg.DB, log *logrus.Logger) {
	migrations.Run(db, "init")

	collection := migrations.NewCollection()
	err := collection.DiscoverSQLMigrationsFromFilesystem(http.FS(migrationsFS), "migrations")
	if err != nil {
		log.Fatalf("Error discovering migrations: %s", err)
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		log.Fatalf("Error running migrations: %s", err)
	}

	if newVersion != oldVersion {
		log.Infof("Migrated from version %d to %d", oldVersion, newVersion)
	} else {
		log.Infof("No new migrations to apply. Current version is %d", oldVersion)
	}
}
