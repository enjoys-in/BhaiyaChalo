package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/model"
)

type DriverRepository interface {
	Create(ctx context.Context, driver *model.Driver) error
	FindByID(ctx context.Context, id string) (*model.Driver, error)
	FindByPhone(ctx context.Context, phone string) (*model.Driver, error)
	Update(ctx context.Context, driver *model.Driver) error
	Delete(ctx context.Context, id string) error
	UpdatePreference(ctx context.Context, pref *model.DriverPreference) error
	GetPreference(ctx context.Context, driverID string) (*model.DriverPreference, error)
	ListByCityID(ctx context.Context, cityID string) ([]*model.Driver, error)
}
