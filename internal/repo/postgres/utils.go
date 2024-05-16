package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sportgroup-hq/api/internal/models"
	"github.com/uptrace/bun"
)

func (p *Postgres) tx(ctx context.Context) bun.IDB {
	tx, ok := ctx.Value(txKey).(bun.Tx)
	if ok {
		return tx
	}

	return p.db
}

func (p *Postgres) err(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return models.ErrNotFound
	default:
		return err
	}
}
