package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/model"
)

type RiskScoringRepository interface {
	Create(ctx context.Context, entity *model.RiskScore) error
	FindByID(ctx context.Context, id string) (*model.RiskScore, error)
	Update(ctx context.Context, entity *model.RiskScore) error
	Delete(ctx context.Context, id string) error
}
