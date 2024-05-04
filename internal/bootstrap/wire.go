//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/repo"
	"github.com/sportgroup-hq/api/internal/repo/postgres"
	"github.com/sportgroup-hq/api/internal/service"
	"github.com/sportgroup-hq/api/internal/service/auth"
	"github.com/sportgroup-hq/api/internal/service/grpcserver"
	"github.com/sportgroup-hq/api/internal/service/httpserver"
	"github.com/sportgroup-hq/api/internal/service/user"
)

func Up() (*Dependencies, error) {
	wire.Build(
		wire.Bind(new(service.User), new(*user.Service)),
		wire.Bind(new(service.Auth), new(*auth.Service)),

		wire.Bind(new(repo.Repo), new(*postgres.Postgres)),

		config.New,
		httpserver.New,
		grpcserver.New,
		getPostgresConfig,

		postgres.New,
		auth.New,
		user.New,
		NewDependencies,
	)
	return &Dependencies{}, nil
}

func getPostgresConfig(cfg config.Config) config.Postgres {
	return cfg.Postgres
}
