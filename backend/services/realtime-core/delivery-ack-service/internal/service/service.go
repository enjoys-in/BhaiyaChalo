package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/delivery-ack-service/internal/ports"
)

type deliveryAckService struct {
	repo      ports.DeliveryAckRepository
	publisher ports.EventPublisher
}

func NewDeliveryAckService(repo ports.DeliveryAckRepository, publisher ports.EventPublisher) ports.DeliveryAckService {
	return &deliveryAckService{repo: repo, publisher: publisher}
}

func (s *deliveryAckService) Create(ctx context.Context, req dto.CreateDeliveryAckRequest) (*dto.DeliveryAckResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *deliveryAckService) GetByID(ctx context.Context, id string) (*dto.DeliveryAckResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *deliveryAckService) Update(ctx context.Context, req dto.UpdateDeliveryAckRequest) (*dto.DeliveryAckResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *deliveryAckService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
