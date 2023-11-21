package grpc

import (
	"context"
	"taskStorage/pkg/application/queries/get_user_info"
)

type Server struct {
	UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	//TODO implement it
	return nil, nil
}

func (s *Server) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	var getUserInfoQuery = get_user_info.GetUserInfoQuery{Id: req.Id}
	var getUserInfoQueryHandler = get_user_info.GetUserInfoQueryHandler{}

	var userInfo = getUserInfoQueryHandler.Handle(getUserInfoQuery)
	var result = ToGrpcResponse(*userInfo)
	return result, nil
}
