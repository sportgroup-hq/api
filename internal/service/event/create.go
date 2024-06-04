package event

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/api/internal/repo"
)

func (s *Service) CreateEvent(ctx context.Context, userID, groupID uuid.UUID, cer *models.CreateEventRequest) (*models.Event, error) {
	groupMember, err := s.groupRepo.GetGroupMember(ctx, userID, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group member: %w", err)
	}

	if groupMember.CanCreateEvent() {
		return nil, models.ErrForbidden
	}

	event := cer.ToEvent()

	event.ID = uuid.New()
	event.GroupID = groupID
	event.CreatedBy = userID

	event.Records = cer.Records.ToEventRecords(event.ID)

	if event.Records.ContainsNotUniqueTitle() {
		return nil, models.ErrRecordTitleNotUnique
	}

	err = s.repo.Atomic(ctx, func(atomicRepo repo.Atomic) error {
		r := atomicRepo.(Repo)

		if err = r.CreateEvent(ctx, event); err != nil {
			return fmt.Errorf("failed to create event: %w", err)
		}

		if cer.AssignType == models.AssignTypeSelected {
			if len(cer.AssignedUserIDs) == 0 {
				return models.ErrAssigneesRequired
			}

			var assignees []models.EventAssignee

			for _, assignedUserID := range cer.AssignedUserIDs {
				assignees = append(assignees, models.EventAssignee{
					EventID: event.ID,
					UserID:  assignedUserID,
				})
			}

			if err = r.CreateEventAssignees(ctx, assignees); err != nil {
				return fmt.Errorf("failed to create event assignees: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *Service) UpdateEvent(ctx context.Context, userID, groupID, eventID uuid.UUID, uer *models.UpdateEventRequest) (*models.Event, error) {
	return nil, nil
}

func (s *Service) DeleteEvent(ctx context.Context, userID, groupID, eventID uuid.UUID) error {
	groupMember, err := s.groupRepo.GetGroupMember(ctx, userID, groupID)
	if err != nil {
		return fmt.Errorf("failed to get group member: %w", err)
	}

	if !groupMember.CanDeleteEvent() {
		return models.ErrForbidden
	}

	// check if event for group exists
	if _, err = s.repo.GetEventByIDAndGroupID(ctx, eventID, groupID); err != nil {
		return fmt.Errorf("failed to get event: %w", err)
	}

	if err = s.repo.DeleteEvent(ctx, eventID); err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}

	return nil
}

func (s *Service) SetEventRecordValue(ctx context.Context, userID, groupID, eventID, recordID uuid.UUID, value *models.EventRecordValue) error {
	groupMember, err := s.groupRepo.GetGroupMember(ctx, userID, groupID)
	if err != nil {
		return fmt.Errorf("failed to get group member: %w", err)
	}

	event, err := s.repo.GetEventByIDAndGroupID(ctx, eventID, groupID)
	if err != nil {
		return fmt.Errorf("failed to get event: %w", err)
	}

	record := event.GetRecord(recordID)

	if record == nil {
		return models.ErrRecordNotFound
	}

	if !record.WriteAccessScopes.AllowedForMemberType(groupMember.Type) {
		return models.ErrForbidden
	}

	value.UserID = userID
	value.EventID = eventID
	value.RecordID = recordID

	if err = s.repo.UpsertEventRecordValue(ctx, value); err != nil {
		return fmt.Errorf("failed to upsert event record value: %w", err)
	}

	return nil
}
