package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/ports"
)

type deliveryAckRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.DeliveryAckRepository {
	return &deliveryAckRepository{db: db}
}

func (r *deliveryAckRepository) Create(ctx context.Context, entity *model.DeliveryAck) error {
	// TODO: implement
	return nil
}

func (r *deliveryAckRepository) FindByID(ctx context.Context, id string) (*model.DeliveryAck, error) {
	// TODO: implement
	return nil, nil
}

func (r *deliveryAckRepository) Update(ctx context.Context, entity *model.DeliveryAck) error {
	// TODO: implement
	return nil
}

func (r *deliveryAckRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
