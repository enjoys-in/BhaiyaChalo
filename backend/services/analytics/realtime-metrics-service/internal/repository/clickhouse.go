package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/realtime-metrics-service/internal/ports"
)

type realtimeMetricsRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.RealtimeMetricsRepository {
	return &realtimeMetricsRepository{db: db}
}

func (r *realtimeMetricsRepository) Create(ctx context.Context, entity *model.Metric) error {
	// TODO: implement
	return nil
}

func (r *realtimeMetricsRepository) FindByID(ctx context.Context, id string) (*model.Metric, error) {
	// TODO: implement
	return nil, nil
}

func (r *realtimeMetricsRepository) Update(ctx context.Context, entity *model.Metric) error {
	// TODO: implement
	return nil
}

func (r *realtimeMetricsRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
