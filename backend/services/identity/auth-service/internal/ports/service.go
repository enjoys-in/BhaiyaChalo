package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
	Refresh(ctx context.Context, req dto.RefreshRequest) (*dto.TokenResponse, error)
	Verify(ctx context.Context, req dto.VerifyRequest) (*dto.VerifyResponse, error)
	Logout(ctx context.Context, req dto.LogoutRequest) error
}
