package event

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/api/internal/repo"
)

type Repo interface {
	repo.Atomic

	CreateEvent(ctx context.Context, event *models.Event) error
	CreateEventAssignees(ctx context.Context, assignees []models.EventAssignee) error

	GetEventsByGroup(ctx context.Context, groupID uuid.UUID, opts ...repo.Option) (models.Events, error)
	GetEventByIDAndGroupID(ctx context.Context, eventID, groupID uuid.UUID) (*models.Event, error)
	EventWithGroupIDExists(ctx context.Context, groupID, eventID uuid.UUID) (bool, error)

	DeleteEvent(ctx context.Context, eventID uuid.UUID) error

	UpsertEventRecordValue(ctx context.Context, recordValue *models.EventRecordValue) error

	GetEventValueByGroupIDAndEventID(ctx context.Context, groupID uuid.UUID, eventID ...uuid.UUID) ([]models.EventRecordValue, error)

	CreateEventComment(ctx context.Context, comment *models.EventComment) error
	GetEventComments(ctx context.Context, eventID uuid.UUID) ([]models.EventComment, error)

	OrRecordAssignTypeAllOrSelected(userID uuid.UUID) repo.Option
}

type GroupRepo interface {
	GetGroupMember(ctx context.Context, userID, groupID uuid.UUID) (*models.GroupMember, error)
	GroupMemberExists(ctx context.Context, userID, groupID uuid.UUID) (bool, error)
}

type Service struct {
	cfg       *config.Config
	repo      Repo
	groupRepo GroupRepo
}

func New(cfg *config.Config, repo Repo, groupRepo GroupRepo) *Service {
	return &Service{
		cfg:       cfg,
		repo:      repo,
		groupRepo: groupRepo,
	}
}
