package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/dto"
)

type GeofenceService interface {
	Create(ctx context.Context, req dto.CreateGeofenceRequest) (*dto.GeofenceResponse, error)
	Get(ctx context.Context, id string) (*dto.GeofenceResponse, error)
	List(ctx context.Context, cityID string) ([]dto.GeofenceResponse, error)
	Update(ctx context.Context, id string, req dto.UpdateGeofenceRequest) (*dto.GeofenceResponse, error)
	Delete(ctx context.Context, id string) error
	CheckPoint(ctx context.Context, req dto.PointInFenceRequest) (*dto.PointCheckResponse, error)
}
