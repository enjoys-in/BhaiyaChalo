package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/model"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *model.Payment) error
	FindByID(ctx context.Context, id string) (*model.Payment, error)
	FindByBookingID(ctx context.Context, bookingID string) (*model.Payment, error)
	UpdateStatus(ctx context.Context, id string, status model.PaymentStatus, gatewayID, gatewayStatus, failureReason string) error
	CreateRefund(ctx context.Context, refund *model.Refund) error
	FindRefundByPaymentID(ctx context.Context, paymentID string) (*model.Refund, error)
}
