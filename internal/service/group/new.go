package group

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/api/internal/repo"
)

type Repo interface {
	repo.Atomic

	//BeginTx(ctx context.Context) (context.Context, error)
	//CommitTx(ctx context.Context) error
	//RollbackTx(ctx context.Context) error

	GetGroup(ctx context.Context, groupID uuid.UUID) (*models.Group, error)
	GetGroupInviteByCode(ctx context.Context, code string) (*models.GroupInvite, error)
	GetGroupsJoinedByUserID(ctx context.Context, userID uuid.UUID) ([]models.Group, error)
	CreateGroup(ctx context.Context, group *models.Group) (*models.Group, error)
	UpdateGroup(ctx context.Context, group *models.Group) error
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error
	CreateGroupInvite(ctx context.Context, groupID uuid.UUID, code string) (*models.GroupInvite, error)

	CreateGroupMember(ctx context.Context, groupID, userID uuid.UUID, memberType models.GroupMemberType) (*models.GroupMember, error)
	GetGroupMember(ctx context.Context, userID, groupID uuid.UUID) (*models.GroupMember, error)
	GroupMemberExists(ctx context.Context, userID, groupID uuid.UUID) (bool, error)

	DeleteGroupMember(ctx context.Context, userID, groupID uuid.UUID) error

	GetGroupRecords(ctx context.Context, groupID uuid.UUID) ([]models.GroupRecord, error)
	CopyDefaultGroupRecords(ctx context.Context, groupID uuid.UUID) error
}

type UserRepo interface {
	GetUserByGroupMembership(ctx context.Context, groupID uuid.UUID) ([]models.User, error)
}

type Service struct {
	cfg      *config.Config
	repo     Repo
	userRepo UserRepo
}

func New(cfg *config.Config, repo Repo, userRepo UserRepo) *Service {
	return &Service{
		cfg:      cfg,
		repo:     repo,
		userRepo: userRepo,
	}
}
