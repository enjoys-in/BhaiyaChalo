package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/dto"
)

type RealtimeMetricsService interface {
	Create(ctx context.Context, req dto.CreateMetricRequest) (*dto.MetricResponse, error)
	GetByID(ctx context.Context, id string) (*dto.MetricResponse, error)
	Update(ctx context.Context, req dto.UpdateMetricRequest) (*dto.MetricResponse, error)
	Delete(ctx context.Context, id string) error
}
