package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/dto"
)

type FraudDetectionService interface {
	Create(ctx context.Context, req dto.CreateFraudSignalRequest) (*dto.FraudSignalResponse, error)
	GetByID(ctx context.Context, id string) (*dto.FraudSignalResponse, error)
	Update(ctx context.Context, req dto.UpdateFraudSignalRequest) (*dto.FraudSignalResponse, error)
	Delete(ctx context.Context, id string) error
}
