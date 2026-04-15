package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/model"
)

type EventPublisher interface {
	PublishBookingCreated(ctx context.Context, booking *model.Booking) error
	PublishBookingConfirmed(ctx context.Context, booking *model.Booking) error
	PublishBookingCancelled(ctx context.Context, bookingID string) error
	PublishDriverAssigned(ctx context.Context, bookingID, driverID string) error
}
