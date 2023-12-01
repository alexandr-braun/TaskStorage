package postgresql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
	"taskStorage/pkg/infrastructure_abstractions"
)

func RunMigrations(dbConnectionFactory infrastructure_abstractions.DBConnectionFactory) error {
	var dbConnection, dbConErr = dbConnectionFactory.NewConnection()
	if dbConErr != nil {
		log.Fatal("Error connecting to DB:", dbConErr)
	}

	if err := goose.Up(dbConnection, "pkg/infrastructure/migrations"); err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	defer func(dbConnection *sql.DB) {
		closeErr := dbConnection.Close()
		if closeErr != nil {
			log.Printf("Failed to close database connection: %v", closeErr)
		}
	}(dbConnection)

	return nil
}
