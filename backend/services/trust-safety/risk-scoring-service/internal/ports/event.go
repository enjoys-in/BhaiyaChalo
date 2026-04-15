package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/model"
)

type EventPublisher interface {
	PublishRiskScored(ctx context.Context, entity *model.RiskScore) error
	PublishRiskFlagged(ctx context.Context, entity *model.RiskScore) error
}
