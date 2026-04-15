package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/model"
)

type SurgeRepository interface {
	GetPolicy(ctx context.Context, cityID string) (*model.SurgePolicy, error)
	UpdatePolicy(ctx context.Context, policy *model.SurgePolicy) error
	SaveZone(ctx context.Context, zone *model.SurgeZone) error
	GetZone(ctx context.Context, zoneID string) (*model.SurgeZone, error)
	SaveHistory(ctx context.Context, history *model.SurgeHistory) error
	GetHistory(ctx context.Context, zoneID string, limit int) ([]*model.SurgeHistory, error)
}
