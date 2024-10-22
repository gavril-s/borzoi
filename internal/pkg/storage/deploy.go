package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) UpsertDeploy(ctx context.Context, deploy domain.Deploy) error {
	jsonData, err := json.Marshal(deploy)
	if err != nil {
		return fmt.Errorf("json.Marshal: %v", err)
	}

	key := fmt.Sprintf("deploy:%s", deploy.Name)
	err = r.conn.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("redis.Client.Set: %v", err)
	}

	return nil
}

func (r *Redis) GetDeployByName(ctx context.Context, name string) (*domain.Deploy, error) {
	key := fmt.Sprintf("deploy:%s", name)
	jsonData, err := r.conn.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("object with \"Name\" %s does not exist", name)
	} else if err != nil {
		return nil, fmt.Errorf("redis.Client.Get: %v", err)
	}

	var deploy domain.Deploy
	err = json.Unmarshal([]byte(jsonData), &deploy)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return &deploy, nil
}

func (r *Redis) DeleteDeployByName(ctx context.Context, name string) error {
	key := fmt.Sprintf("deploy:%s", name)
	err := r.conn.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("redis.Client.Del: %v", err)
	}
	return nil
}
