package grpc

import (
	"taskStorage/pkg/domain"
)

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
