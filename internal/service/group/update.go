package group

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) UpdateGroup(ctx context.Context, userID uuid.UUID, updateGroupRequest models.UpdateGroupRequest) (*models.Group, error) {
	groupMember, err := s.repo.GetGroupMember(ctx, userID, updateGroupRequest.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group member: %w", err)
	}

	if !groupMember.CanEditGroup() {
		return nil, models.ErrForbidden
	}

	group, err := s.repo.GetGroup(ctx, updateGroupRequest.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group: %w", err)
	}

	newGroup := updateGroupRequest.Apply(group)

	if err = s.repo.UpdateGroup(ctx, newGroup); err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	return group, nil
}
