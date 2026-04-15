package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/ports"
)

type escalationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.EscalationRepository {
	return &escalationRepository{db: db}
}

func (r *escalationRepository) Create(ctx context.Context, entity *model.Escalation) error {
	// TODO: implement
	return nil
}

func (r *escalationRepository) FindByID(ctx context.Context, id string) (*model.Escalation, error) {
	// TODO: implement
	return nil, nil
}

func (r *escalationRepository) Update(ctx context.Context, entity *model.Escalation) error {
	// TODO: implement
	return nil
}

func (r *escalationRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
