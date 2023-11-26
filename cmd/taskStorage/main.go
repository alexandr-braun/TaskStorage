package main

import (
	"context"
	"go.uber.org/fx"
	"taskStorage/pkg/application/queries/get_user_info"
	"taskStorage/pkg/configuration"
	"taskStorage/pkg/infrastructure"
	"taskStorage/pkg/infrastructure/repositories"
	"taskStorage/pkg/presentation/grpc"
)

func main() {
	app := fx.New(
		fx.Provide(
			configuration.NewConfig,
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
