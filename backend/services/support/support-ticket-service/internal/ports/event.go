package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/model"
)

type EventPublisher interface {
	PublishTicketCreated(ctx context.Context, entity *model.Ticket) error
	PublishTicketResolved(ctx context.Context, entity *model.Ticket) error
}
