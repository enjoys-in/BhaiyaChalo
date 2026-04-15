package mongodb

import (
	"context"
	"fmt"
	"time"
)

// Config holds MongoDB connection settings.
type Config struct {
	URI      string
	Database string
	Timeout  time.Duration
}

// Client wraps a MongoDB client connection.
// Placeholder — actual implementation uses go.mongodb.org/mongo-driver.
type Client struct {
	cfg Config
}

// Connect creates a new MongoDB connection.
func Connect(ctx context.Context, cfg Config) (*Client, error) {
	if cfg.URI == "" {
		return nil, fmt.Errorf("mongodb: URI is required")
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 10 * time.Second
	}
	// TODO: Replace with actual mongo.Connect when dependency is added.
	return &Client{cfg: cfg}, nil
}

// Database returns the configured database name.
func (c *Client) Database() string { return c.cfg.Database }

// Close disconnects the client.
func (c *Client) Close(ctx context.Context) error {
	// TODO: c.inner.Disconnect(ctx)
	return nil
}
