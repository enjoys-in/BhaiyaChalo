package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/model"
)

type EventPublisher interface {
	PublishIncentiveCreated(ctx context.Context, entity *model.Incentive) error
	PublishIncentiveCompleted(ctx context.Context, entity *model.Incentive) error
}
