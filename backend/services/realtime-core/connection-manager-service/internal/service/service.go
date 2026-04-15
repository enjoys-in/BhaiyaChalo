package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/connection-manager-service/internal/ports"
)

type connectionManager struct {
	repo     ports.ConnectionRepository
	capacity int
}

func NewConnectionManager(repo ports.ConnectionRepository, maxConnsPerNode int) ports.ConnectionManager {
	return &connectionManager{
		repo:     repo,
		capacity: maxConnsPerNode,
	}
}

func (s *connectionManager) Register(ctx context.Context, req dto.RegisterConnectionRequest) (*model.Connection, error) {
	conn := &model.Connection{
		ID:         uuid.New().String(),
		UserID:     req.UserID,
		ServerNode: req.ServerNode,
		Protocol:   req.Protocol,
		CreatedAt:  time.Now().UTC(),
	}

	if err := s.repo.Register(ctx, conn); err != nil {
		return nil, err
	}

	return conn, nil
}

func (s *connectionManager) Remove(ctx context.Context, connectionID string) error {
	return s.repo.Remove(ctx, connectionID)
}

func (s *connectionManager) LocateUser(ctx context.Context, userID string) (*model.Connection, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *connectionManager) GetNodeHealth(ctx context.Context, nodeID string) (*model.NodeStatus, error) {
	status, err := s.repo.GetNodeStatus(ctx, nodeID)
	if err != nil {
		return nil, err
	}
	status.Capacity = s.capacity
	return status, nil
}
