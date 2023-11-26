package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"taskStorage/pkg/configuration"
	_ "taskStorage/pkg/infrastructure/migrations"
)

type DBConnectionFactory interface {
	NewConnection() (*sql.DB, error)
}

type PostgresqlConnectionFactory struct {
	databaseConfig configuration.DatabaseConfig
}

func NewPostgresqlConnectionFactory(config *configuration.Config) DBConnectionFactory {
	return &PostgresqlConnectionFactory{
		databaseConfig: config.Database,
	}
}

func (postgresqlConnectionFactory *PostgresqlConnectionFactory) NewConnection() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresqlConnectionFactory.databaseConfig.Host,
		postgresqlConnectionFactory.databaseConfig.Port,
		postgresqlConnectionFactory.databaseConfig.User,
		postgresqlConnectionFactory.databaseConfig.Password,
		postgresqlConnectionFactory.databaseConfig.DBName,
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
