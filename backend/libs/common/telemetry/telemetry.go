package telemetry

import (
	"context"
	"log/slog"
)

// Config holds telemetry setup configuration.
type Config struct {
	ServiceName string
	Environment string
	OTLPAddr    string
}

// Setup initialises tracing and metrics exporters.
// Placeholder — integrate with OpenTelemetry SDK when dependencies are added.
func Setup(ctx context.Context, cfg Config, logger *slog.Logger) (shutdown func(context.Context) error, err error) {
	logger.Info("telemetry setup placeholder",
		"service", cfg.ServiceName,
		"env", cfg.Environment,
		"otlp", cfg.OTLPAddr,
	)
	return func(_ context.Context) error { return nil }, nil
}
