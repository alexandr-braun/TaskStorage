package user

import (
	"context"
	"taskStorage/pkg/application/queries/get_user_info"
	"taskStorage/pkg/presentation/grpc/services/user/converters"
	pb "taskStorage/pkg/presentation/grpc/services/user/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var getUserInfoQuery = get_user_info.GetUserInfoQuery{Id: req.Id}
	var getUserInfoQueryHandler = get_user_info.GetUserInfoQueryHandler{}

	var userInfo = getUserInfoQueryHandler.Handle(getUserInfoQuery)
	var result = converters.ToGrpcResponse(userInfo)
	return result, nil
}
