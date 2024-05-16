package group

import "github.com/sportgroup-hq/api/internal/config"

type Repo interface {
}

type Service struct {
	cfg  *config.Config
	repo Repo
}

func New(cfg *config.Config, repo Repo) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}
