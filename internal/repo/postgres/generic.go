package postgres

import (
	"context"

	"github.com/uptrace/bun"
)

func (p *Postgres) insert(ctx context.Context, entity any) error {
	_, err := p.tx().NewInsert().
		Model(entity).
		Exec(ctx)

	return p.err(err)
}

func (p *Postgres) updateByPK(ctx context.Context, entity any) error {
	_, err := p.tx().NewUpdate().
		Model(entity).
		WherePK().
		Exec(ctx)

	return p.err(err)
}

func (p *Postgres) deleteByPK(ctx context.Context, obj any) error {
	_, err := p.tx().NewDelete().
		Model(obj).
		WherePK().
		Exec(ctx)

	return p.err(err)
}

func (p *Postgres) upsertByID(ctx context.Context, obj any, onConflictOf string) error {
	_, err := p.tx().NewInsert().
		On("conflict (?) do update", bun.Safe(onConflictOf)).
		Model(obj).
		Exec(ctx)

	return p.err(err)
}
