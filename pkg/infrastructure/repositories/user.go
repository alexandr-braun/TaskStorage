package repositories

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"taskStorage/pkg/domain"
	"taskStorage/pkg/infrastructure"
)

type PostgreSqlUserRepository struct {
	dbConnectionFactory infrastructure.DBConnectionFactory
}

func NewPostgreSqlUserRepository(dbConnectionFactory infrastructure.DBConnectionFactory) (*PostgreSqlUserRepository, error) {
	return &PostgreSqlUserRepository{dbConnectionFactory: dbConnectionFactory}, nil
}

func (userRepository *PostgreSqlUserRepository) GetUser(id int32) (*domain.User, error) {
	var user domain.User
	var dbConnection, _ = userRepository.dbConnectionFactory.NewConnection()
	defer dbConnection.Close()

	err := dbConnection.QueryRow(
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
