//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/repo/postgres"
	"github.com/sportgroup-hq/api/internal/service/grpcserver"
	"github.com/sportgroup-hq/api/internal/service/httpserver"
	"github.com/sportgroup-hq/api/internal/service/user"
)

func Up() (*Dependencies, error) {
	wire.Build(
		config.New,
		httpserver.New,
		grpcserver.New,
		getPostgresConfig,

		postgres.New,

		user.New,
		wire.Bind(new(httpserver.UserService), new(*user.Service)),
		wire.Bind(new(grpcserver.UserService), new(*user.Service)),
		wire.Bind(new(user.Repo), new(*postgres.Postgres)),

		NewDependencies,
	)
	return &Dependencies{}, nil
}

func getPostgresConfig(cfg *config.Config) config.Postgres {
	return cfg.Postgres
}
