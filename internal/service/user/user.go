package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

func (s *Service) UserExistsByID(ctx context.Context, userID uuid.UUID) (bool, error) {
	panic("not implemented")
}

func (s *Service) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	return s.repo.UserExistsByEmail(ctx, email)
}

func (s *Service) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return user, nil
}
