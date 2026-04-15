package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/model"
)

type EventPublisher interface {
	PublishEventIngested(ctx context.Context, entity *model.AnalyticsEvent) error
}
