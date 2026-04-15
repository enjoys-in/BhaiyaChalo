package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/subscription-service/internal/ports"
)

type subscriptionService struct {
	repo ports.SubscriptionRepository
}

func NewSubscriptionService(repo ports.SubscriptionRepository) ports.SubscriptionService {
	return &subscriptionService{repo: repo}
}

func (s *subscriptionService) Subscribe(ctx context.Context, req dto.SubscribeRequest) (*model.Subscription, error) {
	sub := &model.Subscription{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		Channel:   model.Channel(req.Channel),
		Topic:     req.Topic,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.Subscribe(ctx, sub); err != nil {
		return nil, err
	}
	return sub, nil
}

func (s *subscriptionService) Unsubscribe(ctx context.Context, req dto.UnsubscribeRequest) error {
	return s.repo.Unsubscribe(ctx, req.SubscriptionID)
}

func (s *subscriptionService) FindByUser(ctx context.Context, userID string) ([]*model.Subscription, error) {
	return s.repo.FindByUser(ctx, userID)
}

func (s *subscriptionService) FindByChannel(ctx context.Context, channel string, topic string) ([]*model.Subscription, error) {
	return s.repo.FindByChannel(ctx, model.Channel(channel), topic)
}

func (s *subscriptionService) CountByChannel(ctx context.Context, channel string) (int64, error) {
	return s.repo.CountByChannel(ctx, model.Channel(channel))
}
