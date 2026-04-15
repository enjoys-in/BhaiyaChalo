package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/model"
)

type DeliveryAckRepository interface {
	Create(ctx context.Context, entity *model.DeliveryAck) error
	FindByID(ctx context.Context, id string) (*model.DeliveryAck, error)
	Update(ctx context.Context, entity *model.DeliveryAck) error
	Delete(ctx context.Context, id string) error
}
