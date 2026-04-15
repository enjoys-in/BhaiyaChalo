package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/ports"
)

type searchService struct {
	repo      ports.SearchRepository
	publisher ports.EventPublisher
}

func NewSearchService(repo ports.SearchRepository, publisher ports.EventPublisher) ports.SearchService {
	return &searchService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *searchService) Search(ctx context.Context, req *dto.SearchRequest) (*dto.SearchResponse, error) {
	queryID := uuid.New().String()

	query := &model.SearchQuery{
		ID:        queryID,
		UserID:    req.UserID,
		CityID:    req.CityID,
		PickupLat: req.PickupLat,
		PickupLng: req.PickupLng,
		DropLat:   req.DropLat,
		DropLng:   req.DropLng,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.SaveQuery(ctx, query); err != nil {
		return nil, fmt.Errorf("save search query: %w", err)
	}

	// TODO: call ETA service via gRPC for each vehicle type
	// TODO: call Pricing service via gRPC for fare estimates

	vehicleTypes := []string{"auto", "mini", "sedan", "suv"}
	options := make([]dto.VehicleOption, 0, len(vehicleTypes))

	for _, vt := range vehicleTypes {
		options = append(options, dto.VehicleOption{
			VehicleType:   vt,
			Fare:          0, // placeholder: pricing service result
			ETA:           0, // placeholder: ETA service result
			Surge:         1.0,
			DriversNearby: 0, // placeholder: availability check
		})
	}

	_ = s.publisher.PublishSearchPerformed(ctx, query)

	return &dto.SearchResponse{
		QueryID: queryID,
		Results: options,
	}, nil
}
