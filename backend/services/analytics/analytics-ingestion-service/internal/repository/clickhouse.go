package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-ingestion-service/internal/ports"
)

type analyticsIngestionRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.AnalyticsIngestionRepository {
	return &analyticsIngestionRepository{db: db}
}

func (r *analyticsIngestionRepository) Create(ctx context.Context, entity *model.AnalyticsEvent) error {
	// TODO: implement
	return nil
}

func (r *analyticsIngestionRepository) FindByID(ctx context.Context, id string) (*model.AnalyticsEvent, error) {
	// TODO: implement
	return nil, nil
}

func (r *analyticsIngestionRepository) Update(ctx context.Context, entity *model.AnalyticsEvent) error {
	// TODO: implement
	return nil
}

func (r *analyticsIngestionRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
