package postgres

import "context"

func (p *Postgres) insert(ctx context.Context, entity any) error {
	_, err := p.tx(ctx).NewInsert().
		Model(entity).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
