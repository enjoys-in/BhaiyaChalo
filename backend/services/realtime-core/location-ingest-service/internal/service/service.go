package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/ports"
)

type locationIngestService struct {
	repo      ports.LocationIngestRepository
	publisher ports.EventPublisher
}

func NewLocationIngestService(repo ports.LocationIngestRepository, publisher ports.EventPublisher) ports.LocationIngestService {
	return &locationIngestService{repo: repo, publisher: publisher}
}

func (s *locationIngestService) Create(ctx context.Context, req dto.CreateLocationUpdateRequest) (*dto.LocationUpdateResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *locationIngestService) GetByID(ctx context.Context, id string) (*dto.LocationUpdateResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *locationIngestService) Update(ctx context.Context, req dto.UpdateLocationUpdateRequest) (*dto.LocationUpdateResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *locationIngestService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
