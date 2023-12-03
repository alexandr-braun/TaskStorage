package main

import (
	"context"
	"go.uber.org/fx"
	"taskStorage/pkg/application/queries/get_user_info"
	"taskStorage/pkg/configuration"
	"taskStorage/pkg/infrastructure/postgresql"
	_ "taskStorage/pkg/infrastructure/postgresql/migrations"
	"taskStorage/pkg/infrastructure/postgresql/repositories"
	"taskStorage/pkg/infrastructure_abstractions"
	"taskStorage/pkg/presentation/grpc"
)

func main() {
	app := fx.New(
		fx.Provide(
			configuration.NewConfig,
			postgresql.NewPostgresqlConnectionFactory,
			repositories.NewPostgreSqlUserRepository,
			get_user_info.NewGetUserInfoQueryHandler,
		),
		fx.Invoke(
			func(lifecycle fx.Lifecycle, dbFactory infrastructure_abstractions.DBConnectionFactory) {
				lifecycle.Append(fx.Hook{
					OnStart: func(context.Context) error {
						return postgresql.RunMigrations(dbFactory)
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
