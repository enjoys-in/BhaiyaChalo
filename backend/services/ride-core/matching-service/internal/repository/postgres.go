package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.MatchRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) SaveRequest(ctx context.Context, req *model.MatchRequest) error {
	query := `
		INSERT INTO match_requests (id, booking_id, city_id, pickup_lat, pickup_lng, vehicle_type, radius_km, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.db.ExecContext(ctx, query,
		req.ID, req.BookingID, req.CityID,
		req.PickupLat, req.PickupLng,
		req.VehicleType, req.RadiusKM,
		req.Status, req.CreatedAt,
	)
	return err
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, id string, status model.MatchStatus) error {
	query := `UPDATE match_requests SET status = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *postgresRepository) FindByBookingID(ctx context.Context, bookingID string) (*model.MatchRequest, error) {
	query := `
		SELECT id, booking_id, city_id, pickup_lat, pickup_lng, vehicle_type, radius_km, status, created_at
		FROM match_requests WHERE booking_id = $1 ORDER BY created_at DESC LIMIT 1`

	mr := &model.MatchRequest{}
	err := r.db.QueryRowContext(ctx, query, bookingID).Scan(
		&mr.ID, &mr.BookingID, &mr.CityID,
		&mr.PickupLat, &mr.PickupLng,
		&mr.VehicleType, &mr.RadiusKM,
		&mr.Status, &mr.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return mr, nil
}
