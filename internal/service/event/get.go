package event

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (s *Service) GetEventsByGroup(ctx context.Context, userID, groupID uuid.UUID) ([]models.Event, error) {
	groupMember, err := s.groupRepo.GetGroupMember(ctx, userID, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group member: %w", err)
	}

	events, err := s.repo.GetEventsByGroup(ctx, groupID,
		s.repo.OrRecordAssignTypeAllOrSelected(userID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}

	if len(events) == 0 {
		return nil, nil
	}

	events.FilterRecordsByAccess(groupMember.Type)

	values, err := s.repo.GetEventValueByGroupIDAndEventID(ctx, groupID, events.IDs()...)
	if err != nil {
		return nil, fmt.Errorf("failed to get event values: %w", err)
	}

	events.AssignValues(values)

	return events, nil
}

func (s *Service) GetEventByID(ctx context.Context, userID, groupID, eventID uuid.UUID) (*models.Event, error) {
	groupMember, err := s.groupRepo.GetGroupMember(ctx, userID, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group member: %w", err)
	}

	event, err := s.repo.GetEventByIDAndGroupID(ctx, eventID, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get event: %w", err)
	}

	event.FilterRecordsByAccess(groupMember.Type)

	values, err := s.repo.GetEventValueByGroupIDAndEventID(ctx, groupID, event.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get event values: %w", err)
	}

	event.AssignValues(values)

	return event, nil
}
