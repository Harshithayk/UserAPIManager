package dbconnection

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func ConnectRedis() (*redis.Client,error) {
	rdb := redis.NewClient(&redis.Options{
		Password: "",
		DB:       0,
		Addr:     "localhost:6379",
	})
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Error().Err(err).Msg("Connection to the redis is closed")
		return nil, err
	}
	return rdb, nil
}
