package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/ports"
)

type templateService struct {
	repo      ports.TemplateRepository
	publisher ports.EventPublisher
}

func NewTemplateService(repo ports.TemplateRepository, publisher ports.EventPublisher) ports.TemplateService {
	return &templateService{repo: repo, publisher: publisher}
}

func (s *templateService) Create(ctx context.Context, req dto.CreateTemplateRequest) (*dto.TemplateResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *templateService) GetByID(ctx context.Context, id string) (*dto.TemplateResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *templateService) Update(ctx context.Context, req dto.UpdateTemplateRequest) (*dto.TemplateResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *templateService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
