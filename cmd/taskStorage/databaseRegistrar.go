package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// TODO Move to config file
const (
	host     = "localhost"
	port     = 5432
	user     = "task_storage"
	password = "task_storage"
	dbname   = "task_storage"
)

type DatabaseRegistrar struct {
}

func NewDatabaseRegistrar() *DatabaseRegistrar {
	return &DatabaseRegistrar{}
}

func (dr *DatabaseRegistrar) addDatabase() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
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

	fmt.Println("Successfully connected!")
}
