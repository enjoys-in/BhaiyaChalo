package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/model"
)

type EventPublisher interface {
	PublishPaymentInitiated(ctx context.Context, payment *model.Payment) error
	PublishPaymentCaptured(ctx context.Context, payment *model.Payment) error
	PublishPaymentFailed(ctx context.Context, payment *model.Payment) error
	PublishRefundIssued(ctx context.Context, refund *model.Refund) error
}
