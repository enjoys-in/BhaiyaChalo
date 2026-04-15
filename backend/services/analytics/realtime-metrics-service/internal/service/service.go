package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/ports"
)

type realtimeMetricsService struct {
	repo      ports.RealtimeMetricsRepository
	publisher ports.EventPublisher
}

func NewRealtimeMetricsService(repo ports.RealtimeMetricsRepository, publisher ports.EventPublisher) ports.RealtimeMetricsService {
	return &realtimeMetricsService{repo: repo, publisher: publisher}
}

func (s *realtimeMetricsService) Create(ctx context.Context, req dto.CreateMetricRequest) (*dto.MetricResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *realtimeMetricsService) GetByID(ctx context.Context, id string) (*dto.MetricResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *realtimeMetricsService) Update(ctx context.Context, req dto.UpdateMetricRequest) (*dto.MetricResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *realtimeMetricsService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
