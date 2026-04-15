package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/model"
)

type PaymentService interface {
	Initiate(ctx context.Context, req dto.InitiatePaymentRequest) (*model.Payment, error)
	Capture(ctx context.Context, req dto.CapturePaymentRequest) (*model.Payment, error)
	Refund(ctx context.Context, req dto.RefundRequest) (*model.Refund, error)
	GetStatus(ctx context.Context, paymentID string) (*model.Payment, error)
}
