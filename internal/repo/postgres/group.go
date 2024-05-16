package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (p *Postgres) GetOwnerOrJoinedGroupsByUserID(ctx context.Context, userID uuid.UUID) ([]models.Group, error) {
	var groups []models.Group

	err := p.tx(ctx).NewSelect().
		Model(&groups).
		Join(`INNER JOIN group_members gm ON gm.group_id = "group".id`).
		Where("gm.user_id = ?", userID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return groups, nil
}

func (p *Postgres) CreateGroup(ctx context.Context, group *models.Group) (*models.Group, error) {
	if err := p.insert(ctx, group); err != nil {
		return nil, err
	}

	return group, nil
}

func (p *Postgres) CreateGroupMember(ctx context.Context, groupID, userID uuid.UUID, memberType models.GroupMemberType) (*models.GroupMember, error) {
	groupMember := &models.GroupMember{
		GroupID: groupID,
		UserID:  userID,
		Type:    memberType,
	}

	if err := p.insert(ctx, groupMember); err != nil {
		return nil, err
	}

	return groupMember, nil
}
