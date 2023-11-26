package grpc

import (
	"taskStorage/pkg/domain/user"
)

// ToGrpcResponse TODO extract to separate converters catalog
func ToGrpcResponse(user user.User) *GetUserResponse {
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
