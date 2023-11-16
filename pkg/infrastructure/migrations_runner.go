package infrastructure

import (
	"github.com/pressly/goose/v3"
	"log"
)

func RunMigrations(dbConnectionFactory DBConnectionFactory) error {
	var dbConnection, _ = dbConnectionFactory.NewConnection()
	if err := goose.Up(dbConnection, "pkg/infrastructure/migrations"); err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	// TODO handle the error
	defer dbConnection.Close()

	return nil
}
