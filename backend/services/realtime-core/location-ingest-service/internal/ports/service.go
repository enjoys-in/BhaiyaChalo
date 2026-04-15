package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/dto"
)

type LocationIngestService interface {
	Create(ctx context.Context, req dto.CreateLocationUpdateRequest) (*dto.LocationUpdateResponse, error)
	GetByID(ctx context.Context, id string) (*dto.LocationUpdateResponse, error)
	Update(ctx context.Context, req dto.UpdateLocationUpdateRequest) (*dto.LocationUpdateResponse, error)
	Delete(ctx context.Context, id string) error
}
