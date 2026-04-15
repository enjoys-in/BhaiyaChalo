package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/model"
)

type EscalationRepository interface {
	Create(ctx context.Context, entity *model.Escalation) error
	FindByID(ctx context.Context, id string) (*model.Escalation, error)
	Update(ctx context.Context, entity *model.Escalation) error
	Delete(ctx context.Context, id string) error
}
