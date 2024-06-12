package postgres

import (
	"context"
	"fmt"

	"github.com/sportgroup-hq/api/internal/repo"
)

func (p *Postgres) Atomic(ctx context.Context, f func(tx repo.Atomic) error) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	if err = f(p.withTx(tx)); err == nil {
		return tx.Commit()
	}

	if err = tx.Rollback(); err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}

	return fmt.Errorf("failed to execute atomic function: %w", err)
}
