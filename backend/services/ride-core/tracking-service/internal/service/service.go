package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/config"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/ports"
)

type trackingService struct {
	cache     ports.LocationCache
	repo      ports.TrackingRepository
	publisher ports.TrackingEventPublisher
	cfg       *config.Config
}

func NewTrackingService(
	cache ports.LocationCache,
	repo ports.TrackingRepository,
	publisher ports.TrackingEventPublisher,
	cfg *config.Config,
) ports.TrackingService {
	return &trackingService{
		cache:     cache,
		repo:      repo,
		publisher: publisher,
		cfg:       cfg,
	}
}

func (s *trackingService) UpdateLocation(ctx context.Context, req *dto.UpdateLocationRequest) error {
	if err := s.cache.Set(ctx, req.DriverID, req.Lat, req.Lng, s.cfg.LocationTTLSeconds); err != nil {
		return fmt.Errorf("cache set location: %w", err)
	}

	loc := &model.LocationUpdate{
		DriverID:  req.DriverID,
		Lat:       req.Lat,
		Lng:       req.Lng,
		Heading:   req.Heading,
		Speed:     req.Speed,
		Accuracy:  req.Accuracy,
		Timestamp: time.Now().UTC(),
	}

	if err := s.publisher.PublishLocationUpdated(ctx, loc); err != nil {
		return fmt.Errorf("publish location updated: %w", err)
	}

	if err := s.repo.SaveLocation(ctx, loc); err != nil {
		return fmt.Errorf("save location: %w", err)
	}

	return nil
}

func (s *trackingService) GetLocation(ctx context.Context, driverID string) (*model.LocationUpdate, error) {
	lat, lng, err := s.cache.Get(ctx, driverID)
	if err == nil {
		return &model.LocationUpdate{
			DriverID:  driverID,
			Lat:       lat,
			Lng:       lng,
			Timestamp: time.Now().UTC(),
		}, nil
	}

	loc, err := s.repo.GetLatestLocation(ctx, driverID)
	if err != nil {
		return nil, fmt.Errorf("get latest location: %w", err)
	}
	return loc, nil
}

func (s *trackingService) StartTracking(ctx context.Context, req *dto.StartTrackingRequest) (*model.TrackingSession, error) {
	existing, _ := s.repo.FindActiveSession(ctx, req.TripID)
	if existing != nil {
		return existing, nil
	}

	session := &model.TrackingSession{
		ID:        uuid.New().String(),
		TripID:    req.TripID,
		DriverID:  req.DriverID,
		UserID:    req.UserID,
		Active:    true,
		StartedAt: time.Now().UTC(),
	}

	if err := s.repo.StartSession(ctx, session); err != nil {
		return nil, fmt.Errorf("start session: %w", err)
	}

	if err := s.publisher.PublishTrackingStarted(ctx, session); err != nil {
		return session, fmt.Errorf("publish tracking started: %w", err)
	}

	return session, nil
}

func (s *trackingService) StopTracking(ctx context.Context, tripID string) error {
	session, err := s.repo.FindActiveSession(ctx, tripID)
	if err != nil {
		return fmt.Errorf("find active session: %w", err)
	}

	if err := s.repo.EndSession(ctx, tripID); err != nil {
		return fmt.Errorf("end session: %w", err)
	}

	session.Active = false
	now := time.Now().UTC()
	session.EndedAt = &now

	return s.publisher.PublishTrackingStopped(ctx, session)
}
