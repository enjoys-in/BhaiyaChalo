package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/model"
)

type EventPublisher interface {
	PublishReconciliationStarted(ctx context.Context, entity *model.Reconciliation) error
	PublishReconciliationCompleted(ctx context.Context, entity *model.Reconciliation) error
}
