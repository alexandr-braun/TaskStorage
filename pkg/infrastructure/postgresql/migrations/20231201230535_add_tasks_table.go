package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddTasksTable, downAddTasksTable)
}

func upAddTasksTable(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE tasks (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            description TEXT,
            priority INT NOT NULL,
            status INT NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            deadline TIMESTAMP WITH TIME ZONE,
            assigned_to_user_id INT,
            categories TEXT[],
            FOREIGN KEY (assigned_to_user_id) REFERENCES users(id)
        )`)
	return err
}

func downAddTasksTable(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS tasks")
	return err
}
