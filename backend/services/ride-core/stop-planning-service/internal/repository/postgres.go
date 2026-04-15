package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/ports"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.StopRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) AddStop(ctx context.Context, stop *model.Stop) error {
	query := `INSERT INTO stops (id, trip_id, lat, lng, address, stop_order, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.ExecContext(ctx, query,
		stop.ID, stop.TripID, stop.Lat, stop.Lng, stop.Address,
		stop.StopOrder, stop.Status, stop.CreatedAt, stop.UpdatedAt,
	)
	return err
}

func (r *postgresRepo) RemoveStop(ctx context.Context, tripID, stopID string) error {
	query := `DELETE FROM stops WHERE id = $1 AND trip_id = $2`
	result, err := r.db.ExecContext(ctx, query, stopID, tripID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("stop not found")
	}
	return nil
}

func (r *postgresRepo) ReorderStops(ctx context.Context, tripID string, stopIDs []string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for i, id := range stopIDs {
		query := `UPDATE stops SET stop_order = $1, updated_at = $2 WHERE id = $3 AND trip_id = $4`
		_, err := tx.ExecContext(ctx, query, i, time.Now().UTC(), id, tripID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *postgresRepo) UpdateStopStatus(ctx context.Context, tripID, stopID string, status model.StopStatus) (*model.Stop, error) {
	now := time.Now().UTC()
	var query string

	switch status {
	case model.StopStatusArrived:
		query = `UPDATE stops SET status = $1, arrived_at = $2, updated_at = $2 WHERE id = $3 AND trip_id = $4
			RETURNING id, trip_id, lat, lng, address, stop_order, status, arrived_at, departed_at, created_at, updated_at`
	case model.StopStatusCompleted:
		query = `UPDATE stops SET status = $1, departed_at = $2, updated_at = $2 WHERE id = $3 AND trip_id = $4
			RETURNING id, trip_id, lat, lng, address, stop_order, status, arrived_at, departed_at, created_at, updated_at`
	default:
		query = `UPDATE stops SET status = $1, updated_at = $2 WHERE id = $3 AND trip_id = $4
			RETURNING id, trip_id, lat, lng, address, stop_order, status, arrived_at, departed_at, created_at, updated_at`
	}

	var stop model.Stop
	err := r.db.QueryRowContext(ctx, query, status, now, stopID, tripID).Scan(
		&stop.ID, &stop.TripID, &stop.Lat, &stop.Lng, &stop.Address,
		&stop.StopOrder, &stop.Status, &stop.ArrivedAt, &stop.DepartedAt,
		&stop.CreatedAt, &stop.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &stop, nil
}

func (r *postgresRepo) FindByTripID(ctx context.Context, tripID string) (*model.MultiStopTrip, error) {
	query := `SELECT id, trip_id, lat, lng, address, stop_order, status, arrived_at, departed_at, created_at, updated_at
		FROM stops WHERE trip_id = $1 ORDER BY stop_order`
	rows, err := r.db.QueryContext(ctx, query, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stops []model.Stop
	for rows.Next() {
		var s model.Stop
		if err := rows.Scan(
			&s.ID, &s.TripID, &s.Lat, &s.Lng, &s.Address,
			&s.StopOrder, &s.Status, &s.ArrivedAt, &s.DepartedAt,
			&s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		stops = append(stops, s)
	}

	return &model.MultiStopTrip{
		TripID:     tripID,
		Stops:      stops,
		TotalStops: len(stops),
	}, nil
}
