package postgres

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/sportgroup-hq/api/internal/models"
	"github.com/sportgroup-hq/api/internal/repo"
	"github.com/uptrace/bun"
)

func (p *Postgres) tx() bun.IDB {
	if p.transaction != nil {
		return p.transaction
	}

	return p.db
}

func (p *Postgres) err(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return models.ErrNotFound
	default:
		return err
	}
}

func toSelectOptions(opts []repo.Option) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(sq *bun.SelectQuery) *bun.SelectQuery {
		for _, opt := range opts {
			optTyped, ok := opt.(func(sq *bun.SelectQuery) *bun.SelectQuery)
			if !ok {
				slog.Error("invalid option type: %T", opt)
				continue
			}
			sq = optTyped(sq)
		}
		return sq
	}
}
