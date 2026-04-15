package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/dto"
)

type SupportTicketService interface {
	Create(ctx context.Context, req dto.CreateTicketRequest) (*dto.TicketResponse, error)
	GetByID(ctx context.Context, id string) (*dto.TicketResponse, error)
	Update(ctx context.Context, req dto.UpdateTicketRequest) (*dto.TicketResponse, error)
	Delete(ctx context.Context, id string) error
}
