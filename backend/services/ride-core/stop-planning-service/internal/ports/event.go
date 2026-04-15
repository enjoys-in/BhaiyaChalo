package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/model"
)

type EventPublisher interface {
	PublishStopAdded(ctx context.Context, stop *model.Stop) error
	PublishStopArrived(ctx context.Context, stop *model.Stop) error
	PublishStopCompleted(ctx context.Context, stop *model.Stop) error
}
