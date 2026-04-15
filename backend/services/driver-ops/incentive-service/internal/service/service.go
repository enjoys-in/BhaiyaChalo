package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/incentive-service/internal/ports"
)

type incentiveService struct {
	repo      ports.IncentiveRepository
	publisher ports.EventPublisher
}

func NewIncentiveService(repo ports.IncentiveRepository, publisher ports.EventPublisher) ports.IncentiveService {
	return &incentiveService{repo: repo, publisher: publisher}
}

func (s *incentiveService) Create(ctx context.Context, req dto.CreateIncentiveRequest) (*dto.IncentiveResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *incentiveService) GetByID(ctx context.Context, id string) (*dto.IncentiveResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *incentiveService) Update(ctx context.Context, req dto.UpdateIncentiveRequest) (*dto.IncentiveResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *incentiveService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
