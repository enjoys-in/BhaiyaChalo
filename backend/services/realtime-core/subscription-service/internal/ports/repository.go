package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/model"
)

type SubscriptionRepository interface {
	Subscribe(ctx context.Context, sub *model.Subscription) error
	Unsubscribe(ctx context.Context, subscriptionID string) error
	FindByUser(ctx context.Context, userID string) ([]*model.Subscription, error)
	FindByChannel(ctx context.Context, channel model.Channel, topic string) ([]*model.Subscription, error)
	CountByChannel(ctx context.Context, channel model.Channel) (int64, error)
}
