package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/model"
)

type EventPublisher interface {
	PublishGeofenceCreated(ctx context.Context, fence *model.Geofence) error
	PublishGeofenceUpdated(ctx context.Context, fence *model.Geofence) error
}
