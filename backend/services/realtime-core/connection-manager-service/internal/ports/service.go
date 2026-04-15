package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/model"
)

type ConnectionManager interface {
	Register(ctx context.Context, req dto.RegisterConnectionRequest) (*model.Connection, error)
	Remove(ctx context.Context, connectionID string) error
	LocateUser(ctx context.Context, userID string) (*model.Connection, error)
	GetNodeHealth(ctx context.Context, nodeID string) (*model.NodeStatus, error)
}
