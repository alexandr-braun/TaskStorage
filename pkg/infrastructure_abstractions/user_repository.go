package infrastructure_abstractions

import (
	"taskStorage/pkg/domain/user"
)

type UserRepository interface {
	GetUser(id int32) (*user.User, error)
}
