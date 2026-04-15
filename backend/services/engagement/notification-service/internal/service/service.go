package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/ports"
)

type notificationService struct {
	repo      ports.NotificationRepository
	publisher ports.EventPublisher
}

func NewNotificationService(repo ports.NotificationRepository, publisher ports.EventPublisher) ports.NotificationService {
	return &notificationService{repo: repo, publisher: publisher}
}

func (s *notificationService) Create(ctx context.Context, req dto.CreateNotificationRequest) (*dto.NotificationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *notificationService) GetByID(ctx context.Context, id string) (*dto.NotificationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *notificationService) Update(ctx context.Context, req dto.UpdateNotificationRequest) (*dto.NotificationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *notificationService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
