package main

import (
	"context"
	"go.uber.org/fx"
	"log"
	"taskStorage/pkg/application/queries/get_user_info"
	"taskStorage/pkg/infrastructure"
	"taskStorage/pkg/infrastructure/repositories"
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
		},
			infrastructure.NewPostgresqlConnectionFactory,
			repositories.NewPostgreSqlUserRepository,
			get_user_info.NewGetUserInfoQueryHandler,
		),
		fx.Invoke(
			func(lifecycle fx.Lifecycle, dbFactory infrastructure.DBConnectionFactory) {
				lifecycle.Append(fx.Hook{
					OnStart: func(context.Context) error {
						return infrastructure.RunMigrations(dbFactory)
					},
				})
				lifecycle.Append(fx.Hook{
					OnStart: func(context.Context) error {
						return grpc.NewGrpcRegistrar().RegisterGrpcServices()
					},
				})
			},
		),
	)

	app.Run()
}
