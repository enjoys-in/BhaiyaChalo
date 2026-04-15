package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/trust-safety/fraud-detection-service/internal/ports"
)

type fraudDetectionService struct {
	repo      ports.FraudDetectionRepository
	publisher ports.EventPublisher
}

func NewFraudDetectionService(repo ports.FraudDetectionRepository, publisher ports.EventPublisher) ports.FraudDetectionService {
	return &fraudDetectionService{repo: repo, publisher: publisher}
}

func (s *fraudDetectionService) Create(ctx context.Context, req dto.CreateFraudSignalRequest) (*dto.FraudSignalResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *fraudDetectionService) GetByID(ctx context.Context, id string) (*dto.FraudSignalResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *fraudDetectionService) Update(ctx context.Context, req dto.UpdateFraudSignalRequest) (*dto.FraudSignalResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *fraudDetectionService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
