package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/ports"
)

type reconciliationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.ReconciliationRepository {
	return &reconciliationRepository{db: db}
}

func (r *reconciliationRepository) Create(ctx context.Context, entity *model.Reconciliation) error {
	// TODO: implement
	return nil
}

func (r *reconciliationRepository) FindByID(ctx context.Context, id string) (*model.Reconciliation, error) {
	// TODO: implement
	return nil, nil
}

func (r *reconciliationRepository) Update(ctx context.Context, entity *model.Reconciliation) error {
	// TODO: implement
	return nil
}

func (r *reconciliationRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
