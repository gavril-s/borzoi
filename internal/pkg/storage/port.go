package storage

import (
	"context"
	"fmt"
)

func (r *Redis) IsPortBusy(ctx context.Context, port int) (bool, error) {
	key := fmt.Sprintf("port:%d", port)
	exists, err := r.conn.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("redis.Client.Exists: %v", err)
	}
	return exists > 0, nil
}

func (r *Redis) MarkPortAsBusy(ctx context.Context, port int) error {
	key := fmt.Sprintf("port:%d", port)
	err := r.conn.Set(ctx, key, "busy", 0).Err()
	if err != nil {
		return fmt.Errorf("redis.Client.Set: %v", err)
	}
	return nil
}
