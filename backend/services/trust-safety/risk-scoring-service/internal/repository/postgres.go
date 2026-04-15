package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/risk-scoring-service/internal/ports"
)

type riskScoringRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.RiskScoringRepository {
	return &riskScoringRepository{db: db}
}

func (r *riskScoringRepository) Create(ctx context.Context, entity *model.RiskScore) error {
	// TODO: implement
	return nil
}

func (r *riskScoringRepository) FindByID(ctx context.Context, id string) (*model.RiskScore, error) {
	// TODO: implement
	return nil, nil
}

func (r *riskScoringRepository) Update(ctx context.Context, entity *model.RiskScore) error {
	// TODO: implement
	return nil
}

func (r *riskScoringRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
