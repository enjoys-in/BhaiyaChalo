package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.EarningsRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Record(ctx context.Context, earning *model.Earning) error {
	query := `
		INSERT INTO earnings (id, driver_id, trip_id, fare_amount, commission,
			incentive_bonus, tip_amount, net_earning, currency, earned_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.ExecContext(ctx, query,
		earning.ID, earning.DriverID, earning.TripID,
		earning.FareAmount, earning.Commission, earning.IncentiveBonus,
		earning.TipAmount, earning.NetEarning, earning.Currency,
		earning.EarnedAt, earning.CreatedAt,
	)
	return err
}

func (r *postgresRepository) FindByDriverID(ctx context.Context, driverID string, from, to time.Time) ([]*model.Earning, error) {
	query := `
		SELECT id, driver_id, trip_id, fare_amount, commission,
			incentive_bonus, tip_amount, net_earning, currency, earned_at, created_at
		FROM earnings
		WHERE driver_id = $1 AND earned_at >= $2 AND earned_at <= $3
		ORDER BY earned_at DESC`

	rows, err := r.db.QueryContext(ctx, query, driverID, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var earnings []*model.Earning
	for rows.Next() {
		e := &model.Earning{}
		if err := rows.Scan(
			&e.ID, &e.DriverID, &e.TripID, &e.FareAmount, &e.Commission,
			&e.IncentiveBonus, &e.TipAmount, &e.NetEarning, &e.Currency,
			&e.EarnedAt, &e.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan earning: %w", err)
		}
		earnings = append(earnings, e)
	}
	return earnings, rows.Err()
}

func (r *postgresRepository) GetDailySummary(ctx context.Context, driverID string, date time.Time) (*model.DailySummary, error) {
	query := `
		SELECT driver_id,
			$2::date AS date,
			COUNT(*) AS total_trips,
			COALESCE(SUM(fare_amount), 0) AS total_fare,
			COALESCE(SUM(commission), 0) AS total_commission,
			COALESCE(SUM(incentive_bonus), 0) AS total_incentive,
			COALESCE(SUM(tip_amount), 0) AS total_tips,
			COALESCE(SUM(net_earning), 0) AS net_earning
		FROM earnings
		WHERE driver_id = $1 AND earned_at::date = $2::date
		GROUP BY driver_id`

	s := &model.DailySummary{}
	err := r.db.QueryRowContext(ctx, query, driverID, date).Scan(
		&s.DriverID, &s.Date, &s.TotalTrips, &s.TotalFare,
		&s.TotalCommission, &s.TotalIncentive, &s.TotalTips, &s.NetEarning,
	)
	if err == sql.ErrNoRows {
		return &model.DailySummary{DriverID: driverID, Date: date}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("scan daily summary: %w", err)
	}
	return s, nil
}

func (r *postgresRepository) GetWeeklySummary(ctx context.Context, driverID string, weekStart time.Time) (*model.WeeklySummary, error) {
	weekEnd := weekStart.AddDate(0, 0, 7)
	query := `
		SELECT driver_id,
			$2::date AS week_start,
			$3::date AS week_end,
			COUNT(*) AS total_trips,
			COALESCE(SUM(fare_amount), 0) AS total_fare,
			COALESCE(SUM(commission), 0) AS total_commission,
			COALESCE(SUM(incentive_bonus), 0) AS total_incentive,
			COALESCE(SUM(tip_amount), 0) AS total_tips,
			COALESCE(SUM(net_earning), 0) AS net_earning
		FROM earnings
		WHERE driver_id = $1 AND earned_at >= $2 AND earned_at < $3
		GROUP BY driver_id`

	s := &model.WeeklySummary{}
	err := r.db.QueryRowContext(ctx, query, driverID, weekStart, weekEnd).Scan(
		&s.DriverID, &s.WeekStart, &s.WeekEnd, &s.TotalTrips, &s.TotalFare,
		&s.TotalCommission, &s.TotalIncentive, &s.TotalTips, &s.NetEarning,
	)
	if err == sql.ErrNoRows {
		return &model.WeeklySummary{DriverID: driverID, WeekStart: weekStart, WeekEnd: weekEnd}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("scan weekly summary: %w", err)
	}
	return s, nil
}

func (r *postgresRepository) GetTotalByDateRange(ctx context.Context, driverID string, from, to time.Time) (float64, error) {
	query := `SELECT COALESCE(SUM(net_earning), 0) FROM earnings WHERE driver_id = $1 AND earned_at >= $2 AND earned_at <= $3`
	var total float64
	err := r.db.QueryRowContext(ctx, query, driverID, from, to).Scan(&total)
	return total, err
}
