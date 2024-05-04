package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s Service) GetUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return &models.User{
		ID: userID,
	}, nil

	user, err := s.repo.User().GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
