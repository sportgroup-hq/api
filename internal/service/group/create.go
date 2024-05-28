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

func (s *Service) UpdateGroup(ctx context.Context, userID uuid.UUID, updateGroupRequest models.UpdateGroupRequest) (*models.Group, error) {
	groupMember, err := s.repo.GetGroupMember(ctx, userID, updateGroupRequest.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group member: %w", err)
	}

	if groupMember.CanEditGroup() {
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
