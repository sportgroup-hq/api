package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/api/internal/repo"
	"github.com/uptrace/bun"
)

func (p *Postgres) CreateEvent(ctx context.Context, event *models.Event) error {
	return p.insert(ctx, event)
}

func (p *Postgres) CreateEventAssignees(ctx context.Context, assignees []models.EventAssignee) error {
	return p.insert(ctx, &assignees)
}

func (p *Postgres) GetEventsByGroup(ctx context.Context, groupID uuid.UUID, opts ...repo.Option) (models.Events, error) {
	var events []models.Event

	err := p.tx().NewSelect().
		Model(&events).
		Where("group_id = ?", groupID).
		Apply(toSelectOptions(opts)).
		Relation("AssignedUsers").
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return events, nil
}

func (p *Postgres) GetEventByIDAndGroupID(ctx context.Context, eventID, groupID uuid.UUID) (*models.Event, error) {
	event := new(models.Event)

	err := p.tx().NewSelect().
		Model(event).
		Where("id = ?", eventID).
		Where("group_id = ?", groupID).
		Relation("AssignedUsers").
		Relation("Comments.User").
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return event, nil
}

func (p *Postgres) EventWithGroupIDExists(ctx context.Context, groupID, eventID uuid.UUID) (bool, error) {
	exists, err := p.tx().NewSelect().
		Model((*models.Event)(nil)).
		Where("id = ?", eventID).
		Where("group_id = ?", groupID).
		Exists(ctx)
	if err != nil {
		return false, p.err(err)
	}

	return exists, nil
}

func (p *Postgres) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	return p.deleteByPK(ctx, &models.Event{ID: eventID})
}

func (p *Postgres) UpsertEventRecordValue(ctx context.Context, recordValue *models.EventRecordValue) error {
	return p.upsertByID(ctx, recordValue, "event_id, user_id, record_id")
}

func (p *Postgres) GetEventValueByGroupIDAndEventID(ctx context.Context, groupID uuid.UUID, eventIDs ...uuid.UUID) ([]models.EventRecordValue, error) {
	var values []models.EventRecordValue

	err := p.tx().NewSelect().
		Model(&values).
		Join("LEFT JOIN events ON events.id = event_record_value.event_id").
		Where("events.group_id = ?", groupID).
		Where("event_id IN (?)", bun.In(eventIDs)).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return values, nil
}

func (p *Postgres) CreateEventComment(ctx context.Context, comment *models.EventComment) error {
	return p.insert(ctx, comment)
}

func (p *Postgres) GetEventComments(ctx context.Context, eventID uuid.UUID) ([]models.EventComment, error) {
	var comments []models.EventComment

	err := p.tx().NewSelect().
		Model(&comments).
		Where("event_id = ?", eventID).
		Relation("User").
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return comments, nil
}

func (p *Postgres) OrRecordAssignTypeAllOrSelected(userID uuid.UUID) repo.Option {
	return func(sq *bun.SelectQuery) *bun.SelectQuery {
		return sq.Where("assign_type = ?", models.AssignTypeAll).
			WhereGroup(" OR ", func(sq *bun.SelectQuery) *bun.SelectQuery {
				return sq.Where("assign_type = ?", models.AssignTypeSelected).
					Join("LEFT JOIN event_assignees ON event_assignees.event_id = event.id").
					Where("event_assignees.user_id = ?", userID)
			})
	}
}
