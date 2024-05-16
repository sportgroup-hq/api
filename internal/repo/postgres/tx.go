package postgres

import (
	"context"
	"errors"

	"github.com/uptrace/bun"
)

func (p *Postgres) BeginTx(ctx context.Context) (context.Context, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, p.err(err)
	}

	txCommited := false

	ctx = context.WithValue(ctx, txKey, tx)
	ctx = context.WithValue(ctx, txCommitedKey, &txCommited)
	return ctx, nil
}

func (p *Postgres) CommitTx(ctx context.Context) error {
	tx, ok := ctx.Value(txKey).(bun.Tx)
	if !ok {
		return p.err(errors.New("transaction not found in context"))
	}

	if err := tx.Commit(); err != nil {
		return p.err(err)
	}

	txCommited, ok := ctx.Value(txCommitedKey).(*bool)
	if !ok {
		return p.err(errors.New("transaction not found in context"))
	}

	*txCommited = true

	return nil
}

func (p *Postgres) RollbackTx(ctx context.Context) error {
	tx, ok := ctx.Value(txKey).(bun.Tx)
	if !ok {
		return p.err(errors.New("transaction not found in context"))
	}

	txCommited, ok := ctx.Value(txCommitedKey).(*bool)
	if !ok || *txCommited {
		// ignoring rollback if transaction is already commited
		return nil
	}

	if err := tx.Rollback(); err != nil {
		return p.err(err)
	}

	return nil
}
