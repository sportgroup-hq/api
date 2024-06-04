package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/api/internal/models"
)

func (p *Postgres) GetGroupRecords(ctx context.Context, groupID uuid.UUID) ([]models.GroupRecord, error) {
	var records []models.GroupRecord

	err := p.tx().NewSelect().
		Model(&records).
		Where("group_id = ?", groupID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return records, nil
}

func (p *Postgres) CopyDefaultGroupRecords(ctx context.Context, groupID uuid.UUID) error {
	_, err := p.tx().QueryContext(ctx, `
INSERT INTO group_records (group_id, title, type, read_access_scopes, write_access_scopes)
SELECT ?, title, type, read_access_scopes, write_access_scopes FROM group_records_default;`, groupID)
	return p.err(err)
}
