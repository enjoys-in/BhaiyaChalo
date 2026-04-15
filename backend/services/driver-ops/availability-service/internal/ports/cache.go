package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/model"
)

type AvailabilityCache interface {
	Set(ctx context.Context, avail *model.DriverAvailability) error
	Get(ctx context.Context, driverID string) (*model.DriverAvailability, error)
	Delete(ctx context.Context, driverID string) error
	CountByCity(ctx context.Context, cityID, vehicleType string) (int, error)
}
