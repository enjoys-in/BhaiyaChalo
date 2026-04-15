package grpc

import (
	"fmt"
	"log/slog"
)

// ServerConfig holds gRPC server settings.
type ServerConfig struct {
	Port int
}

// ClientConfig holds gRPC client connection settings.
type ClientConfig struct {
	Addr string
}

// NewServerAddr returns the listen address string.
func NewServerAddr(cfg ServerConfig) string {
	return fmt.Sprintf(":%d", cfg.Port)
}

// Dial creates a gRPC client connection.
// Placeholder — actual implementation uses google.golang.org/grpc.
func Dial(cfg ClientConfig, logger *slog.Logger) error {
	logger.Info("grpc dial placeholder", "addr", cfg.Addr)
	// TODO: grpc.Dial(cfg.Addr, ...)
	return nil
}
