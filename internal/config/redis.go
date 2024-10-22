package config

import (
	"context"
	"fmt"
	"log"

	"github.com/gavril-s/borzoi/internal/pkg/storage"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis(ctx context.Context, cfg Config) *storage.Redis {
	db, err := storage.NewRedis(ctx, &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	return db
}
