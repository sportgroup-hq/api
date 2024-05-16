package group

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) GetUserGroups(ctx context.Context, userID uuid.UUID) ([]models.Group, error) {
	groups, err := s.repo.GetOwnerOrJoinedGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user groups: %w", err)
	}

	return groups, nil
}
