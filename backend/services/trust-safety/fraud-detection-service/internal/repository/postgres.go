package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/ports"
)

type fraudDetectionRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.FraudDetectionRepository {
	return &fraudDetectionRepository{db: db}
}

func (r *fraudDetectionRepository) Create(ctx context.Context, entity *model.FraudSignal) error {
	// TODO: implement
	return nil
}

func (r *fraudDetectionRepository) FindByID(ctx context.Context, id string) (*model.FraudSignal, error) {
	// TODO: implement
	return nil, nil
}

func (r *fraudDetectionRepository) Update(ctx context.Context, entity *model.FraudSignal) error {
	// TODO: implement
	return nil
}

func (r *fraudDetectionRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
