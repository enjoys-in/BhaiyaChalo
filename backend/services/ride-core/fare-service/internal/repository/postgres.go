package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/ports"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.FareRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) SaveCalculation(ctx context.Context, calc *model.FareCalculation) error {
	query := `INSERT INTO fare_calculations
		(id, booking_id, base_price, distance_charge, time_charge, surge_multiplier, surge_amount,
		 toll_charges, tax_amount, promo_discount, total_fare, currency, city_id, vehicle_type, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`
	_, err := r.db.ExecContext(ctx, query,
		calc.ID, calc.BookingID, calc.BasePrice, calc.DistanceCharge, calc.TimeCharge,
		calc.SurgeMultiplier, calc.SurgeAmount, calc.TollCharges, calc.TaxAmount,
		calc.PromoDiscount, calc.TotalFare, calc.Currency, calc.CityID, calc.VehicleType,
		calc.CreatedAt,
	)
	return err
}

func (r *postgresRepo) FindByBookingID(ctx context.Context, bookingID string) (*model.FareCalculation, error) {
	query := `SELECT id, booking_id, base_price, distance_charge, time_charge, surge_multiplier, surge_amount,
		toll_charges, tax_amount, promo_discount, total_fare, currency, city_id, vehicle_type, created_at
		FROM fare_calculations WHERE booking_id = $1 ORDER BY created_at DESC LIMIT 1`

	var calc model.FareCalculation
	err := r.db.QueryRowContext(ctx, query, bookingID).Scan(
		&calc.ID, &calc.BookingID, &calc.BasePrice, &calc.DistanceCharge, &calc.TimeCharge,
		&calc.SurgeMultiplier, &calc.SurgeAmount, &calc.TollCharges, &calc.TaxAmount,
		&calc.PromoDiscount, &calc.TotalFare, &calc.Currency, &calc.CityID, &calc.VehicleType,
		&calc.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("fare calculation not found")
		}
		return nil, err
	}
	return &calc, nil
}

func (r *postgresRepo) GetConfig(ctx context.Context, cityID, vehicleType string) (*model.FareConfig, error) {
	query := `SELECT city_id, vehicle_type, base_price_per_km, base_price_per_min, min_fare, cancellation_fee
		FROM fare_configs WHERE city_id = $1 AND vehicle_type = $2`

	var cfg model.FareConfig
	err := r.db.QueryRowContext(ctx, query, cityID, vehicleType).Scan(
		&cfg.CityID, &cfg.VehicleType, &cfg.BasePricePerKM, &cfg.BasePricePerMin,
		&cfg.MinFare, &cfg.CancellationFee,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("fare config not found for city=%s vehicle=%s", cityID, vehicleType)
		}
		return nil, err
	}
	return &cfg, nil
}
