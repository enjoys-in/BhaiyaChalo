package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/dto"
)

type VehicleService interface {
	Create(ctx context.Context, req dto.CreateVehicleRequest) (*dto.VehicleResponse, error)
	GetByID(ctx context.Context, id string) (*dto.VehicleResponse, error)
	GetByDriverID(ctx context.Context, driverID string) ([]dto.VehicleResponse, error)
	Update(ctx context.Context, id string, req dto.UpdateVehicleRequest) (*dto.VehicleResponse, error)
	Delete(ctx context.Context, id string) error
	ListByType(ctx context.Context, vehicleType string) ([]dto.VehicleResponse, error)
}
