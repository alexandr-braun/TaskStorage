package infrastructure

import (
	"github.com/pressly/goose/v3"
	"log"
)

func RunMigrations(databaseConfig DatabaseConfig) {
	var connect, _ = NewPostgresqlConnectionFactory().NewConnection(databaseConfig)
	if err := goose.Up(connect, "pkg/infrastructure/migrations"); err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	// TODO handle the error
	defer connect.Close()
}
