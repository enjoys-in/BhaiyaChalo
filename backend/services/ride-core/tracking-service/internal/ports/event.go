package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/model"
)

type TrackingEventPublisher interface {
	PublishLocationUpdated(ctx context.Context, loc *model.LocationUpdate) error
	PublishTrackingStarted(ctx context.Context, session *model.TrackingSession) error
	PublishTrackingStopped(ctx context.Context, session *model.TrackingSession) error
}
