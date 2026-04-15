package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/model"
)

type SessionRepository interface {
	Register(ctx context.Context, session *model.SocketSession) error
	Unregister(ctx context.Context, sessionID string) error
	FindByUserID(ctx context.Context, userID string) ([]*model.SocketSession, error)
	FindActiveByServer(ctx context.Context, serverID string) ([]*model.SocketSession, error)
	CountActive(ctx context.Context) (int64, error)
}
