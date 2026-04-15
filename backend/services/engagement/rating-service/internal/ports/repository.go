package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/model"
)

type RatingRepository interface {
	Create(ctx context.Context, entity *model.Rating) error
	FindByID(ctx context.Context, id string) (*model.Rating, error)
	Update(ctx context.Context, entity *model.Rating) error
	Delete(ctx context.Context, id string) error
}
