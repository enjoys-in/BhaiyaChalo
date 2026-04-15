package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/engagement/rating-service/internal/ports"
)

type ratingRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.RatingRepository {
	return &ratingRepository{db: db}
}

func (r *ratingRepository) Create(ctx context.Context, entity *model.Rating) error {
	// TODO: implement
	return nil
}

func (r *ratingRepository) FindByID(ctx context.Context, id string) (*model.Rating, error) {
	// TODO: implement
	return nil, nil
}

func (r *ratingRepository) Update(ctx context.Context, entity *model.Rating) error {
	// TODO: implement
	return nil
}

func (r *ratingRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
