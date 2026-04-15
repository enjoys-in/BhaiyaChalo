package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/model"
)

type ConnectionRepository interface {
	Register(ctx context.Context, conn *model.Connection) error
	Remove(ctx context.Context, connectionID string) error
	FindByUserID(ctx context.Context, userID string) (*model.Connection, error)
	FindByNode(ctx context.Context, nodeID string) ([]*model.Connection, error)
	GetNodeStatus(ctx context.Context, nodeID string) (*model.NodeStatus, error)
}
