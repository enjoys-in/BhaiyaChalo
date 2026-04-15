package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/pricing-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.PricingRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) SaveEstimate(ctx context.Context, estimate *model.PriceEstimate) error {
	query := `
		INSERT INTO price_estimates (id, city_id, vehicle_type, distance_km, duration_min,
			base_fare, surge_multiplier, estimated_fare, currency, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := r.db.ExecContext(ctx, query,
		estimate.ID, estimate.CityID, estimate.VehicleType,
		estimate.DistanceKM, estimate.DurationMin, estimate.BaseFare,
		estimate.SurgeMultiplier, estimate.EstimatedFare, estimate.Currency,
		estimate.CreatedAt,
	)
	return err
}

func (r *postgresRepository) GetRule(ctx context.Context, cityID, vehicleType string) (*model.PricingRule, error) {
	query := `
		SELECT id, city_id, vehicle_type, base_fare_per_km, base_fare_per_min,
			min_fare, max_fare, booking_fee, active, created_at, updated_at
		FROM pricing_rules WHERE city_id = $1 AND vehicle_type = $2 AND active = true`

	rule := &model.PricingRule{}
	err := r.db.QueryRowContext(ctx, query, cityID, vehicleType).Scan(
		&rule.ID, &rule.CityID, &rule.VehicleType,
		&rule.BaseFarePerKM, &rule.BaseFarePerMin,
		&rule.MinFare, &rule.MaxFare, &rule.BookingFee,
		&rule.Active, &rule.CreatedAt, &rule.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return rule, nil
}

func (r *postgresRepository) ListRules(ctx context.Context, cityID string) ([]*model.PricingRule, error) {
	query := `
		SELECT id, city_id, vehicle_type, base_fare_per_km, base_fare_per_min,
			min_fare, max_fare, booking_fee, active, created_at, updated_at
		FROM pricing_rules WHERE city_id = $1 ORDER BY vehicle_type`

	rows, err := r.db.QueryContext(ctx, query, cityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanRules(rows)
}

func (r *postgresRepository) CreateRule(ctx context.Context, rule *model.PricingRule) error {
	query := `
		INSERT INTO pricing_rules (id, city_id, vehicle_type, base_fare_per_km, base_fare_per_min,
			min_fare, max_fare, booking_fee, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.ExecContext(ctx, query,
		rule.ID, rule.CityID, rule.VehicleType,
		rule.BaseFarePerKM, rule.BaseFarePerMin,
		rule.MinFare, rule.MaxFare, rule.BookingFee,
		rule.Active, rule.CreatedAt, rule.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) UpdateRule(ctx context.Context, rule *model.PricingRule) error {
	query := `
		UPDATE pricing_rules SET base_fare_per_km = $1, base_fare_per_min = $2,
			min_fare = $3, max_fare = $4, booking_fee = $5, active = $6, updated_at = $7
		WHERE id = $8`

	_, err := r.db.ExecContext(ctx, query,
		rule.BaseFarePerKM, rule.BaseFarePerMin,
		rule.MinFare, rule.MaxFare, rule.BookingFee,
		rule.Active, rule.UpdatedAt, rule.ID,
	)
	return err
}

func scanRules(rows *sql.Rows) ([]*model.PricingRule, error) {
	var rules []*model.PricingRule
	for rows.Next() {
		rule := &model.PricingRule{}
		if err := rows.Scan(
			&rule.ID, &rule.CityID, &rule.VehicleType,
			&rule.BaseFarePerKM, &rule.BaseFarePerMin,
			&rule.MinFare, &rule.MaxFare, &rule.BookingFee,
			&rule.Active, &rule.CreatedAt, &rule.UpdatedAt,
		); err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, rows.Err()
}
