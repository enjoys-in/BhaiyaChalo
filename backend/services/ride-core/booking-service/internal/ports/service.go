package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/model"
)

type BookingService interface {
	CreateBooking(ctx context.Context, req *dto.CreateBookingRequest) (*model.Booking, error)
	GetBooking(ctx context.Context, id string) (*model.Booking, error)
	CancelBooking(ctx context.Context, req *dto.CancelBookingRequest) error
	UpdateStatus(ctx context.Context, req *dto.UpdateStatusRequest) error
	AssignDriver(ctx context.Context, bookingID, driverID string) error
}
