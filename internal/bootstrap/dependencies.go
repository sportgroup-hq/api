package bootstrap

import (
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/repo"
	"github.com/sportgroup-hq/api/internal/service"
	"github.com/sportgroup-hq/api/internal/service/grpcserver"
	"github.com/sportgroup-hq/api/internal/service/httpserver"
)

type Dependencies struct {
	Config config.Config

	HTTPServer *httpserver.Server
	GRPCServer *grpcserver.Server

	AuthService service.Auth
	UserService service.User

	repo repo.Repo
}

func NewDependencies(config config.Config, httpServer *httpserver.Server,
	grpcServer *grpcserver.Server, authSrv service.Auth,
	userSrv service.User, repo repo.Repo) *Dependencies {
	return &Dependencies{
		Config:      config,
		HTTPServer:  httpServer,
		GRPCServer:  grpcServer,
		AuthService: authSrv,
		UserService: userSrv,
		repo:        repo,
	}
}
