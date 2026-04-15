package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/ports"
)

type incentiveRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.IncentiveRepository {
	return &incentiveRepository{db: db}
}

func (r *incentiveRepository) Create(ctx context.Context, entity *model.Incentive) error {
	// TODO: implement
	return nil
}

func (r *incentiveRepository) FindByID(ctx context.Context, id string) (*model.Incentive, error) {
	// TODO: implement
	return nil, nil
}

func (r *incentiveRepository) Update(ctx context.Context, entity *model.Incentive) error {
	// TODO: implement
	return nil
}

func (r *incentiveRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
