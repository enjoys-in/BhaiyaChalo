package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/geofence-service/internal/ports"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.GeofenceRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, fence *model.Geofence) error {
	polygonJSON, err := json.Marshal(fence.Polygon)
	if err != nil {
		return fmt.Errorf("marshal polygon: %w", err)
	}
	metadataJSON, err := json.Marshal(fence.Metadata)
	if err != nil {
		return fmt.Errorf("marshal metadata: %w", err)
	}

	query := `
		INSERT INTO geofences (id, city_id, name, type, polygon, center_lat, center_lng,
			radius_km, active, metadata, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err = r.db.ExecContext(ctx, query,
		fence.ID, fence.CityID, fence.Name, fence.Type,
		polygonJSON, fence.CenterLat, fence.CenterLng,
		fence.RadiusKM, fence.Active, metadataJSON,
		fence.CreatedAt, fence.UpdatedAt,
	)
	return err
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (*model.Geofence, error) {
	query := `
		SELECT id, city_id, name, type, polygon, center_lat, center_lng,
			radius_km, active, metadata, created_at, updated_at
		FROM geofences WHERE id = $1`

	return r.scanGeofence(r.db.QueryRowContext(ctx, query, id))
}

func (r *postgresRepository) FindByCityID(ctx context.Context, cityID string) ([]*model.Geofence, error) {
	query := `
		SELECT id, city_id, name, type, polygon, center_lat, center_lng,
			radius_km, active, metadata, created_at, updated_at
		FROM geofences WHERE city_id = $1 ORDER BY name`

	rows, err := r.db.QueryContext(ctx, query, cityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGeofences(rows)
}

func (r *postgresRepository) Update(ctx context.Context, fence *model.Geofence) error {
	polygonJSON, err := json.Marshal(fence.Polygon)
	if err != nil {
		return fmt.Errorf("marshal polygon: %w", err)
	}
	metadataJSON, err := json.Marshal(fence.Metadata)
	if err != nil {
		return fmt.Errorf("marshal metadata: %w", err)
	}

	query := `
		UPDATE geofences SET name = $1, type = $2, polygon = $3, center_lat = $4,
			center_lng = $5, radius_km = $6, active = $7, metadata = $8, updated_at = $9
		WHERE id = $10`

	_, err = r.db.ExecContext(ctx, query,
		fence.Name, fence.Type, polygonJSON,
		fence.CenterLat, fence.CenterLng, fence.RadiusKM,
		fence.Active, metadataJSON, fence.UpdatedAt, fence.ID,
	)
	return err
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM geofences WHERE id = $1`, id)
	return err
}

// FindContaining uses PostGIS ST_Contains to find geofences containing the given point.
func (r *postgresRepository) FindContaining(ctx context.Context, lat, lng float64) ([]*model.Geofence, error) {
	query := `
		SELECT id, city_id, name, type, polygon, center_lat, center_lng,
			radius_km, active, metadata, created_at, updated_at
		FROM geofences
		WHERE active = true
			AND ST_Contains(
				ST_SetSRID(ST_GeomFromGeoJSON(polygon::text), 4326),
				ST_SetSRID(ST_MakePoint($1, $2), 4326)
			)
		ORDER BY name`

	rows, err := r.db.QueryContext(ctx, query, lng, lat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanGeofences(rows)
}

func (r *postgresRepository) scanGeofence(row *sql.Row) (*model.Geofence, error) {
	f := &model.Geofence{}
	var polygonJSON, metadataJSON []byte

	err := row.Scan(
		&f.ID, &f.CityID, &f.Name, &f.Type,
		&polygonJSON, &f.CenterLat, &f.CenterLng,
		&f.RadiusKM, &f.Active, &metadataJSON,
		&f.CreatedAt, &f.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(polygonJSON, &f.Polygon); err != nil {
		return nil, fmt.Errorf("unmarshal polygon: %w", err)
	}
	if len(metadataJSON) > 0 {
		if err := json.Unmarshal(metadataJSON, &f.Metadata); err != nil {
			return nil, fmt.Errorf("unmarshal metadata: %w", err)
		}
	}

	return f, nil
}

func (r *postgresRepository) scanGeofences(rows *sql.Rows) ([]*model.Geofence, error) {
	var fences []*model.Geofence
	for rows.Next() {
		f := &model.Geofence{}
		var polygonJSON, metadataJSON []byte

		if err := rows.Scan(
			&f.ID, &f.CityID, &f.Name, &f.Type,
			&polygonJSON, &f.CenterLat, &f.CenterLng,
			&f.RadiusKM, &f.Active, &metadataJSON,
			&f.CreatedAt, &f.UpdatedAt,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(polygonJSON, &f.Polygon); err != nil {
			return nil, fmt.Errorf("unmarshal polygon: %w", err)
		}
		if len(metadataJSON) > 0 {
			if err := json.Unmarshal(metadataJSON, &f.Metadata); err != nil {
				return nil, fmt.Errorf("unmarshal metadata: %w", err)
			}
		}

		fences = append(fences, f)
	}
	return fences, rows.Err()
}
