package auth

import (
	"github.com/sportgroup-hq/api/internal/config"
	"github.com/sportgroup-hq/api/internal/repo"
)

type Service struct {
	cfg config.Config

	repo repo.Repo
}

func New(cfg config.Config, repo repo.Repo) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}
