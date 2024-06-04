package group

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) GetGroupRecords(ctx context.Context, userID, groupID uuid.UUID) ([]models.GroupRecord, error) {
	groupMember, err := s.repo.GetGroupMember(ctx, userID, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to check group membership: %w", err)
	}

	if !groupMember.CanAccessGroupRecords() {
		return nil, models.ErrForbidden
	}

	records, err := s.repo.GetGroupRecords(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group records: %w", err)
	}

	return records, nil
}
