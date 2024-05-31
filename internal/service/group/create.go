package group

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) CreateGroup(ctx context.Context, creatorID uuid.UUID, group *models.Group) (*models.Group, error) {
	ctx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer s.repo.RollbackTx(ctx)

	// create group
	group, err = s.repo.CreateGroup(ctx, group)
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	// add owner to group
	_, err = s.repo.CreateGroupMember(ctx, group.ID, creatorID, models.GroupMemberTypeOwner)
	if err != nil {
		return nil, fmt.Errorf("failed to add owner to group: %w", err)
	}

	if _, err = s.repo.CreateGroupInvite(ctx, group.ID, randGroupCode()); err != nil {
		return nil, fmt.Errorf("failed to create group invite: %w", err)
	}

	if err = s.repo.CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return group, nil
}

func (s *Service) JoinGroup(ctx context.Context, userID uuid.UUID, code string) error {
	groupInvite, err := s.repo.GetGroupInviteByCode(ctx, code)
	if err != nil {
		return fmt.Errorf("failed to get group invite by code: %w", err)
	}

	exists, err := s.repo.GroupMemberExists(ctx, userID, groupInvite.GroupID)
	if err != nil {
		return fmt.Errorf("failed to check if group member exists: %w", err)
	}

	if exists {
		return models.ErrAlreadyJoined
	}

	if !groupInvite.Active {
		return models.ErrGroupInviteInactive
	}

	_, err = s.repo.CreateGroupMember(ctx, groupInvite.GroupID, userID, models.GroupMemberTypeStudent)
	if err != nil {
		if errors.Is(err, models.ErrDuplicate) {
			return models.ErrAlreadyJoined
		}

		return fmt.Errorf("failed to create group member: %w", err)
	}

	return nil
}
