package infrastructure_abstractions

import "taskStorage/pkg/domain"

type UserRepository interface {
	GetUser(id string) (*domain.User, error)
}
