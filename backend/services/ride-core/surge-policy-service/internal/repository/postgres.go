package repository

import (
	"context"
	"database/sql"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/surge-policy-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.SurgeRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetPolicy(ctx context.Context, cityID string) (*model.SurgePolicy, error) {
	query := `
		SELECT id, city_id, min_demand_supply_ratio, max_multiplier,
			step_size, cooldown_minutes, active, created_at
		FROM surge_policies WHERE city_id = $1 AND active = true`

	p := &model.SurgePolicy{}
	err := r.db.QueryRowContext(ctx, query, cityID).Scan(
		&p.ID, &p.CityID, &p.MinDemandSupplyRatio,
		&p.MaxMultiplier, &p.StepSize, &p.CooldownMinutes,
		&p.Active, &p.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *postgresRepository) UpdatePolicy(ctx context.Context, policy *model.SurgePolicy) error {
	query := `
		UPDATE surge_policies SET min_demand_supply_ratio = $1, max_multiplier = $2,
			step_size = $3, cooldown_minutes = $4, active = $5
		WHERE id = $6`

	_, err := r.db.ExecContext(ctx, query,
		policy.MinDemandSupplyRatio, policy.MaxMultiplier,
		policy.StepSize, policy.CooldownMinutes, policy.Active, policy.ID,
	)
	return err
}

func (r *postgresRepository) SaveZone(ctx context.Context, zone *model.SurgeZone) error {
	query := `
		INSERT INTO surge_zones (id, city_id, geofence_id, current_multiplier, demand_count, supply_count, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
			current_multiplier = EXCLUDED.current_multiplier,
			demand_count = EXCLUDED.demand_count,
			supply_count = EXCLUDED.supply_count,
			updated_at = EXCLUDED.updated_at`

	_, err := r.db.ExecContext(ctx, query,
		zone.ID, zone.CityID, zone.GeofenceID,
		zone.CurrentMultiplier, zone.DemandCount, zone.SupplyCount,
		zone.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) GetZone(ctx context.Context, zoneID string) (*model.SurgeZone, error) {
	query := `
		SELECT id, city_id, geofence_id, current_multiplier, demand_count, supply_count, updated_at
		FROM surge_zones WHERE id = $1`

	z := &model.SurgeZone{}
	err := r.db.QueryRowContext(ctx, query, zoneID).Scan(
		&z.ID, &z.CityID, &z.GeofenceID,
		&z.CurrentMultiplier, &z.DemandCount, &z.SupplyCount,
		&z.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return z, nil
}

func (r *postgresRepository) SaveHistory(ctx context.Context, history *model.SurgeHistory) error {
	query := `
		INSERT INTO surge_history (id, zone_id, multiplier, demand_count, supply_count, calculated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query,
		history.ID, history.ZoneID, history.Multiplier,
		history.DemandCount, history.SupplyCount, history.CalculatedAt,
	)
	return err
}

func (r *postgresRepository) GetHistory(ctx context.Context, zoneID string, limit int) ([]*model.SurgeHistory, error) {
	query := `
		SELECT id, zone_id, multiplier, demand_count, supply_count, calculated_at
		FROM surge_history WHERE zone_id = $1 ORDER BY calculated_at DESC LIMIT $2`

	rows, err := r.db.QueryContext(ctx, query, zoneID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []*model.SurgeHistory
	for rows.Next() {
		h := &model.SurgeHistory{}
		if err := rows.Scan(
			&h.ID, &h.ZoneID, &h.Multiplier,
			&h.DemandCount, &h.SupplyCount, &h.CalculatedAt,
		); err != nil {
			return nil, err
		}
		history = append(history, h)
	}
	return history, rows.Err()
}
