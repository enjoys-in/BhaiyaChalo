package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/template-service/internal/model"
)

type EventPublisher interface {
	PublishTemplateCreated(ctx context.Context, entity *model.Template) error
	PublishTemplateUpdated(ctx context.Context, entity *model.Template) error
}
