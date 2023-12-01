package infrastructure_abstractions

import "database/sql"

type DBConnectionFactory interface {
	NewConnection() (*sql.DB, error)
}
