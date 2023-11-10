package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddUsersTable, downAddUsersTable)
}

func upAddUsersTable(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(255),
        middle_name VARCHAR(255),
        last_name VARCHAR(255)
    )`)
	return err
}

func downAddUsersTable(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS users")
	return err
}
