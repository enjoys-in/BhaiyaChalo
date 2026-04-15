package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/support/support-ticket-service/internal/ports"
)

type supportTicketService struct {
	repo      ports.SupportTicketRepository
	publisher ports.EventPublisher
}

func NewSupportTicketService(repo ports.SupportTicketRepository, publisher ports.EventPublisher) ports.SupportTicketService {
	return &supportTicketService{repo: repo, publisher: publisher}
}

func (s *supportTicketService) Create(ctx context.Context, req dto.CreateTicketRequest) (*dto.TicketResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *supportTicketService) GetByID(ctx context.Context, id string) (*dto.TicketResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *supportTicketService) Update(ctx context.Context, req dto.UpdateTicketRequest) (*dto.TicketResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *supportTicketService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
