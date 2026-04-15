package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/payout-service/internal/model"
)

type EventPublisher interface {
	PublishPayoutInitiated(ctx context.Context, entity *model.Payout) error
	PublishPayoutCompleted(ctx context.Context, entity *model.Payout) error
}
