package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/model"
)

type EventPublisher interface {
	PublishSurgeUpdated(ctx context.Context, zone *model.SurgeZone) error
	PublishSurgePolicyChanged(ctx context.Context, policy *model.SurgePolicy) error
}
