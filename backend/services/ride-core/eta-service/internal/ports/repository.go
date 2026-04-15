package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/model"
)

type ETARepository interface {
	SaveCalculation(ctx context.Context, req *model.ETARequest, result *model.ETAResult) error
	FindRecent(ctx context.Context, req *model.ETARequest) (*model.ETAResult, error)
}
