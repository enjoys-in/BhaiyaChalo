package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/model"
)

type ETAService interface {
	Calculate(ctx context.Context, req *dto.CalculateETARequest) (*model.ETAResult, error)
}
