package cache

import (
	"context"
	"encoding/json"
	"time"
	
	"github.com/redis/go-redis/v9"
)

// RedisCache provides a Redis-backed caching implementation
type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisCache creates a new Redis cache client
func NewRedisCache(addr, password string, ttlSeconds int) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	
	return &RedisCache{
		client: client,
		ttl:    time.Duration(ttlSeconds) * time.Second,
	}, nil
}

// Get retrieves a value from cache and unmarshals it into the destination
func (rc *RedisCache) Get(ctx context.Context, key string, dest interface{}) (bool, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil // Key not found
	} else if err != nil {
		return false, err
	}
	
	if err := json.Unmarshal([]byte(val), dest); err != nil {
		return false, err
	}
	
	return true, nil
}

// Set serializes and stores a value in the cache with the default TTL
func (rc *RedisCache) Set(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	
	return rc.client.Set(ctx, key, data, rc.ttl).Err()
}

// SetWithTTL serializes and stores a value with a custom TTL
func (rc *RedisCache) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	
	return rc.client.Set(ctx, key, data, ttl).Err()
}

// Delete removes a key from the cache
func (rc *RedisCache) Delete(ctx context.Context, key string) error {
	return rc.client.Del(ctx, key).Err()
}

// Clear removes all keys matching the given pattern
func (rc *RedisCache) Clear(ctx context.Context, pattern string) error {
	keys, err := rc.client.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}
	
	if len(keys) > 0 {
		return rc.client.Del(ctx, keys...).Err()
	}
	
	return nil
}

// Close closes the underlying Redis connection
func (rc *RedisCache) Close() error {
	return rc.client.Close()
}