package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/model"
)

type EventPublisher interface {
	PublishNotificationSent(ctx context.Context, entity *model.Notification) error
	PublishNotificationFailed(ctx context.Context, entity *model.Notification) error
}
