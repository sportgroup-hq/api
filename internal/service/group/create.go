package group

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/api/internal/repo"
)

func (s *Service) CreateGroup(ctx context.Context, creatorID uuid.UUID, group *models.Group) (*models.Group, error) {
	err := s.repo.Atomic(ctx, func(atomicRepo repo.Atomic) error {
		r := atomicRepo.(Repo)

		newGroup, err := r.CreateGroup(ctx, group)
		if err != nil {
			return fmt.Errorf("failed to create group: %w", err)
		}

		if _, err = r.CreateGroupInvite(ctx, newGroup.ID, randGroupCode()); err != nil {
			return fmt.Errorf("failed to create group invite: %w", err)
		}

		// add coach to group
		_, err = r.CreateGroupMember(ctx, newGroup.ID, creatorID, models.GroupMemberTypeCoach)
		if err != nil {
			return fmt.Errorf("failed to add coach to group: %w", err)
		}

		if err = r.CopyDefaultGroupRecords(ctx, newGroup.ID); err != nil {
			return fmt.Errorf("failed to copy default group records: %w", err)
		}

		group = newGroup

		return nil
	})
	if err != nil {
		return nil, err
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
