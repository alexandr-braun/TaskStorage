package grpc

import (
	"taskStorage/pkg/domain"
)

// ToGrpcResponse TODO extract to separate converters catalog
func ToGrpcResponse(user domain.User) *GetUserResponse {
	userGrpcModel := &User{
		Id:         user.Id,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
	}

	grpcResponse := &GetUserResponse{
		User: userGrpcModel,
	}

	return grpcResponse
}
