package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (p *Postgres) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	err := p.tx(ctx).NewSelect().
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

	err := p.tx(ctx).NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &user, nil
}

func (p *Postgres) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := p.tx(ctx).NewSelect().
		Model((*models.User)(nil)).
		Where("email = ?", email).
		Exists(ctx)
	if err != nil {
		return false, p.err(err)
	}

	return exists, nil
}

func (p *Postgres) CreateUser(ctx context.Context, user *models.User) error {
	if err := p.insert(ctx, user); err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateUser(ctx context.Context, user *models.User) error {
	return p.updateByPK(ctx, user)
}
