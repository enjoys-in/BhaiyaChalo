package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/payout-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/payout-service/internal/ports"
)

type payoutRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.PayoutRepository {
	return &payoutRepository{db: db}
}

func (r *payoutRepository) Create(ctx context.Context, entity *model.Payout) error {
	// TODO: implement
	return nil
}

func (r *payoutRepository) FindByID(ctx context.Context, id string) (*model.Payout, error) {
	// TODO: implement
	return nil, nil
}

func (r *payoutRepository) Update(ctx context.Context, entity *model.Payout) error {
	// TODO: implement
	return nil
}

func (r *payoutRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
