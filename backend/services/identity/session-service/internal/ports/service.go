package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/dto"
)

type SessionService interface {
	Create(ctx context.Context, req *dto.CreateSessionRequest) (*dto.SessionResponse, error)
	GetByID(ctx context.Context, req *dto.GetSessionRequest) (*dto.SessionResponse, error)
	Invalidate(ctx context.Context, req *dto.InvalidateRequest) error
	InvalidateAll(ctx context.Context, req *dto.InvalidateAllRequest) error
}
