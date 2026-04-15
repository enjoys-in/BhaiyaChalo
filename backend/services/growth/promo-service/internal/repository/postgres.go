package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/promo-service/internal/ports"
)

type postgresPromoRepo struct {
	db *sql.DB
}

func NewPostgresPromoRepository(db *sql.DB) ports.PromoRepository {
	return &postgresPromoRepo{db: db}
}

func (r *postgresPromoRepo) Create(ctx context.Context, promo *model.PromoCode) error {
	query := `INSERT INTO promo_codes (id, code, city_id, type, discount_value, max_discount, min_order_value, usage_limit, used_count, valid_from, valid_until, active, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	_, err := r.db.ExecContext(ctx, query,
		promo.ID, promo.Code, promo.CityID, promo.Type,
		promo.DiscountValue, promo.MaxDiscount, promo.MinOrderValue,
		promo.UsageLimit, promo.UsedCount, promo.ValidFrom, promo.ValidUntil,
		promo.Active, promo.CreatedAt,
	)
	return err
}

func (r *postgresPromoRepo) FindByCode(ctx context.Context, code string) (*model.PromoCode, error) {
	query := `SELECT id, code, city_id, type, discount_value, max_discount, min_order_value, usage_limit, used_count, valid_from, valid_until, active, created_at
		FROM promo_codes WHERE code = $1`
	p := &model.PromoCode{}
	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&p.ID, &p.Code, &p.CityID, &p.Type,
		&p.DiscountValue, &p.MaxDiscount, &p.MinOrderValue,
		&p.UsageLimit, &p.UsedCount, &p.ValidFrom, &p.ValidUntil,
		&p.Active, &p.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return p, err
}

func (r *postgresPromoRepo) FindByID(ctx context.Context, id string) (*model.PromoCode, error) {
	query := `SELECT id, code, city_id, type, discount_value, max_discount, min_order_value, usage_limit, used_count, valid_from, valid_until, active, created_at
		FROM promo_codes WHERE id = $1`
	p := &model.PromoCode{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.Code, &p.CityID, &p.Type,
		&p.DiscountValue, &p.MaxDiscount, &p.MinOrderValue,
		&p.UsageLimit, &p.UsedCount, &p.ValidFrom, &p.ValidUntil,
		&p.Active, &p.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return p, err
}

func (r *postgresPromoRepo) ListActive(ctx context.Context) ([]model.PromoCode, error) {
	query := `SELECT id, code, city_id, type, discount_value, max_discount, min_order_value, usage_limit, used_count, valid_from, valid_until, active, created_at
		FROM promo_codes WHERE active = true AND valid_until > $1 ORDER BY created_at DESC`
	return r.scanPromos(ctx, query, time.Now())
}

func (r *postgresPromoRepo) ListByCityID(ctx context.Context, cityID string) ([]model.PromoCode, error) {
	query := `SELECT id, code, city_id, type, discount_value, max_discount, min_order_value, usage_limit, used_count, valid_from, valid_until, active, created_at
		FROM promo_codes WHERE city_id = $1 AND active = true AND valid_until > $2 ORDER BY created_at DESC`
	return r.scanPromos(ctx, query, cityID, time.Now())
}

func (r *postgresPromoRepo) IncrementUsage(ctx context.Context, promoID string) error {
	query := `UPDATE promo_codes SET used_count = used_count + 1 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, promoID)
	return err
}

func (r *postgresPromoRepo) RecordUsage(ctx context.Context, usage *model.PromoUsage) error {
	query := `INSERT INTO promo_usages (id, promo_id, user_id, booking_id, discount_applied, used_at)
		VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := r.db.ExecContext(ctx, query,
		usage.ID, usage.PromoID, usage.UserID, usage.BookingID,
		usage.DiscountApplied, usage.UsedAt,
	)
	return err
}

func (r *postgresPromoRepo) HasUserUsed(ctx context.Context, promoID, userID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM promo_usages WHERE promo_id = $1 AND user_id = $2)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, promoID, userID).Scan(&exists)
	return exists, err
}

func (r *postgresPromoRepo) scanPromos(ctx context.Context, query string, args ...interface{}) ([]model.PromoCode, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var promos []model.PromoCode
	for rows.Next() {
		var p model.PromoCode
		if err := rows.Scan(
			&p.ID, &p.Code, &p.CityID, &p.Type,
			&p.DiscountValue, &p.MaxDiscount, &p.MinOrderValue,
			&p.UsageLimit, &p.UsedCount, &p.ValidFrom, &p.ValidUntil,
			&p.Active, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		promos = append(promos, p)
	}
	return promos, rows.Err()
}
