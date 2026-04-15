package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/vehicle-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.VehicleRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, vehicle *model.Vehicle) error {
	query := `
		INSERT INTO vehicles (id, driver_id, make, model, year, color, plate_number, vehicle_type,
			insurance_expiry, fitness_expiry, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := r.db.ExecContext(ctx, query,
		vehicle.ID, vehicle.DriverID, vehicle.Make, vehicle.Model,
		vehicle.Year, vehicle.Color, vehicle.PlateNumber, vehicle.VehicleType,
		vehicle.InsuranceExpiry, vehicle.FitnessExpiry, vehicle.Status,
		vehicle.CreatedAt, vehicle.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*model.Vehicle, error) {
	query := `
		SELECT id, driver_id, make, model, year, color, plate_number, vehicle_type,
			insurance_expiry, fitness_expiry, status, created_at, updated_at
		FROM vehicles WHERE id = $1`

	v := &model.Vehicle{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&v.ID, &v.DriverID, &v.Make, &v.Model,
		&v.Year, &v.Color, &v.PlateNumber, &v.VehicleType,
		&v.InsuranceExpiry, &v.FitnessExpiry, &v.Status,
		&v.CreatedAt, &v.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (r *postgresRepository) FindByDriverID(ctx context.Context, driverID string) ([]*model.Vehicle, error) {
	query := `
		SELECT id, driver_id, make, model, year, color, plate_number, vehicle_type,
			insurance_expiry, fitness_expiry, status, created_at, updated_at
		FROM vehicles WHERE driver_id = $1 ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, driverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanVehicles(rows)
}

func (r *postgresRepository) Update(ctx context.Context, vehicle *model.Vehicle) error {
	query := `
		UPDATE vehicles SET make = $1, model = $2, year = $3, color = $4,
			plate_number = $5, vehicle_type = $6, insurance_expiry = $7,
			fitness_expiry = $8, status = $9, updated_at = $10
		WHERE id = $11`

	_, err := r.db.ExecContext(ctx, query,
		vehicle.Make, vehicle.Model, vehicle.Year, vehicle.Color,
		vehicle.PlateNumber, vehicle.VehicleType, vehicle.InsuranceExpiry,
		vehicle.FitnessExpiry, vehicle.Status, vehicle.UpdatedAt,
		vehicle.ID,
	)
	return err
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM vehicles WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *postgresRepository) ListByType(ctx context.Context, vehicleType model.VehicleType) ([]*model.Vehicle, error) {
	query := `
		SELECT id, driver_id, make, model, year, color, plate_number, vehicle_type,
			insurance_expiry, fitness_expiry, status, created_at, updated_at
		FROM vehicles WHERE vehicle_type = $1 ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, vehicleType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanVehicles(rows)
}

func scanVehicles(rows *sql.Rows) ([]*model.Vehicle, error) {
	var vehicles []*model.Vehicle
	for rows.Next() {
		v := &model.Vehicle{}
		if err := rows.Scan(
			&v.ID, &v.DriverID, &v.Make, &v.Model,
			&v.Year, &v.Color, &v.PlateNumber, &v.VehicleType,
			&v.InsuranceExpiry, &v.FitnessExpiry, &v.Status,
			&v.CreatedAt, &v.UpdatedAt,
		); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, v)
	}
	return vehicles, rows.Err()
}
