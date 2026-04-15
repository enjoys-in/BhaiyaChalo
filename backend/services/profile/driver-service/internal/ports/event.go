package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/model"
)

type EventPublisher interface {
	PublishDriverCreated(ctx context.Context, driver *model.Driver) error
	PublishDriverUpdated(ctx context.Context, driver *model.Driver) error
	PublishDriverStatusChanged(ctx context.Context, driver *model.Driver, oldStatus string) error
}
