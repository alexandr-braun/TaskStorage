package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"taskStorage/pkg/domain"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(connStr string) (*UserRepository, error) {
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return &UserRepository{db: db}, nil
}

func (userRepository *UserRepository) Close() error {
	err := userRepository.db.Close(context.Background())
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	return nil
}

func (userRepository *UserRepository) GetUser(id string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}
