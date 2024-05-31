package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

type Repo interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, request *models.User) error
}

type Service struct {
	repo Repo
}

func New(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}
