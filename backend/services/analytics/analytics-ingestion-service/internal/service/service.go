package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/ports"
)

type analyticsIngestionService struct {
	repo      ports.AnalyticsIngestionRepository
	publisher ports.EventPublisher
}

func NewAnalyticsIngestionService(repo ports.AnalyticsIngestionRepository, publisher ports.EventPublisher) ports.AnalyticsIngestionService {
	return &analyticsIngestionService{repo: repo, publisher: publisher}
}

func (s *analyticsIngestionService) Create(ctx context.Context, req dto.CreateAnalyticsEventRequest) (*dto.AnalyticsEventResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *analyticsIngestionService) GetByID(ctx context.Context, id string) (*dto.AnalyticsEventResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *analyticsIngestionService) Update(ctx context.Context, req dto.UpdateAnalyticsEventRequest) (*dto.AnalyticsEventResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *analyticsIngestionService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
