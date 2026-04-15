package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/ports"
)

type availabilityService struct {
	repo      ports.AvailabilityRepository
	cache     ports.AvailabilityCache
	publisher ports.EventPublisher
}

func NewAvailabilityService(
	repo ports.AvailabilityRepository,
	cache ports.AvailabilityCache,
	publisher ports.EventPublisher,
) ports.AvailabilityService {
	return &availabilityService{
		repo:      repo,
		cache:     cache,
		publisher: publisher,
	}
}

func (s *availabilityService) GoOnline(ctx context.Context, req dto.GoOnlineRequest) (*dto.AvailabilityResponse, error) {
	now := time.Now().UTC()
	avail := &model.DriverAvailability{
		DriverID:    req.DriverID,
		CityID:      req.CityID,
		Online:      true,
		OnTrip:      false,
		VehicleType: req.VehicleType,
		Lat:         req.Lat,
		Lng:         req.Lng,
		LastSeenAt:  now,
		UpdatedAt:   now,
	}

	if err := s.repo.SetOnline(ctx, avail); err != nil {
		return nil, fmt.Errorf("set online: %w", err)
	}

	if err := s.cache.Set(ctx, avail); err != nil {
		return nil, fmt.Errorf("cache set: %w", err)
	}

	s.logAction(ctx, req.DriverID, model.ActionWentOnline)
	_ = s.publisher.PublishDriverOnline(ctx, req.DriverID, req.CityID, req.VehicleType)

	return toResponse(avail), nil
}

func (s *availabilityService) GoOffline(ctx context.Context, req dto.GoOfflineRequest) error {
	if err := s.repo.SetOffline(ctx, req.DriverID); err != nil {
		return fmt.Errorf("set offline: %w", err)
	}

	if err := s.cache.Delete(ctx, req.DriverID); err != nil {
		return fmt.Errorf("cache delete: %w", err)
	}

	s.logAction(ctx, req.DriverID, model.ActionWentOffline)
	_ = s.publisher.PublishDriverOffline(ctx, req.DriverID)

	return nil
}

func (s *availabilityService) UpdateTripStatus(ctx context.Context, req dto.UpdateTripStatusRequest) error {
	if req.OnTrip {
		if err := s.repo.SetOnTrip(ctx, req.DriverID); err != nil {
			return fmt.Errorf("set on trip: %w", err)
		}
		s.logAction(ctx, req.DriverID, model.ActionTripStarted)
		_ = s.publisher.PublishDriverBusy(ctx, req.DriverID)
	} else {
		if err := s.repo.SetFree(ctx, req.DriverID); err != nil {
			return fmt.Errorf("set free: %w", err)
		}
		s.logAction(ctx, req.DriverID, model.ActionTripEnded)
		_ = s.publisher.PublishDriverFree(ctx, req.DriverID)
	}

	cached, _ := s.cache.Get(ctx, req.DriverID)
	if cached != nil {
		cached.OnTrip = req.OnTrip
		cached.UpdatedAt = time.Now().UTC()
		_ = s.cache.Set(ctx, cached)
	}

	return nil
}

func (s *availabilityService) GetStatus(ctx context.Context, driverID string) (*dto.AvailabilityResponse, error) {
	cached, err := s.cache.Get(ctx, driverID)
	if err == nil && cached != nil {
		return toResponse(cached), nil
	}

	avail, err := s.repo.GetStatus(ctx, driverID)
	if err != nil {
		return nil, fmt.Errorf("get status: %w", err)
	}
	if avail == nil {
		return nil, fmt.Errorf("driver not found")
	}
	return toResponse(avail), nil
}

func (s *availabilityService) CountOnlineDrivers(ctx context.Context, cityID, vehicleType string) (*dto.OnlineDriversResponse, error) {
	count, err := s.cache.CountByCity(ctx, cityID, vehicleType)
	if err != nil {
		count, err = s.repo.CountOnlineByCityAndType(ctx, cityID, vehicleType)
		if err != nil {
			return nil, fmt.Errorf("count online: %w", err)
		}
	}

	return &dto.OnlineDriversResponse{
		CityID:      cityID,
		VehicleType: vehicleType,
		Count:       count,
	}, nil
}

func (s *availabilityService) logAction(ctx context.Context, driverID string, action model.ActionType) {
	log := &model.AvailabilityLog{
		ID:        uuid.NewString(),
		DriverID:  driverID,
		Action:    action,
		Timestamp: time.Now().UTC(),
	}
	_ = s.repo.LogAction(ctx, log)
}

func toResponse(a *model.DriverAvailability) *dto.AvailabilityResponse {
	return &dto.AvailabilityResponse{
		DriverID:    a.DriverID,
		CityID:      a.CityID,
		Online:      a.Online,
		OnTrip:      a.OnTrip,
		VehicleType: a.VehicleType,
		Lat:         a.Lat,
		Lng:         a.Lng,
		LastSeenAt:  a.LastSeenAt,
		UpdatedAt:   a.UpdatedAt,
	}
}
