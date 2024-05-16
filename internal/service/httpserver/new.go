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

type GroupService interface {
	GetUserGroups(ctx context.Context, userID uuid.UUID) ([]models.Group, error)
}

type Server struct {
	cfg      *config.Config
	userSrv  UserService
	groupSrv GroupService
}

func New(cfg *config.Config, userSrv UserService, groupSrv GroupService) *Server {
	return &Server{
		cfg:      cfg,
		userSrv:  userSrv,
		groupSrv: groupSrv,
	}
}
