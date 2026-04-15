package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/model"
)

type SessionService interface {
	Register(ctx context.Context, req dto.RegisterSessionRequest) (*model.SocketSession, error)
	Unregister(ctx context.Context, req dto.UnregisterSessionRequest) error
	FindByUserID(ctx context.Context, userID string) ([]*model.SocketSession, error)
	FindActiveByServer(ctx context.Context, serverID string) ([]*model.SocketSession, error)
	CountActive(ctx context.Context) (int64, error)
}
