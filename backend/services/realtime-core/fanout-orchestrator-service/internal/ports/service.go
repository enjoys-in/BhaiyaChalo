package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/fanout-orchestrator-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/fanout-orchestrator-service/internal/model"
)

type FanoutService interface {
	Send(ctx context.Context, req dto.SendMessageRequest) (*model.FanoutMessage, error)
	Broadcast(ctx context.Context, req dto.BroadcastRequest) (*model.FanoutMessage, error)
	GetStats(ctx context.Context, messageID string) (*dto.DeliveryStatsResponse, error)
}
