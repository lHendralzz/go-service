package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLock struct {
	client *redis.Client
}

type Option struct {
	Host     string `env:"REDIS_HOST"`
	Password string `env:"REDIS_PASSWORD"`
	DB       string `env:"REDIS_DB"`
}

func NewRedisLock(client *redis.Client) *RedisLock {
	return &RedisLock{client: client}
}

// AcquireLock attempts to acquire a lock with the specified key and timeout.
func (rl *RedisLock) AcquireLock(ctx context.Context, key string, timeout time.Duration) (bool, error) {
	// Attempt to set the lock key
	success, err := rl.client.SetNX(ctx, key, true, timeout).Result()
	if err != nil {
		return false, err
	}
	return success, nil
}

// ReleaseLock releases the lock by deleting the key.
func (rl *RedisLock) ReleaseLock(ctx context.Context, key string) error {
	_, err := rl.client.Del(ctx, key).Result()
	return err
}
