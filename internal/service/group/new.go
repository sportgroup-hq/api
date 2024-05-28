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

	GetGroup(ctx context.Context, groupID uuid.UUID) (*models.Group, error)
	GetGroupInviteByCode(ctx context.Context, code string) (*models.GroupInvite, error)
	GetOwnerOrJoinedGroupsByUserID(ctx context.Context, userID uuid.UUID) ([]models.Group, error)
	CreateGroup(ctx context.Context, group *models.Group) (*models.Group, error)
	UpdateGroup(ctx context.Context, group *models.Group) error
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error
	CreateGroupInvite(ctx context.Context, groupID uuid.UUID, code string) (*models.GroupInvite, error)

	CreateGroupMember(ctx context.Context, groupID, userID uuid.UUID, memberType models.GroupMemberType) (*models.GroupMember, error)
	GetGroupMember(ctx context.Context, userID, groupID uuid.UUID) (*models.GroupMember, error)
	GroupMemberExists(ctx context.Context, userID, groupID uuid.UUID) (bool, error)

	DeleteGroupMember(ctx context.Context, userID, groupID uuid.UUID) error
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
