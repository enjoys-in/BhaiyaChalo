package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/dto"
)

type TemplateService interface {
	Create(ctx context.Context, req dto.CreateTemplateRequest) (*dto.TemplateResponse, error)
	GetByID(ctx context.Context, id string) (*dto.TemplateResponse, error)
	Update(ctx context.Context, req dto.UpdateTemplateRequest) (*dto.TemplateResponse, error)
	Delete(ctx context.Context, id string) error
}
