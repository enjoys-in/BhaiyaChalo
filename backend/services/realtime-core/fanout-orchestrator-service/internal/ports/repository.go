package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/fanout-orchestrator-service/internal/model"
)

type FanoutRepository interface {
	SaveMessage(ctx context.Context, msg *model.FanoutMessage) error
	GetDeliveryStats(ctx context.Context, messageID string) (total, delivered, failed int64, err error)
}
