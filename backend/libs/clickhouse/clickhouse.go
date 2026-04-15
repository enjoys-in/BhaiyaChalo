package clickhouse

import (
	"context"
	"fmt"
)

// Config holds ClickHouse connection settings.
type Config struct {
	Addr     string
	Database string
	User     string
	Password string
}

// Client wraps a ClickHouse connection.
type Client struct {
	cfg Config
}

// Connect creates a ClickHouse client.
func Connect(ctx context.Context, cfg Config) (*Client, error) {
	if cfg.Addr == "" {
		return nil, fmt.Errorf("clickhouse: Addr is required")
	}
	// TODO: actual connection when dependency is added
	return &Client{cfg: cfg}, nil
}

// Close disconnects the client.
func (c *Client) Close() error { return nil }
