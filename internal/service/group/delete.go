package group

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) LeaveGroup(ctx context.Context, userID, groupID uuid.UUID) error {
	groupMember, err := s.repo.GetGroupMember(ctx, userID, groupID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.ErrNotJoined
		}

		return fmt.Errorf("failed to check if group member exists: %w", err)
	}

	if groupMember.Type == models.GroupMemberTypeOwner {
		return models.ErrOwnerCannotLeave
	}

	if err = s.repo.DeleteGroupMember(ctx, userID, groupID); err != nil {
		return fmt.Errorf("failed to delete group member: %w", err)
	}

	return nil
}

func (s *Service) DeleteGroup(ctx context.Context, initiatorID uuid.UUID, groupID uuid.UUID) error {
	groupMember, err := s.repo.GetGroupMember(ctx, initiatorID, groupID)
	if err != nil {
		return fmt.Errorf("failed to get group member: %w", err)
	}

	if groupMember.Type != models.GroupMemberTypeOwner {
		return models.ErrNotOwner
	}

	return s.repo.DeleteGroup(ctx, groupID)
}
