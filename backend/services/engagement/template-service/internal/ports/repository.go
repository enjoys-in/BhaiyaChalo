package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/model"
)

type TemplateRepository interface {
	Create(ctx context.Context, entity *model.Template) error
	FindByID(ctx context.Context, id string) (*model.Template, error)
	Update(ctx context.Context, entity *model.Template) error
	Delete(ctx context.Context, id string) error
}
