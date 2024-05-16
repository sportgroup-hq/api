package group

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/models"
)

type Repo interface {
	BeginTx(ctx context.Context) (context.Context, error)
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error

	GetOwnerOrJoinedGroupsByUserID(ctx context.Context, userID uuid.UUID) ([]models.Group, error)
	CreateGroup(ctx context.Context, group *models.Group) (*models.Group, error)
	CreateGroupMember(ctx context.Context, groupID, userID uuid.UUID, memberType models.GroupMemberType) (*models.GroupMember, error)
}

type Service struct {
	cfg  *config.Config
	repo Repo
}

func New(cfg *config.Config, repo Repo) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}
