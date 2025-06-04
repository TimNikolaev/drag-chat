package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func New(ctx context.Context, port, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + port,
		Password: password,
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
