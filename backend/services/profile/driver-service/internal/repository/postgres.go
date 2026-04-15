package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/driver-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.DriverRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, driver *model.Driver) error {
	query := `
		INSERT INTO drivers (id, first_name, last_name, phone, email, avatar_url, license_number,
			city_id, rating, total_trips, status, onboarding_step, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`

	_, err := r.db.ExecContext(ctx, query,
		driver.ID, driver.FirstName, driver.LastName, driver.Phone,
		driver.Email, driver.AvatarURL, driver.LicenseNumber,
		driver.CityID, driver.Rating, driver.TotalTrips,
		driver.Status, driver.OnboardingStep,
		driver.CreatedAt, driver.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*model.Driver, error) {
	query := `
		SELECT id, first_name, last_name, phone, email, avatar_url, license_number,
			city_id, rating, total_trips, status, onboarding_step, created_at, updated_at
		FROM drivers WHERE id = $1`

	d := &model.Driver{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&d.ID, &d.FirstName, &d.LastName, &d.Phone,
		&d.Email, &d.AvatarURL, &d.LicenseNumber,
		&d.CityID, &d.Rating, &d.TotalTrips,
		&d.Status, &d.OnboardingStep,
		&d.CreatedAt, &d.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *postgresRepository) FindByPhone(ctx context.Context, phone string) (*model.Driver, error) {
	query := `
		SELECT id, first_name, last_name, phone, email, avatar_url, license_number,
			city_id, rating, total_trips, status, onboarding_step, created_at, updated_at
		FROM drivers WHERE phone = $1`

	d := &model.Driver{}
	err := r.db.QueryRowContext(ctx, query, phone).Scan(
		&d.ID, &d.FirstName, &d.LastName, &d.Phone,
		&d.Email, &d.AvatarURL, &d.LicenseNumber,
		&d.CityID, &d.Rating, &d.TotalTrips,
		&d.Status, &d.OnboardingStep,
		&d.CreatedAt, &d.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *postgresRepository) Update(ctx context.Context, driver *model.Driver) error {
	query := `
		UPDATE drivers SET first_name=$1, last_name=$2, email=$3, avatar_url=$4,
			license_number=$5, city_id=$6, status=$7, onboarding_step=$8, updated_at=$9
		WHERE id = $10`

	_, err := r.db.ExecContext(ctx, query,
		driver.FirstName, driver.LastName, driver.Email, driver.AvatarURL,
		driver.LicenseNumber, driver.CityID, driver.Status, driver.OnboardingStep,
		driver.UpdatedAt, driver.ID,
	)
	return err
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM drivers WHERE id = $1`, id)
	return err
}

func (r *postgresRepository) UpdatePreference(ctx context.Context, pref *model.DriverPreference) error {
	query := `
		INSERT INTO driver_preferences (driver_id, auto_accept, max_distance, preferred_zones)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (driver_id) DO UPDATE SET
			auto_accept = EXCLUDED.auto_accept,
			max_distance = EXCLUDED.max_distance,
			preferred_zones = EXCLUDED.preferred_zones`

	zones := strings.Join(pref.PreferredZones, ",")
	_, err := r.db.ExecContext(ctx, query, pref.DriverID, pref.AutoAccept, pref.MaxDistance, zones)
	return err
}

func (r *postgresRepository) GetPreference(ctx context.Context, driverID string) (*model.DriverPreference, error) {
	query := `SELECT driver_id, auto_accept, max_distance, preferred_zones FROM driver_preferences WHERE driver_id = $1`

	p := &model.DriverPreference{}
	var zones string
	err := r.db.QueryRowContext(ctx, query, driverID).Scan(&p.DriverID, &p.AutoAccept, &p.MaxDistance, &zones)
	if err != nil {
		return nil, err
	}
	if zones != "" {
		p.PreferredZones = strings.Split(zones, ",")
	}
	return p, nil
}

func (r *postgresRepository) ListByCityID(ctx context.Context, cityID string) ([]*model.Driver, error) {
	query := `
		SELECT id, first_name, last_name, phone, email, avatar_url, license_number,
			city_id, rating, total_trips, status, onboarding_step, created_at, updated_at
		FROM drivers WHERE city_id = $1 ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, cityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []*model.Driver
	for rows.Next() {
		d := &model.Driver{}
		if err := rows.Scan(
			&d.ID, &d.FirstName, &d.LastName, &d.Phone,
			&d.Email, &d.AvatarURL, &d.LicenseNumber,
			&d.CityID, &d.Rating, &d.TotalTrips,
			&d.Status, &d.OnboardingStep,
			&d.CreatedAt, &d.UpdatedAt,
		); err != nil {
			return nil, err
		}
		drivers = append(drivers, d)
	}
	return drivers, rows.Err()
}

func scanDriver(rows *sql.Rows) (*model.Driver, error) {
	d := &model.Driver{}
	err := rows.Scan(
		&d.ID, &d.FirstName, &d.LastName, &d.Phone,
		&d.Email, &d.AvatarURL, &d.LicenseNumber,
		&d.CityID, &d.Rating, &d.TotalTrips,
		&d.Status, &d.OnboardingStep,
		&d.CreatedAt, &d.UpdatedAt,
	)
	return d, err
}
