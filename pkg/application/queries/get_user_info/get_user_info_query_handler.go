package get_user_info

import (
	"taskStorage/pkg/domain"
	"taskStorage/pkg/infrastructure_abstractions"
)

type GetUserInfoQueryHandler struct {
	userRepository infrastructure_abstractions.UserRepository
}

func NewGetUserInfoQueryHandler(userRepository infrastructure_abstractions.UserRepository) *GetUserInfoQueryHandler {
	return &GetUserInfoQueryHandler{userRepository: userRepository}
}

func (h *GetUserInfoQueryHandler) Handle(query GetUserInfoQuery) *domain.User {
	var user, _ = h.userRepository.GetUser(query.Id)
	return user
}
