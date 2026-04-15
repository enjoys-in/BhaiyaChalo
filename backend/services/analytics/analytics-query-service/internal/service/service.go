package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-query-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-query-service/internal/ports"
)

type analyticsQueryService struct {
	repo      ports.AnalyticsQueryRepository
	publisher ports.EventPublisher
}

func NewAnalyticsQueryService(repo ports.AnalyticsQueryRepository, publisher ports.EventPublisher) ports.AnalyticsQueryService {
	return &analyticsQueryService{repo: repo, publisher: publisher}
}

func (s *analyticsQueryService) Create(ctx context.Context, req dto.CreateQueryResultRequest) (*dto.QueryResultResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *analyticsQueryService) GetByID(ctx context.Context, id string) (*dto.QueryResultResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *analyticsQueryService) Update(ctx context.Context, req dto.UpdateQueryResultRequest) (*dto.QueryResultResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *analyticsQueryService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
