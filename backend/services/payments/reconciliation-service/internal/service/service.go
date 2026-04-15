package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/reconciliation-service/internal/ports"
)

type reconciliationService struct {
	repo      ports.ReconciliationRepository
	publisher ports.EventPublisher
}

func NewReconciliationService(repo ports.ReconciliationRepository, publisher ports.EventPublisher) ports.ReconciliationService {
	return &reconciliationService{repo: repo, publisher: publisher}
}

func (s *reconciliationService) Create(ctx context.Context, req dto.CreateReconciliationRequest) (*dto.ReconciliationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *reconciliationService) GetByID(ctx context.Context, id string) (*dto.ReconciliationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *reconciliationService) Update(ctx context.Context, req dto.UpdateReconciliationRequest) (*dto.ReconciliationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *reconciliationService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
