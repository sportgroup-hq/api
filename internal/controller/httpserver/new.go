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
	GetGroupsByUser(ctx context.Context, userID uuid.UUID) ([]models.Group, error)
	GetGroupByID(ctx context.Context, userID, groupID uuid.UUID) (*models.Group, error)
	GetGroupMembers(ctx context.Context, userID, groupID uuid.UUID) ([]models.User, error)

	CreateGroup(ctx context.Context, creatorID uuid.UUID, group *models.Group) (*models.Group, error)
	JoinGroup(ctx context.Context, userID uuid.UUID, code string) error

	UpdateGroup(ctx context.Context, userID uuid.UUID, group models.UpdateGroupRequest) (*models.Group, error)

	LeaveGroup(ctx context.Context, userID, groupID uuid.UUID) error
	DeleteGroup(ctx context.Context, initiatorID, groupID uuid.UUID) error

	GetGroupRecords(ctx context.Context, userID, groupID uuid.UUID) ([]models.GroupRecord, error)
}

type EventService interface {
	GetEventsByGroup(ctx context.Context, userID, groupID uuid.UUID) ([]models.Event, error)
	GetEventByID(ctx context.Context, userID, groupID, eventID uuid.UUID) (*models.Event, error)
	CreateEvent(ctx context.Context, userID, groupID uuid.UUID, cer *models.CreateEventRequest) (*models.Event, error)
	UpdateEvent(ctx context.Context, userID, groupID, eventID uuid.UUID, uer *models.UpdateEventRequest) (*models.Event, error)
	DeleteEvent(ctx context.Context, userID, groupID, eventID uuid.UUID) error

	SetEventRecordValue(ctx context.Context, userID, groupID, eventID, recordID uuid.UUID, value *models.EventRecordValue) error
}

type Server struct {
	cfg      *config.Config
	userSrv  UserService
	groupSrv GroupService
	eventSrv EventService
}

func New(cfg *config.Config, userSrv UserService, groupSrv GroupService, eventSrv EventService) *Server {
	return &Server{
		cfg:      cfg,
		userSrv:  userSrv,
		groupSrv: groupSrv,
		eventSrv: eventSrv,
	}
}
