package main

import (
	"go.uber.org/fx"
	"log"
	"taskStorage/pkg/infrastructure"
	"taskStorage/pkg/presentation/grpc"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	app := fx.New(
		fx.Provide(infrastructure.DatabaseConfig{Host: cfg.Database.Host,
			Port:     cfg.Database.Port,
			User:     cfg.Database.User,
			Password: cfg.Database.Password,
			DBName:   cfg.Database.DBName,
		}),
	)

	app.Run()

	infrastructure.NewDatabaseConnection().Connect(cfg.Database)
	grpc.NewGrpcRegistrar().RegisterGrpcServices()

	select {}
}
