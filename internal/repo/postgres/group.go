package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (p *Postgres) CreateGroup(ctx context.Context, group *models.Group) (*models.Group, error) {
	if err := p.insert(ctx, group); err != nil {
		return nil, err
	}

	return group, nil
}

func (p *Postgres) CreateGroupInvite(ctx context.Context, groupID uuid.UUID, code string) (*models.GroupInvite, error) {
	groupInvite := &models.GroupInvite{
		GroupID: groupID,
		Code:    code,
	}

	if err := p.insert(ctx, groupInvite); err != nil {
		return nil, err
	}

	return groupInvite, nil
}

func (p *Postgres) GetGroup(ctx context.Context, groupID uuid.UUID) (*models.Group, error) {
	group := new(models.Group)

	err := p.tx(ctx).NewSelect().
		Model(group).
		Where("id = ?", groupID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return group, nil
}

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

func (p *Postgres) GetGroupInviteByCode(ctx context.Context, code string) (*models.GroupInvite, error) {
	var groupInvite models.GroupInvite

	err := p.tx(ctx).NewSelect().
		Model(&groupInvite).
		Where("code = ?", code).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &groupInvite, nil
}

func (p *Postgres) UpdateGroup(ctx context.Context, group *models.Group) error {
	return p.updateByPK(ctx, group)
}

func (p *Postgres) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	_, err := p.tx(ctx).NewDelete().
		Model((*models.Group)(nil)).
		Where("id = ?", groupID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) GetGroupMember(ctx context.Context, userID, groupID uuid.UUID) (*models.GroupMember, error) {
	var groupMember models.GroupMember

	err := p.tx(ctx).NewSelect().
		Model(&groupMember).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &groupMember, nil
}

func (p *Postgres) GroupMemberExists(ctx context.Context, userID, groupID uuid.UUID) (bool, error) {
	exists, err := p.tx(ctx).NewSelect().
		Model((*models.GroupMember)(nil)).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		Exists(ctx)
	if err != nil {
		return false, p.err(err)
	}

	return exists, nil
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

func (p *Postgres) DeleteGroupMember(ctx context.Context, userID, groupID uuid.UUID) error {
	_, err := p.tx(ctx).NewDelete().
		Model((*models.GroupMember)(nil)).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
