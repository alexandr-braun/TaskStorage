package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddCommentsTable, downAddCommentsTable)
}

func upAddCommentsTable(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    task_id INT NOT NULL,
    text TEXT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (task_id) REFERENCES tasks(id)
)`)
	return err
}

func downAddCommentsTable(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS comments")
	return err
}
