package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/location-ingest-service/internal/ports"
)

type locationIngestRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ports.LocationIngestRepository {
	return &locationIngestRepository{db: db}
}

func (r *locationIngestRepository) Create(ctx context.Context, entity *model.LocationUpdate) error {
	// TODO: implement
	return nil
}

func (r *locationIngestRepository) FindByID(ctx context.Context, id string) (*model.LocationUpdate, error) {
	// TODO: implement
	return nil, nil
}

func (r *locationIngestRepository) Update(ctx context.Context, entity *model.LocationUpdate) error {
	// TODO: implement
	return nil
}

func (r *locationIngestRepository) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
