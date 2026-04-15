package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/model"
)

type AvailabilityRepository interface {
	SetOnline(ctx context.Context, avail *model.DriverAvailability) error
	SetOffline(ctx context.Context, driverID string) error
	SetOnTrip(ctx context.Context, driverID string) error
	SetFree(ctx context.Context, driverID string) error
	GetStatus(ctx context.Context, driverID string) (*model.DriverAvailability, error)
	CountOnlineByCityAndType(ctx context.Context, cityID, vehicleType string) (int, error)
	LogAction(ctx context.Context, log *model.AvailabilityLog) error
}
