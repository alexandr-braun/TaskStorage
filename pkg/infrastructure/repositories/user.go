package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"taskStorage/pkg/domain"
)

type PostgreSqlUserRepository struct {
	db *pgx.Conn
}

func NewPostgreSqlUserRepository(connStr string) (*PostgreSqlUserRepository, error) {
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return &PostgreSqlUserRepository{db: db}, nil
}

func (userRepository *PostgreSqlUserRepository) Close() error {
	err := userRepository.db.Close(context.Background())
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	return nil
}

func (userRepository *PostgreSqlUserRepository) GetUser(id string) (*domain.User, error) {
	var user domain.User

	err := userRepository.db.QueryRow(context.Background(),
		"SELECT id, first_name, middle_name, last_name"+
			"FROM users"+
			"WHERE id = $1", id).Scan(&user.Id, &user.FirstName, &user.MiddleName, &user.LastName)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
