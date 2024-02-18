// File: core/cache/cache.go

package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCacheService implements CacheService using Redis
type RedisCacheService struct {
	client *redis.Client
}

// NewRedisCacheService creates a new instance of RedisCacheService
func NewRedisCacheService(ctx context.Context, addr, password string) (*RedisCacheService, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // Use default DB
	})

	// Ping the Redis server to ensure connectivity
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisCacheService{
		client: client,
	}, nil
}

// Get retrieves value from cache by key
func (svc *RedisCacheService) Get(ctx context.Context, key string) (string, error) {
	val, err := svc.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil // Cache miss
		}
		return "", err
	}
	return val, nil
}

// Set sets value in cache with specified key
func (svc *RedisCacheService) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	err := svc.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
