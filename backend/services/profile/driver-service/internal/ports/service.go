package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/dto"
)

type DriverService interface {
	Create(ctx context.Context, req *dto.CreateDriverRequest) (*dto.DriverResponse, error)
	GetByID(ctx context.Context, id string) (*dto.DriverResponse, error)
	GetByPhone(ctx context.Context, phone string) (*dto.DriverResponse, error)
	Update(ctx context.Context, id string, req *dto.UpdateDriverRequest) (*dto.DriverResponse, error)
	Delete(ctx context.Context, id string) error
	UpdatePreference(ctx context.Context, driverID string, req *dto.UpdatePreferenceRequest) (*dto.DriverPreferenceResponse, error)
	GetPreference(ctx context.Context, driverID string) (*dto.DriverPreferenceResponse, error)
	ListByCityID(ctx context.Context, cityID string) ([]*dto.DriverResponse, error)
}
