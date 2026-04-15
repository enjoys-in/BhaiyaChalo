package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-query-service/internal/dto"
)

type AnalyticsQueryService interface {
	Create(ctx context.Context, req dto.CreateQueryResultRequest) (*dto.QueryResultResponse, error)
	GetByID(ctx context.Context, id string) (*dto.QueryResultResponse, error)
	Update(ctx context.Context, req dto.UpdateQueryResultRequest) (*dto.QueryResultResponse, error)
	Delete(ctx context.Context, id string) error
}
