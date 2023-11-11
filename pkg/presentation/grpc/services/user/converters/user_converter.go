package converters

import (
	"taskStorage/pkg/domain"
	pb "taskStorage/pkg/presentation/grpc/services/user/grpc"
)

func ToGrpcResponse(user domain.User) *pb.GetUserResponse {
	userGrpcModel := &pb.User{
		Id:         user.Id,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
	}

	grpcResponse := &pb.GetUserResponse{
		User: userGrpcModel,
	}

	return grpcResponse
}
