package user

import "github.com/sportgroup-hq/api/internal/repo"

type Service struct {
	repo repo.Repo
}

func New(repo repo.Repo) *Service {
	return &Service{
		repo: repo,
	}
}
