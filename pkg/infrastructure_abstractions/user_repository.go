package infrastructure_abstractions

import "taskStorage/pkg/domain"

type UserRepository interface {
	GetUser(id int32) (*domain.User, error)
}
