package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/ports"
)

type riskScoringService struct {
	repo      ports.RiskScoringRepository
	publisher ports.EventPublisher
}

func NewRiskScoringService(repo ports.RiskScoringRepository, publisher ports.EventPublisher) ports.RiskScoringService {
	return &riskScoringService{repo: repo, publisher: publisher}
}

func (s *riskScoringService) Create(ctx context.Context, req dto.CreateRiskScoreRequest) (*dto.RiskScoreResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *riskScoringService) GetByID(ctx context.Context, id string) (*dto.RiskScoreResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *riskScoringService) Update(ctx context.Context, req dto.UpdateRiskScoreRequest) (*dto.RiskScoreResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *riskScoringService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
