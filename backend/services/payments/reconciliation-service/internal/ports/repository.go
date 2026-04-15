package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/model"
)

type ReconciliationRepository interface {
	Create(ctx context.Context, entity *model.Reconciliation) error
	FindByID(ctx context.Context, id string) (*model.Reconciliation, error)
	Update(ctx context.Context, entity *model.Reconciliation) error
	Delete(ctx context.Context, id string) error
}
