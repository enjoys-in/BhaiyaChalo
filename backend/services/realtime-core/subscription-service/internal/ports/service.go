package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/model"
)

type SubscriptionService interface {
	Subscribe(ctx context.Context, req dto.SubscribeRequest) (*model.Subscription, error)
	Unsubscribe(ctx context.Context, req dto.UnsubscribeRequest) error
	FindByUser(ctx context.Context, userID string) ([]*model.Subscription, error)
	FindByChannel(ctx context.Context, channel string, topic string) ([]*model.Subscription, error)
	CountByChannel(ctx context.Context, channel string) (int64, error)
}
