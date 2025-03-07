package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type reides struct {
	redis *redis.Client
}
type UserRedies interface {
	AddTokenToCache(ctx context.Context, email string, token string) error
	GetTokenFromCache(ctx context.Context, id string) (string, error)
}

func NewRediers(redies *redis.Client) (*reides, error) {
	return &reides{
		redis: redies,
	}, nil
}
