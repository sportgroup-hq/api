package grpcserver

import (
	"log"
	"log/slog"
	"net"

	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/service"
	"github.com/sportgroup-hq/common-lib/api"
	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedApiServer

	UserSrv service.User
}

func New(userSrv service.User) *Server {
	return &Server{UserSrv: userSrv}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", config.Get().GRPC.ApiAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterApiServer(grpcServer, s)

	slog.Info("Starting GRPC server on " + config.Get().GRPC.ApiAddress + "...")

	return grpcServer.Serve(lis)
}
