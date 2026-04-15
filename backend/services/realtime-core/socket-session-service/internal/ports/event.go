package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/model"
)

type EventPublisher interface {
	PublishSessionConnected(ctx context.Context, session *model.SocketSession) error
	PublishSessionDisconnected(ctx context.Context, session *model.SocketSession) error
}
