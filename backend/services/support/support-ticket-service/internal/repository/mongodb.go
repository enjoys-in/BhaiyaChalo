package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/ports"
)

type supportTicketRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.SupportTicketRepository {
	return &supportTicketRepository{db: db}
}

func (r *supportTicketRepository) Create(ctx context.Context, entity *model.Ticket) error {
	// TODO: implement
	return nil
}

func (r *supportTicketRepository) FindByID(ctx context.Context, id string) (*model.Ticket, error) {
	// TODO: implement
	return nil, nil
}

func (r *supportTicketRepository) Update(ctx context.Context, entity *model.Ticket) error {
	// TODO: implement
	return nil
}

func (r *supportTicketRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
