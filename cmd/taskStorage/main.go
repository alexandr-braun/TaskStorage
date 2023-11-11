package main

import (
	"log"
	"taskStorage/pkg/infrastructure"
	"taskStorage/pkg/presentation/grpc"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	infrastructure.NewDatabaseRegistrar().ConnectToDatabase(cfg.Database)
	grpc.NewGrpcRegistrar().RegisterGrpcServices()

	select {}
}
