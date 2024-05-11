package httpserver

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/models"
)

type UserService interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type Server struct {
	cfg  *config.Config
	user UserService
}

func New(cfg *config.Config, userSrv UserService) *Server {
	return &Server{
		cfg:  cfg,
		user: userSrv,
	}
}
