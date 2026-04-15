package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/ports"
)

type stopService struct {
	repo      ports.StopRepository
	publisher ports.EventPublisher
}

func NewStopService(repo ports.StopRepository, publisher ports.EventPublisher) ports.StopService {
	return &stopService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *stopService) AddStop(ctx context.Context, req dto.AddStopRequest) (*model.Stop, error) {
	now := time.Now().UTC()
	stop := &model.Stop{
		ID:        uuid.New().String(),
		TripID:    req.TripID,
		Lat:       req.Lat,
		Lng:       req.Lng,
		Address:   req.Address,
		StopOrder: req.StopOrder,
		Status:    model.StopStatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.AddStop(ctx, stop); err != nil {
		return nil, err
	}

	if err := s.publisher.PublishStopAdded(ctx, stop); err != nil {
		return nil, err
	}

	return stop, nil
}

func (s *stopService) RemoveStop(ctx context.Context, req dto.RemoveStopRequest) error {
	return s.repo.RemoveStop(ctx, req.TripID, req.StopID)
}

func (s *stopService) ReorderStops(ctx context.Context, req dto.ReorderStopsRequest) (*model.MultiStopTrip, error) {
	if err := s.repo.ReorderStops(ctx, req.TripID, req.StopIDs); err != nil {
		return nil, err
	}
	return s.repo.FindByTripID(ctx, req.TripID)
}

func (s *stopService) UpdateStopStatus(ctx context.Context, req dto.UpdateStopStatusRequest) (*model.Stop, error) {
	status := model.StopStatus(req.Status)

	stop, err := s.repo.UpdateStopStatus(ctx, req.TripID, req.StopID, status)
	if err != nil {
		return nil, err
	}

	switch status {
	case model.StopStatusArrived:
		_ = s.publisher.PublishStopArrived(ctx, stop)
	case model.StopStatusCompleted:
		_ = s.publisher.PublishStopCompleted(ctx, stop)
	}

	return stop, nil
}

func (s *stopService) GetStopsByTrip(ctx context.Context, tripID string) (*model.MultiStopTrip, error) {
	return s.repo.FindByTripID(ctx, tripID)
}
