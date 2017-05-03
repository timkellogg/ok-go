package tasks

import (
	"database/sql"
	"errors"
	"os"

	"github.com/DavidHuie/gomigrate"
	"github.com/chuckpreslar/gofer"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupEnv(arguments []string) error {
	if len(arguments) > 0 {
		return godotenv.Load(arguments[0])
	}
	return nil
}

func DBConnect() (*sql.DB, error) {
	godotenv.Load()
	db, dbError := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return db, dbError
}

func SetupMigrator() (*gomigrate.Migrator, error) {
	db, dbError := DBConnect()
	if dbError != nil {
		return nil, errors.New("DB connection failed")
	}
	migrator, migError := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./migrations")
	return migrator, migError
}

var DBMigrate = gofer.Register(gofer.Task{
	Namespace:   "db",
	Label:       "migrate",
	Description: "Migrates a database",
	Action: func(arguments ...string) error {

		loadError := SetupEnv(arguments)
		if loadError != nil {
			return errors.New("env file does not exist")
		}

		migrator, migError := SetupMigrator()
		if migError != nil {
			return migError
		}

		migrateError := migrator.Migrate()
		if migrateError != nil {
			return errors.New("Migration failed")
		}

		return nil
	},
})

var DBRollback = gofer.Register(gofer.Task{
	Namespace:   "db",
	Label:       "rollback",
	Description: "Rolls back a database",
	Action: func(arguments ...string) error {

		loadError := SetupEnv(arguments)
		if loadError != nil {
			return errors.New("env file does not exist")
		}

		migrator, migError := SetupMigrator()
		if migError != nil {
			return migError
		}

		migrateError := migrator.Rollback()
		if migrateError != nil {
			return errors.New("Migration failed")
		}

		return nil
	},
})
