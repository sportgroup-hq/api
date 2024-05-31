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
	UpdateUser(c context.Context, updateUserRequest models.UpdateUserRequest) (*models.User, error)
}

type GroupService interface {
	GetUserGroups(ctx context.Context, userID uuid.UUID) ([]models.Group, error)
	CreateGroup(ctx context.Context, creatorID uuid.UUID, group *models.Group) (*models.Group, error)
	UpdateGroup(ctx context.Context, userID uuid.UUID, group models.UpdateGroupRequest) (*models.Group, error)
	JoinGroup(ctx context.Context, userID uuid.UUID, code string) error
	LeaveGroup(ctx context.Context, userID, groupID uuid.UUID) error
	DeleteGroup(ctx context.Context, initiatorID uuid.UUID, groupID uuid.UUID) error
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
