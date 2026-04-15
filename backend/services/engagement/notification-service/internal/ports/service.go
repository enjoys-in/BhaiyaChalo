package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/dto"
)

type NotificationService interface {
	Create(ctx context.Context, req dto.CreateNotificationRequest) (*dto.NotificationResponse, error)
	GetByID(ctx context.Context, id string) (*dto.NotificationResponse, error)
	Update(ctx context.Context, req dto.UpdateNotificationRequest) (*dto.NotificationResponse, error)
	Delete(ctx context.Context, id string) error
}
