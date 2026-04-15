package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/model"
)

type RealtimeMetricsRepository interface {
	Create(ctx context.Context, entity *model.Metric) error
	FindByID(ctx context.Context, id string) (*model.Metric, error)
	Update(ctx context.Context, entity *model.Metric) error
	Delete(ctx context.Context, id string) error
}
