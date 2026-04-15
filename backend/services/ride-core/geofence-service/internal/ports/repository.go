package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/model"
)

type GeofenceRepository interface {
	Create(ctx context.Context, fence *model.Geofence) error
	FindByID(ctx context.Context, id string) (*model.Geofence, error)
	FindByCityID(ctx context.Context, cityID string) ([]*model.Geofence, error)
	Update(ctx context.Context, fence *model.Geofence) error
	Delete(ctx context.Context, id string) error
	FindContaining(ctx context.Context, lat, lng float64) ([]*model.Geofence, error)
}
