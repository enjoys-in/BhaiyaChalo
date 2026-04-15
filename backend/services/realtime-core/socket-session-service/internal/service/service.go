package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/ports"
)

type sessionService struct {
	repo      ports.SessionRepository
	publisher ports.EventPublisher
}

func NewSessionService(repo ports.SessionRepository, publisher ports.EventPublisher) ports.SessionService {
	return &sessionService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *sessionService) Register(ctx context.Context, req dto.RegisterSessionRequest) (*model.SocketSession, error) {
	session := &model.SocketSession{
		ID:          uuid.New().String(),
		UserID:      req.UserID,
		Role:        req.Role,
		DeviceID:    req.DeviceID,
		ServerID:    req.ServerID,
		ConnectedAt: time.Now().UTC(),
		Active:      true,
	}

	if err := s.repo.Register(ctx, session); err != nil {
		return nil, err
	}

	_ = s.publisher.PublishSessionConnected(ctx, session)

	return session, nil
}

func (s *sessionService) Unregister(ctx context.Context, req dto.UnregisterSessionRequest) error {
	sessions, err := s.repo.FindByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}

	var target *model.SocketSession
	for _, sess := range sessions {
		if sess.ID == req.SessionID {
			target = sess
			break
		}
	}

	if err := s.repo.Unregister(ctx, req.SessionID); err != nil {
		return err
	}

	if target != nil {
		_ = s.publisher.PublishSessionDisconnected(ctx, target)
	}

	return nil
}

func (s *sessionService) FindByUserID(ctx context.Context, userID string) ([]*model.SocketSession, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *sessionService) FindActiveByServer(ctx context.Context, serverID string) ([]*model.SocketSession, error) {
	return s.repo.FindActiveByServer(ctx, serverID)
}

func (s *sessionService) CountActive(ctx context.Context) (int64, error) {
	return s.repo.CountActive(ctx)
}
