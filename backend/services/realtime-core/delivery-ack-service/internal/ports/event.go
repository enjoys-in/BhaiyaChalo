package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/model"
)

type EventPublisher interface {
	PublishDeliveryAck(ctx context.Context, entity *model.DeliveryAck) error
	PublishDeliveryTimeout(ctx context.Context, entity *model.DeliveryAck) error
}
