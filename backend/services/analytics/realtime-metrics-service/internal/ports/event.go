package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/model"
)

type EventPublisher interface {
	PublishMetricsComputed(ctx context.Context, entity *model.Metric) error
}
