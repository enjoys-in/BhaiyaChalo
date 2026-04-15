package ports

import "context"

type EventPublisher interface {
	PublishEarningRecorded(ctx context.Context, driverID, tripID string, netEarning float64) error
}
