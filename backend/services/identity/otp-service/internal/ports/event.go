package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/model"
)

type EventPublisher interface {
	PublishOTPSent(ctx context.Context, otp *model.OTP) error
	PublishOTPVerified(ctx context.Context, otp *model.OTP) error
}
