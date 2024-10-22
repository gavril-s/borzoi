package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	conn *redis.Client
}

func NewRedis(ctx context.Context, opts *redis.Options) (*Redis, error) {
	conn := redis.NewClient(opts)

	_, err := conn.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{conn: conn}, nil
}
