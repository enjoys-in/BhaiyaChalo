package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/payout-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/payout-service/internal/ports"
)

type payoutService struct {
	repo      ports.PayoutRepository
	publisher ports.EventPublisher
}

func NewPayoutService(repo ports.PayoutRepository, publisher ports.EventPublisher) ports.PayoutService {
	return &payoutService{repo: repo, publisher: publisher}
}

func (s *payoutService) Create(ctx context.Context, req dto.CreatePayoutRequest) (*dto.PayoutResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *payoutService) GetByID(ctx context.Context, id string) (*dto.PayoutResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *payoutService) Update(ctx context.Context, req dto.UpdatePayoutRequest) (*dto.PayoutResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *payoutService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
