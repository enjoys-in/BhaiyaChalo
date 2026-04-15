package redis

import (
	"context"
	"fmt"
	"time"
)

// Config holds Valkey/Redis connection settings.
type Config struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

// Client wraps a Valkey/Redis client.
// Placeholder — actual implementation uses github.com/redis/go-redis.
type Client struct {
	cfg Config
}

// Connect creates a new Valkey client.
func Connect(ctx context.Context, cfg Config) (*Client, error) {
	if cfg.Addr == "" {
		return nil, fmt.Errorf("redis: Addr is required")
	}
	// TODO: Replace with actual redis.NewClient when dependency is added.
	return &Client{cfg: cfg}, nil
}

// Set stores a key with TTL.
func (c *Client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	// TODO: c.inner.Set(ctx, key, value, ttl).Err()
	return nil
}

// Get retrieves a key.
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	// TODO: c.inner.Get(ctx, key).Result()
	return "", nil
}

// Del removes a key.
func (c *Client) Del(ctx context.Context, keys ...string) error {
	// TODO: c.inner.Del(ctx, keys...).Err()
	return nil
}

// Close shuts down the client.
func (c *Client) Close() error {
	// TODO: c.inner.Close()
	return nil
}
