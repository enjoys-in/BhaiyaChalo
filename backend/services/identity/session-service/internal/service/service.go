package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/session-service/internal/ports"
)

const defaultSessionTTL = 24 * time.Hour

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

func (s *sessionService) Create(ctx context.Context, req *dto.CreateSessionRequest) (*dto.SessionResponse, error) {
	now := time.Now().UTC()
	session := &model.Session{
		ID:        uuid.NewString(),
		UserID:    req.UserID,
		Role:      req.Role,
		DeviceID:  req.DeviceID,
		IP:        req.IP,
		UserAgent: req.UserAgent,
		Active:    true,
		ExpiresAt: now.Add(defaultSessionTTL),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(ctx, session); err != nil {
		return nil, err
	}

	_ = s.publisher.PublishSessionCreated(ctx, session)

	return toSessionResponse(session), nil
}

func (s *sessionService) GetByID(ctx context.Context, req *dto.GetSessionRequest) (*dto.SessionResponse, error) {
	session, err := s.repo.FindByID(ctx, req.SessionID)
	if err != nil {
		return nil, err
	}
	return toSessionResponse(session), nil
}

func (s *sessionService) Invalidate(ctx context.Context, req *dto.InvalidateRequest) error {
	session, err := s.repo.FindByID(ctx, req.SessionID)
	if err != nil {
		return err
	}

	if err := s.repo.Invalidate(ctx, req.SessionID); err != nil {
		return err
	}

	_ = s.publisher.PublishSessionInvalidated(ctx, session.ID, session.UserID)
	return nil
}

func (s *sessionService) InvalidateAll(ctx context.Context, req *dto.InvalidateAllRequest) error {
	if err := s.repo.InvalidateAllByUser(ctx, req.UserID); err != nil {
		return err
	}

	_ = s.publisher.PublishSessionInvalidated(ctx, "", req.UserID)
	return nil
}

func toSessionResponse(session *model.Session) *dto.SessionResponse {
	return &dto.SessionResponse{
		ID:        session.ID,
		UserID:    session.UserID,
		Role:      session.Role,
		DeviceID:  session.DeviceID,
		Active:    session.Active,
		ExpiresAt: session.ExpiresAt,
	}
}
