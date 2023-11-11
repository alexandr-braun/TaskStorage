package main

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"

	_ "github.com/lib/pq"
	_ "taskStorage/pkg/infrastructure/migrations"
)

type DatabaseRegistrar struct {
}

func NewDatabaseRegistrar() *DatabaseRegistrar {
	return &DatabaseRegistrar{}
}

func (dr *DatabaseRegistrar) connectToDatabase(databaseConfig Database) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Password, databaseConfig.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "pkg/infrastructure/migrations"); err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	fmt.Println("Successfully connected!")
}
