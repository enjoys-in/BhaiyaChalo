package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/dto"
)

type ReconciliationService interface {
	Create(ctx context.Context, req dto.CreateReconciliationRequest) (*dto.ReconciliationResponse, error)
	GetByID(ctx context.Context, id string) (*dto.ReconciliationResponse, error)
	Update(ctx context.Context, req dto.UpdateReconciliationRequest) (*dto.ReconciliationResponse, error)
	Delete(ctx context.Context, id string) error
}
