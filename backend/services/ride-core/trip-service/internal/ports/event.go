package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/model"
)

type EventPublisher interface {
	PublishTripStarted(ctx context.Context, trip *model.Trip) error
	PublishTripCompleted(ctx context.Context, trip *model.Trip) error
	PublishTripCancelled(ctx context.Context, trip *model.Trip) error
}
