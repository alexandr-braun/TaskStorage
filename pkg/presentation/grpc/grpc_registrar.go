package grpc

import (
	"google.golang.org/grpc"
	userServiceServer "taskStorage/pkg/presentation/grpc/services/user"
	userServiceGrpc "taskStorage/pkg/presentation/grpc/services/user/grpc"
)

type Registrar struct {
}

func NewGrpcRegistrar() *Registrar {
	return &Registrar{}
}

func (registrar *Registrar) RegisterGrpcServices() error {
	newGrpcServer := grpc.NewServer()
	RegisterServices(newGrpcServer)

	return nil
}

func RegisterServices(server *grpc.Server) {
	userServiceGrpc.RegisterUserServiceServer(
		server,
		&userServiceServer.Server{})
}
