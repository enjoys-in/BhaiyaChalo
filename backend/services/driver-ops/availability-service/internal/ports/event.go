package ports

import "context"

type EventPublisher interface {
	PublishDriverOnline(ctx context.Context, driverID, cityID, vehicleType string) error
	PublishDriverOffline(ctx context.Context, driverID string) error
	PublishDriverBusy(ctx context.Context, driverID string) error
	PublishDriverFree(ctx context.Context, driverID string) error
}
