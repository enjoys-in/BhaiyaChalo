package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/model"
)

type ReviewModerationRepository interface {
	Create(ctx context.Context, entity *model.Review) error
	FindByID(ctx context.Context, id string) (*model.Review, error)
	Update(ctx context.Context, entity *model.Review) error
	Delete(ctx context.Context, id string) error
}
