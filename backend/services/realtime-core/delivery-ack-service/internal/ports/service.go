package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/dto"
)

type DeliveryAckService interface {
	Create(ctx context.Context, req dto.CreateDeliveryAckRequest) (*dto.DeliveryAckResponse, error)
	GetByID(ctx context.Context, id string) (*dto.DeliveryAckResponse, error)
	Update(ctx context.Context, req dto.UpdateDeliveryAckRequest) (*dto.DeliveryAckResponse, error)
	Delete(ctx context.Context, id string) error
}
