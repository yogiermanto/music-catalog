package memberships

import (
	"context"
	"music-catalog/internal/configs"
	"music-catalog/internal/models/memberships"
)

type repository interface {
	CreateUser(ctx context.Context, model *memberships.User) error
	GetUser(ctx context.Context, email, username string, id int) (resp memberships.User, err error)
}

type Service struct {
	cfg  *configs.Config
	repo repository
}

func NewService(cfg *configs.Config, repo repository) *Service {
	return &Service{
		cfg,
		repo,
	}
}
