package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/model"
)

type LocationIngestRepository interface {
	Create(ctx context.Context, entity *model.LocationUpdate) error
	FindByID(ctx context.Context, id string) (*model.LocationUpdate, error)
	Update(ctx context.Context, entity *model.LocationUpdate) error
	Delete(ctx context.Context, id string) error
}
