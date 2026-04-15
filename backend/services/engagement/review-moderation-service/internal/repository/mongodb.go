package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/review-moderation-service/internal/ports"
)

type reviewModerationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.ReviewModerationRepository {
	return &reviewModerationRepository{db: db}
}

func (r *reviewModerationRepository) Create(ctx context.Context, entity *model.Review) error {
	// TODO: implement
	return nil
}

func (r *reviewModerationRepository) FindByID(ctx context.Context, id string) (*model.Review, error) {
	// TODO: implement
	return nil, nil
}

func (r *reviewModerationRepository) Update(ctx context.Context, entity *model.Review) error {
	// TODO: implement
	return nil
}

func (r *reviewModerationRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
