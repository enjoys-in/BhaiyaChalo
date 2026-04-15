package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/model"
)

type FraudDetectionRepository interface {
	Create(ctx context.Context, entity *model.FraudSignal) error
	FindByID(ctx context.Context, id string) (*model.FraudSignal, error)
	Update(ctx context.Context, entity *model.FraudSignal) error
	Delete(ctx context.Context, id string) error
}
