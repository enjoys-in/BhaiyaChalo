package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/model"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *model.Booking) error
	FindByID(ctx context.Context, id string) (*model.Booking, error)
	FindByUserID(ctx context.Context, userID string, limit, offset int) ([]*model.Booking, error)
	UpdateStatus(ctx context.Context, id string, status model.BookingStatus) error
	UpdateDriver(ctx context.Context, id string, driverID string) error
	UpdateFare(ctx context.Context, id string, finalFare float64) error
	Cancel(ctx context.Context, id string) error
}
