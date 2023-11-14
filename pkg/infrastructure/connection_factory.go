package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "taskStorage/pkg/infrastructure/migrations"
)

type DBConnectionFactory interface {
	NewConnection(databaseConfig DatabaseConfig) (*sql.DB, error)
}

type PostgresqlConnectionFactory struct {
}

func NewPostgresqlConnectionFactory() DBConnectionFactory {
	return &PostgresqlConnectionFactory{}
}

func (*PostgresqlConnectionFactory) NewConnection(databaseConfig DatabaseConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Password, databaseConfig.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
