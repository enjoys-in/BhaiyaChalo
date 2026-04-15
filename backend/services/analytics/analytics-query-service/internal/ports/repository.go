package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-query-service/internal/model"
)

type AnalyticsQueryRepository interface {
	Create(ctx context.Context, entity *model.QueryResult) error
	FindByID(ctx context.Context, id string) (*model.QueryResult, error)
	Update(ctx context.Context, entity *model.QueryResult) error
	Delete(ctx context.Context, id string) error
}
