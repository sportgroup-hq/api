package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

type Auth interface {
	GenerateToken(ctx context.Context, user *models.User) (*models.UserCredentials, error)
	Login(ctx context.Context, accessToken string) (*models.User, error)
}

type User interface {
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	//GetByEmail(ctx context.Context, email string) (*models.User, error)
	//UpdateUser(ctx context.Context, user *models.User) error
}
