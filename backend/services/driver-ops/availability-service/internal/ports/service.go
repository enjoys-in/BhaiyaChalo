package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/dto"
)

type AvailabilityService interface {
	GoOnline(ctx context.Context, req dto.GoOnlineRequest) (*dto.AvailabilityResponse, error)
	GoOffline(ctx context.Context, req dto.GoOfflineRequest) error
	UpdateTripStatus(ctx context.Context, req dto.UpdateTripStatusRequest) error
	GetStatus(ctx context.Context, driverID string) (*dto.AvailabilityResponse, error)
	CountOnlineDrivers(ctx context.Context, cityID, vehicleType string) (*dto.OnlineDriversResponse, error)
}
