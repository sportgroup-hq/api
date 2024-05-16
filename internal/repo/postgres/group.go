package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (p *Postgres) GetOwnerOrJoinedGroupsByUserID(ctx context.Context, userID uuid.UUID) ([]models.Group, error) {
	var groups []models.Group

	err := p.db.NewSelect().
		Model(&groups).
		Join(`INNER JOIN group_members gm ON gm.group_id = "group".id`).
		Where("gm.user_id = ?", userID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return groups, nil
}
