package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/model"
)

type NotificationRepository interface {
	Create(ctx context.Context, entity *model.Notification) error
	FindByID(ctx context.Context, id string) (*model.Notification, error)
	Update(ctx context.Context, entity *model.Notification) error
	Delete(ctx context.Context, id string) error
}
