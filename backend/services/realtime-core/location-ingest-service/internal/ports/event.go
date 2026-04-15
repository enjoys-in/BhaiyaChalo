package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/model"
)

type EventPublisher interface {
	PublishLocationUpdated(ctx context.Context, entity *model.LocationUpdate) error
	PublishLocationBatch(ctx context.Context, entity *model.LocationUpdate) error
}
