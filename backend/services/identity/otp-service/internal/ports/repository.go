package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/model"
)

type OTPRepository interface {
	Create(ctx context.Context, otp *model.OTP) error
	FindLatest(ctx context.Context, phone string, purpose model.Purpose) (*model.OTP, error)
	MarkVerified(ctx context.Context, id string) error
	IncrementAttempts(ctx context.Context, id string) error
}
