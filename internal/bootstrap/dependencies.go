package bootstrap

import (
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/repo/postgres"
	"github.com/sportgroup-hq/api/internal/service/grpcserver"
	"github.com/sportgroup-hq/api/internal/service/httpserver"
)

type Dependencies struct {
	Config *config.Config

	HTTPServer *httpserver.Server
	GRPCServer *grpcserver.Server

	postgres *postgres.Postgres
}

func NewDependencies(config *config.Config, httpServer *httpserver.Server,
	grpcServer *grpcserver.Server, postgres *postgres.Postgres) *Dependencies {
	return &Dependencies{
		Config:     config,
		HTTPServer: httpServer,
		GRPCServer: grpcServer,
		postgres:   postgres,
	}
}
