package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/dto"
)

type RiskScoringService interface {
	Create(ctx context.Context, req dto.CreateRiskScoreRequest) (*dto.RiskScoreResponse, error)
	GetByID(ctx context.Context, id string) (*dto.RiskScoreResponse, error)
	Update(ctx context.Context, req dto.UpdateRiskScoreRequest) (*dto.RiskScoreResponse, error)
	Delete(ctx context.Context, id string) error
}
