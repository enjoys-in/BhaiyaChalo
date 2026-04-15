package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/model"
)

type EventPublisher interface {
	PublishSessionCreated(ctx context.Context, session *model.Session) error
	PublishSessionInvalidated(ctx context.Context, sessionID string, userID string) error
}
