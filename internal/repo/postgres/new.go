package postgres

import (
	"database/sql"
	"fmt"

	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Postgres struct {
	db          *bun.DB
	transaction *bun.Tx
}

func New(cfg config.Postgres) *Postgres {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDB, pgdialect.New())
	if cfg.Log {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	db.RegisterModel((*models.EventAssignee)(nil))

	return &Postgres{
		db: db,
	}
}

func (p *Postgres) clone() *Postgres {
	return &Postgres{
		db: p.db,
	}
}

func (p *Postgres) withTx(tx bun.Tx) *Postgres {
	return &Postgres{
		transaction: &tx,
	}
}
