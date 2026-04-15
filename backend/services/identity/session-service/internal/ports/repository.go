package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/model"
)

type SessionRepository interface {
	Create(ctx context.Context, session *model.Session) error
	FindByID(ctx context.Context, id string) (*model.Session, error)
	FindByUserID(ctx context.Context, userID string) ([]*model.Session, error)
	Invalidate(ctx context.Context, id string) error
	InvalidateAllByUser(ctx context.Context, userID string) error
}
