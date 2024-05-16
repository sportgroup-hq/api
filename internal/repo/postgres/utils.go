package postgres

import (
	"database/sql"
	"errors"

	"github.com/sportgroup-hq/api/internal/models"
)

func (p *Postgres) err(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return models.NotFoundError
	}

	return err
}
