package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"taskStorage/pkg/configuration"
	"taskStorage/pkg/infrastructure_abstractions"
)

type ConnectionFactory struct {
	databaseConfig configuration.DatabaseConfig
}

func NewPostgresqlConnectionFactory(config *configuration.Config) infrastructure_abstractions.DBConnectionFactory {
	return &ConnectionFactory{
		databaseConfig: config.Database,
	}
}

func (cf *ConnectionFactory) NewConnection() (*sql.DB, error) {
	conStr := cf.getConnectionString()
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (cf *ConnectionFactory) getConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cf.databaseConfig.Host,
		cf.databaseConfig.Port,
		cf.databaseConfig.User,
		cf.databaseConfig.Password,
		cf.databaseConfig.DBName,
	)
}
