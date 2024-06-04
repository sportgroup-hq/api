package group

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) GetGroupsByUser(ctx context.Context, userID uuid.UUID) ([]models.Group, error) {
	groups, err := s.repo.GetGroupsJoinedByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user groups: %w", err)
	}

	return groups, nil
}

func (s *Service) GetGroupByID(ctx context.Context, userID, groupID uuid.UUID) (*models.Group, error) {
	isMember, err := s.repo.GroupMemberExists(ctx, userID, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user is a member of the group: %w", err)
	}

	if !isMember {
		return nil, models.ErrNotFound
	}

	group, err := s.repo.GetGroup(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group by id: %w", err)
	}

	return group, nil
}

func (s *Service) GetGroupMembers(ctx context.Context, userID, groupID uuid.UUID) ([]models.User, error) {
	isMember, err := s.repo.GroupMemberExists(ctx, userID, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user is a member of the group: %w", err)
	}

	if !isMember {
		return nil, models.ErrNotFound
	}

	members, err := s.userRepo.GetUserByGroupMembership(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group members: %w", err)
	}

	return members, nil
}
