package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/user-service/internal/model"
)

type EventPublisher interface {
	PublishUserCreated(ctx context.Context, user *model.User) error
	PublishUserUpdated(ctx context.Context, user *model.User) error
	PublishUserDeleted(ctx context.Context, user *model.User) error
}
