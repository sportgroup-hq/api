package postgres

import (
	"database/sql"
	"fmt"

	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/repo"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Postgres struct {
	db *bun.DB
}

func New(cfg config.Postgres) *Postgres {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDB, pgdialect.New())
	if cfg.Log {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Auth() repo.Auth {
	return p
}

func (p *Postgres) User() repo.User {
	return p
}
