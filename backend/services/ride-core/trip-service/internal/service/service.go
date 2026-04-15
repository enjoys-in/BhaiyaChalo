package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/ports"
)

type tripService struct {
	repo      ports.TripRepository
	publisher ports.EventPublisher
}

func NewTripService(repo ports.TripRepository, publisher ports.EventPublisher) ports.TripService {
	return &tripService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *tripService) Create(ctx context.Context, req dto.CreateTripRequest) (*model.Trip, error) {
	now := time.Now().UTC()
	trip := &model.Trip{
		ID:            uuid.New().String(),
		BookingID:     req.BookingID,
		UserID:        req.UserID,
		DriverID:      req.DriverID,
		CityID:        req.CityID,
		VehicleType:   req.VehicleType,
		Status:        model.TripStatusDriverEnroute,
		PickupLat:     req.PickupLat,
		PickupLng:     req.PickupLng,
		DropLat:       req.DropLat,
		DropLng:       req.DropLng,
		FareAmount:    req.FareAmount,
		PaymentMethod: req.PaymentMethod,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := s.repo.Create(ctx, trip); err != nil {
		return nil, err
	}

	timeline := &model.TripTimeline{
		TripID:    trip.ID,
		Event:     "trip_created",
		Timestamp: now,
	}
	_ = s.repo.AddTimelineEvent(ctx, timeline)

	return trip, nil
}

func (s *tripService) Get(ctx context.Context, tripID string) (*model.Trip, error) {
	return s.repo.FindByID(ctx, tripID)
}

func (s *tripService) UpdateStatus(ctx context.Context, req dto.UpdateTripStatusRequest) error {
	status := model.TripStatus(req.Status)

	if err := s.repo.UpdateStatus(ctx, req.TripID, status); err != nil {
		return err
	}

	timeline := &model.TripTimeline{
		TripID:    req.TripID,
		Event:     req.Status,
		Timestamp: time.Now().UTC(),
	}
	_ = s.repo.AddTimelineEvent(ctx, timeline)

	trip, err := s.repo.FindByID(ctx, req.TripID)
	if err != nil {
		return nil
	}

	switch status {
	case model.TripStatusStarted:
		_ = s.publisher.PublishTripStarted(ctx, trip)
	case model.TripStatusCompleted:
		_ = s.publisher.PublishTripCompleted(ctx, trip)
	case model.TripStatusCancelled:
		_ = s.publisher.PublishTripCancelled(ctx, trip)
	}

	return nil
}

func (s *tripService) ListByUser(ctx context.Context, userID string, limit, offset int) ([]model.Trip, error) {
	return s.repo.FindByUserID(ctx, userID, limit, offset)
}

func (s *tripService) ListByDriver(ctx context.Context, driverID string, limit, offset int) ([]model.Trip, error) {
	return s.repo.FindByDriverID(ctx, driverID, limit, offset)
}
