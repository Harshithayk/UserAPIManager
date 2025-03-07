package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (c *reides) AddTokenToCache(ctx context.Context, id string, token string) error {
	err := c.redis.Set(ctx, id, token, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *reides) GetTokenFromCache(ctx context.Context, id string) (string, error) {
	val, err := c.redis.Get(ctx, id).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}
