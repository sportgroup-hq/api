package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (p *Postgres) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	err := p.tx().NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &user, nil
}

func (p *Postgres) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := p.tx().NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &user, nil
}

func (p *Postgres) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := p.tx().NewSelect().
		Model((*models.User)(nil)).
		Where("email = ?", email).
		Exists(ctx)
	if err != nil {
		return false, p.err(err)
	}

	return exists, nil
}

func (p *Postgres) GetUserByGroupMembership(ctx context.Context, groupID uuid.UUID) ([]models.User, error) {
	var users []models.User

	err := p.tx().NewSelect().
		Model(&users).
		Join(`INNER JOIN group_members ON group_members.user_id = "user".id`).
		Where("group_members.group_id = ?", groupID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return users, nil
}

func (p *Postgres) CreateUser(ctx context.Context, user *models.User) error {
	return p.insert(ctx, user)
}

func (p *Postgres) UpdateUser(ctx context.Context, user *models.User) error {
	return p.updateByPK(ctx, user)
}
