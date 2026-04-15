package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/notification-service/internal/ports"
)

type notificationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) Create(ctx context.Context, entity *model.Notification) error {
	// TODO: implement
	return nil
}

func (r *notificationRepository) FindByID(ctx context.Context, id string) (*model.Notification, error) {
	// TODO: implement
	return nil, nil
}

func (r *notificationRepository) Update(ctx context.Context, entity *model.Notification) error {
	// TODO: implement
	return nil
}

func (r *notificationRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
