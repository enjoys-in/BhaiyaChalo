package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/availability-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.AvailabilityRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) SetOnline(ctx context.Context, avail *model.DriverAvailability) error {
	query := `
		INSERT INTO driver_availability (driver_id, city_id, online, on_trip, vehicle_type, lat, lng, last_seen_at, updated_at)
		VALUES ($1, $2, true, false, $3, $4, $5, $6, $7)
		ON CONFLICT (driver_id) DO UPDATE SET
			city_id = $2, online = true, on_trip = false, vehicle_type = $3,
			lat = $4, lng = $5, last_seen_at = $6, updated_at = $7`

	_, err := r.db.ExecContext(ctx, query,
		avail.DriverID, avail.CityID, avail.VehicleType,
		avail.Lat, avail.Lng, avail.LastSeenAt, avail.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) SetOffline(ctx context.Context, driverID string) error {
	query := `UPDATE driver_availability SET online = false, on_trip = false, updated_at = NOW() WHERE driver_id = $1`
	_, err := r.db.ExecContext(ctx, query, driverID)
	return err
}

func (r *postgresRepository) SetOnTrip(ctx context.Context, driverID string) error {
	query := `UPDATE driver_availability SET on_trip = true, updated_at = NOW() WHERE driver_id = $1`
	_, err := r.db.ExecContext(ctx, query, driverID)
	return err
}

func (r *postgresRepository) SetFree(ctx context.Context, driverID string) error {
	query := `UPDATE driver_availability SET on_trip = false, updated_at = NOW() WHERE driver_id = $1`
	_, err := r.db.ExecContext(ctx, query, driverID)
	return err
}

func (r *postgresRepository) GetStatus(ctx context.Context, driverID string) (*model.DriverAvailability, error) {
	query := `
		SELECT driver_id, city_id, online, on_trip, vehicle_type, lat, lng, last_seen_at, updated_at
		FROM driver_availability WHERE driver_id = $1`

	a := &model.DriverAvailability{}
	err := r.db.QueryRowContext(ctx, query, driverID).Scan(
		&a.DriverID, &a.CityID, &a.Online, &a.OnTrip,
		&a.VehicleType, &a.Lat, &a.Lng, &a.LastSeenAt, &a.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("scan availability: %w", err)
	}
	return a, nil
}

func (r *postgresRepository) CountOnlineByCityAndType(ctx context.Context, cityID, vehicleType string) (int, error) {
	query := `SELECT COUNT(*) FROM driver_availability WHERE city_id = $1 AND vehicle_type = $2 AND online = true AND on_trip = false`
	var count int
	err := r.db.QueryRowContext(ctx, query, cityID, vehicleType).Scan(&count)
	return count, err
}

func (r *postgresRepository) LogAction(ctx context.Context, log *model.AvailabilityLog) error {
	query := `INSERT INTO availability_logs (id, driver_id, action, timestamp) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, log.ID, log.DriverID, log.Action, log.Timestamp)
	return err
}
