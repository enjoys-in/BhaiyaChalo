package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/dto"
)

type SurgeService interface {
	Calculate(ctx context.Context, req dto.CalculateSurgeRequest) (*dto.SurgeResponse, error)
	GetCurrentSurge(ctx context.Context, zoneID string) (*dto.SurgeResponse, error)
	UpdatePolicy(ctx context.Context, cityID string, req dto.UpdatePolicyRequest) (*dto.SurgePolicyResponse, error)
}
