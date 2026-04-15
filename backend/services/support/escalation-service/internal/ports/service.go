package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/dto"
)

type EscalationService interface {
	Create(ctx context.Context, req dto.CreateEscalationRequest) (*dto.EscalationResponse, error)
	GetByID(ctx context.Context, id string) (*dto.EscalationResponse, error)
	Update(ctx context.Context, req dto.UpdateEscalationRequest) (*dto.EscalationResponse, error)
	Delete(ctx context.Context, id string) error
}
