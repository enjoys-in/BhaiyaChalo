package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/campaign-service/internal/ports"
)

type postgresCampaignRepo struct {
	db *sql.DB
}

func NewPostgresCampaignRepository(db *sql.DB) ports.CampaignRepository {
	return &postgresCampaignRepo{db: db}
}

func (r *postgresCampaignRepo) Create(ctx context.Context, c *model.Campaign) error {
	query := `INSERT INTO campaigns (id, name, description, type, target_audience, city_id, promo_code_id, status, scheduled_at, started_at, completed_at, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err := r.db.ExecContext(ctx, query,
		c.ID, c.Name, c.Description, c.Type, c.TargetAudience,
		c.CityID, c.PromoCodeID, c.Status, c.ScheduledAt,
		c.StartedAt, c.CompletedAt, c.CreatedAt,
	)
	return err
}

func (r *postgresCampaignRepo) FindByID(ctx context.Context, id string) (*model.Campaign, error) {
	query := `SELECT id, name, description, type, target_audience, city_id, promo_code_id, status, scheduled_at, started_at, completed_at, created_at
		FROM campaigns WHERE id = $1`
	c := &model.Campaign{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.Name, &c.Description, &c.Type, &c.TargetAudience,
		&c.CityID, &c.PromoCodeID, &c.Status, &c.ScheduledAt,
		&c.StartedAt, &c.CompletedAt, &c.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return c, err
}

func (r *postgresCampaignRepo) List(ctx context.Context, cityID string) ([]model.Campaign, error) {
	var query string
	var args []interface{}

	if cityID != "" {
		query = `SELECT id, name, description, type, target_audience, city_id, promo_code_id, status, scheduled_at, started_at, completed_at, created_at
			FROM campaigns WHERE city_id = $1 ORDER BY created_at DESC`
		args = append(args, cityID)
	} else {
		query = `SELECT id, name, description, type, target_audience, city_id, promo_code_id, status, scheduled_at, started_at, completed_at, created_at
			FROM campaigns ORDER BY created_at DESC`
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []model.Campaign
	for rows.Next() {
		var c model.Campaign
		if err := rows.Scan(
			&c.ID, &c.Name, &c.Description, &c.Type, &c.TargetAudience,
			&c.CityID, &c.PromoCodeID, &c.Status, &c.ScheduledAt,
			&c.StartedAt, &c.CompletedAt, &c.CreatedAt,
		); err != nil {
			return nil, err
		}
		campaigns = append(campaigns, c)
	}
	return campaigns, rows.Err()
}

func (r *postgresCampaignRepo) UpdateStatus(ctx context.Context, id string, status model.CampaignStatus) error {
	query := `UPDATE campaigns SET status = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *postgresCampaignRepo) SaveStats(ctx context.Context, stats *model.CampaignStats) error {
	query := `INSERT INTO campaign_stats (campaign_id, sent, delivered, opened, clicked, converted)
		VALUES ($1,$2,$3,$4,$5,$6)
		ON CONFLICT (campaign_id) DO UPDATE SET sent=$2, delivered=$3, opened=$4, clicked=$5, converted=$6`
	_, err := r.db.ExecContext(ctx, query,
		stats.CampaignID, stats.Sent, stats.Delivered,
		stats.Opened, stats.Clicked, stats.Converted,
	)
	return err
}

func (r *postgresCampaignRepo) GetStats(ctx context.Context, campaignID string) (*model.CampaignStats, error) {
	query := `SELECT campaign_id, sent, delivered, opened, clicked, converted FROM campaign_stats WHERE campaign_id = $1`
	s := &model.CampaignStats{}
	err := r.db.QueryRowContext(ctx, query, campaignID).Scan(
		&s.CampaignID, &s.Sent, &s.Delivered,
		&s.Opened, &s.Clicked, &s.Converted,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &model.CampaignStats{CampaignID: campaignID}, nil
	}
	return s, err
}
