package group

import (
	"context"
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

	if err = s.repo.CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return group, nil
}
