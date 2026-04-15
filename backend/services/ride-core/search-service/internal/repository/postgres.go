package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.SearchRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) SaveQuery(ctx context.Context, query *model.SearchQuery) error {
	q := `
		INSERT INTO search_queries (id, user_id, city_id, pickup_lat, pickup_lng, drop_lat, drop_lng, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(ctx, q,
		query.ID, query.UserID, query.CityID,
		query.PickupLat, query.PickupLng,
		query.DropLat, query.DropLng,
		query.CreatedAt,
	)
	return err
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*model.SearchQuery, error) {
	q := `
		SELECT id, user_id, city_id, pickup_lat, pickup_lng, drop_lat, drop_lng, created_at
		FROM search_queries WHERE id = $1`

	sq := &model.SearchQuery{}
	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&sq.ID, &sq.UserID, &sq.CityID,
		&sq.PickupLat, &sq.PickupLng,
		&sq.DropLat, &sq.DropLng,
		&sq.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return sq, nil
}
