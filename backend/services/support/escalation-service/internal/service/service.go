package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/ports"
)

type escalationService struct {
	repo      ports.EscalationRepository
	publisher ports.EventPublisher
}

func NewEscalationService(repo ports.EscalationRepository, publisher ports.EventPublisher) ports.EscalationService {
	return &escalationService{repo: repo, publisher: publisher}
}

func (s *escalationService) Create(ctx context.Context, req dto.CreateEscalationRequest) (*dto.EscalationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *escalationService) GetByID(ctx context.Context, id string) (*dto.EscalationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *escalationService) Update(ctx context.Context, req dto.UpdateEscalationRequest) (*dto.EscalationResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *escalationService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
