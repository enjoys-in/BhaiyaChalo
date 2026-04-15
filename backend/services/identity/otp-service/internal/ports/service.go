package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/dto"
)

type OTPService interface {
	Send(ctx context.Context, req dto.SendOTPRequest) (*dto.SendOTPResponse, error)
	Verify(ctx context.Context, req dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, error)
}
