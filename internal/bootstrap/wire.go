//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/repo/postgres"
	"github.com/sportgroup-hq/api/internal/service/group"
	"github.com/sportgroup-hq/api/internal/service/grpcserver"
	"github.com/sportgroup-hq/api/internal/service/httpserver"
	"github.com/sportgroup-hq/api/internal/service/user"
)

func Up() (*Dependencies, error) {
	wire.Build(
		config.New,

		getPostgresConfig,
		postgres.New,

		user.New,
		wire.Bind(new(user.Repo), new(*postgres.Postgres)),

		group.New,
		wire.Bind(new(group.Repo), new(*postgres.Postgres)),

		httpserver.New,
		wire.Bind(new(httpserver.UserService), new(*user.Service)),
		wire.Bind(new(httpserver.GroupService), new(*group.Service)),

		grpcserver.New,
		wire.Bind(new(grpcserver.UserService), new(*user.Service)),

		NewDependencies,
	)
	return &Dependencies{}, nil
}

func getPostgresConfig(cfg *config.Config) config.Postgres {
	return cfg.Postgres
}
