package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/model"
)

type EventPublisher interface {
	PublishFraudDetected(ctx context.Context, entity *model.FraudSignal) error
	PublishFraudCleared(ctx context.Context, entity *model.FraudSignal) error
}
