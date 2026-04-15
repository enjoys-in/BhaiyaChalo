package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/model"
)

type EventPublisher interface {
	PublishVehicleRegistered(ctx context.Context, vehicle *model.Vehicle) error
	PublishVehicleApproved(ctx context.Context, vehicle *model.Vehicle) error
	PublishVehicleExpired(ctx context.Context, vehicle *model.Vehicle) error
}
