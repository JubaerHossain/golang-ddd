// File: core/cache/cache.go

package cache

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheService interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, expiration time.Duration) error
	Remove(ctx context.Context, key string) error
}

// RedisCacheService implements CacheService using Redis
type RedisCacheService struct {
	client *redis.Client
}

// NewRedisCacheService creates a new instance of RedisCacheService
func NewRedisCacheService(ctx context.Context) (*RedisCacheService, error) {
	// Get Redis server address and password from environment variables
	if os.Getenv("IS_REDIS") != "true" {
		return nil, nil		
	}
	redisURI := os.Getenv("REDIS_URI")
	if redisURI == "" {
		redisURI = "localhost:6379" // Default Redis server address
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		redisPassword = "" // No password if not provided
	}

	redisDB := os.Getenv("REDIS_DB")
	if redisDB == "" {
		redisDB = "0" // Default Redis DB
	}
	redisDBInt, _ := strconv.Atoi(redisDB)
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDBInt,
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

// Remove implements CacheService.
func (svc *RedisCacheService) Remove(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err := svc.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
