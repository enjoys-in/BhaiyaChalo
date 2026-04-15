package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/model"
)

type SupportTicketRepository interface {
	Create(ctx context.Context, entity *model.Ticket) error
	FindByID(ctx context.Context, id string) (*model.Ticket, error)
	Update(ctx context.Context, entity *model.Ticket) error
	Delete(ctx context.Context, id string) error
}
