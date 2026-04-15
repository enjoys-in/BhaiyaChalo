package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/model"
)

type AnalyticsIngestionRepository interface {
	Create(ctx context.Context, entity *model.AnalyticsEvent) error
	FindByID(ctx context.Context, id string) (*model.AnalyticsEvent, error)
	Update(ctx context.Context, entity *model.AnalyticsEvent) error
	Delete(ctx context.Context, id string) error
}
