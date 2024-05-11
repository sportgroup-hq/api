package grpcserver

import (
	"context"
	"log"
	"log/slog"
	"net"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/common-lib/api"
	"google.golang.org/grpc"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UserExistsByID(ctx context.Context, userID uuid.UUID) (bool, error)
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}

type Server struct {
	api.UnimplementedApiServer

	cfg *config.Config

	userSrv UserService
}

func New(cfg *config.Config, userSrv UserService) *Server {
	return &Server{
		cfg:     cfg,
		userSrv: userSrv,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.cfg.GRPC.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterApiServer(grpcServer, s)

	slog.Info("Starting GRPC server on " + config.Get().GRPC.Address + "...")

	return grpcServer.Serve(lis)
}
