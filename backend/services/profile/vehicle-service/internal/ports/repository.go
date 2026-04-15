package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/model"
)

type VehicleRepository interface {
	Create(ctx context.Context, vehicle *model.Vehicle) error
	FindByID(ctx context.Context, id string) (*model.Vehicle, error)
	FindByDriverID(ctx context.Context, driverID string) ([]*model.Vehicle, error)
	Update(ctx context.Context, vehicle *model.Vehicle) error
	Delete(ctx context.Context, id string) error
	ListByType(ctx context.Context, vehicleType model.VehicleType) ([]*model.Vehicle, error)
}
