package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/ports"
)

type bookingService struct {
	repo      ports.BookingRepository
	publisher ports.EventPublisher
}

func NewBookingService(repo ports.BookingRepository, publisher ports.EventPublisher) ports.BookingService {
	return &bookingService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *bookingService) CreateBooking(ctx context.Context, req *dto.CreateBookingRequest) (*model.Booking, error) {
	now := time.Now().UTC()

	// TODO: call pricing-service via gRPC to get estimated fare
	estimatedFare := 0.0

	// TODO: call promo-service via gRPC to validate promo code and get discount
	discountAmount := 0.0

	booking := &model.Booking{
		ID:             uuid.New().String(),
		UserID:         req.UserID,
		CityID:         req.CityID,
		PickupLat:      req.PickupLat,
		PickupLng:      req.PickupLng,
		PickupAddress:  req.PickupAddress,
		DropLat:        req.DropLat,
		DropLng:        req.DropLng,
		DropAddress:    req.DropAddress,
		VehicleType:    req.VehicleType,
		EstimatedFare:  estimatedFare,
		FinalFare:      0,
		PromoCode:      req.PromoCode,
		DiscountAmount: discountAmount,
		Status:         model.StatusPending,
		PaymentMethod:  req.PaymentMethod,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repo.Create(ctx, booking); err != nil {
		return nil, fmt.Errorf("create booking: %w", err)
	}

	_ = s.publisher.PublishBookingCreated(ctx, booking)

	return booking, nil
}

func (s *bookingService) GetBooking(ctx context.Context, id string) (*model.Booking, error) {
	booking, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get booking: %w", err)
	}
	return booking, nil
}

func (s *bookingService) CancelBooking(ctx context.Context, req *dto.CancelBookingRequest) error {
	booking, err := s.repo.FindByID(ctx, req.BookingID)
	if err != nil {
		return fmt.Errorf("find booking: %w", err)
	}

	if booking.Status == model.StatusCompleted || booking.Status == model.StatusCancelled {
		return fmt.Errorf("cannot cancel booking with status %s", booking.Status)
	}

	if err := s.repo.Cancel(ctx, req.BookingID); err != nil {
		return fmt.Errorf("cancel booking: %w", err)
	}

	_ = s.publisher.PublishBookingCancelled(ctx, req.BookingID)

	return nil
}

func (s *bookingService) UpdateStatus(ctx context.Context, req *dto.UpdateStatusRequest) error {
	status := model.BookingStatus(req.Status)

	if err := s.repo.UpdateStatus(ctx, req.BookingID, status); err != nil {
		return fmt.Errorf("update status: %w", err)
	}

	if status == model.StatusConfirmed {
		booking, _ := s.repo.FindByID(ctx, req.BookingID)
		if booking != nil {
			_ = s.publisher.PublishBookingConfirmed(ctx, booking)
		}
	}

	return nil
}

func (s *bookingService) AssignDriver(ctx context.Context, bookingID, driverID string) error {
	if err := s.repo.UpdateDriver(ctx, bookingID, driverID); err != nil {
		return fmt.Errorf("assign driver: %w", err)
	}

	_ = s.publisher.PublishDriverAssigned(ctx, bookingID, driverID)

	return nil
}
