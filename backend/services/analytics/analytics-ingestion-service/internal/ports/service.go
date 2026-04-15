package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/dto"
)

type AnalyticsIngestionService interface {
	Create(ctx context.Context, req dto.CreateAnalyticsEventRequest) (*dto.AnalyticsEventResponse, error)
	GetByID(ctx context.Context, id string) (*dto.AnalyticsEventResponse, error)
	Update(ctx context.Context, req dto.UpdateAnalyticsEventRequest) (*dto.AnalyticsEventResponse, error)
	Delete(ctx context.Context, id string) error
}
