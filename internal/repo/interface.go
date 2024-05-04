package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

type Repo interface {
	Auth() Auth
	User() User
}

type Auth interface {
	//CreateCredentials(ctx context.Context, credentials *models.UserCredentials) error
}

type User interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	//GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
