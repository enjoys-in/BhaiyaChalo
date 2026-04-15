package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/growth/referral-service/internal/ports"
)

type postgresReferralRepo struct {
	db *sql.DB
}

func NewPostgresReferralRepository(db *sql.DB) ports.ReferralRepository {
	return &postgresReferralRepo{db: db}
}

func (r *postgresReferralRepo) CreateCode(ctx context.Context, code *model.ReferralCode) error {
	query := `INSERT INTO referral_codes (id, user_id, code, active, created_at)
		VALUES ($1,$2,$3,$4,$5)`
	_, err := r.db.ExecContext(ctx, query, code.ID, code.UserID, code.Code, code.Active, code.CreatedAt)
	return err
}

func (r *postgresReferralRepo) FindByCode(ctx context.Context, code string) (*model.ReferralCode, error) {
	query := `SELECT id, user_id, code, active, created_at FROM referral_codes WHERE code = $1`
	rc := &model.ReferralCode{}
	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&rc.ID, &rc.UserID, &rc.Code, &rc.Active, &rc.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return rc, err
}

func (r *postgresReferralRepo) FindByUserID(ctx context.Context, userID string) (*model.ReferralCode, error) {
	query := `SELECT id, user_id, code, active, created_at FROM referral_codes WHERE user_id = $1`
	rc := &model.ReferralCode{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&rc.ID, &rc.UserID, &rc.Code, &rc.Active, &rc.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return rc, err
}

func (r *postgresReferralRepo) CreateReferral(ctx context.Context, referral *model.Referral) error {
	query := `INSERT INTO referrals (id, referrer_id, referee_id, referral_code_id, status, completed_at, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := r.db.ExecContext(ctx, query,
		referral.ID, referral.ReferrerID, referral.RefereeID,
		referral.ReferralCodeID, referral.Status, referral.CompletedAt, referral.CreatedAt,
	)
	return err
}

func (r *postgresReferralRepo) FindReferralByID(ctx context.Context, id string) (*model.Referral, error) {
	query := `SELECT id, referrer_id, referee_id, referral_code_id, status, completed_at, created_at
		FROM referrals WHERE id = $1`
	ref := &model.Referral{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&ref.ID, &ref.ReferrerID, &ref.RefereeID,
		&ref.ReferralCodeID, &ref.Status, &ref.CompletedAt, &ref.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return ref, err
}

func (r *postgresReferralRepo) UpdateReferralStatus(ctx context.Context, id string, status model.ReferralStatus) error {
	query := `UPDATE referrals SET status = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *postgresReferralRepo) CreateReward(ctx context.Context, reward *model.ReferralReward) error {
	query := `INSERT INTO referral_rewards (id, referral_id, user_id, type, amount, currency, credited_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := r.db.ExecContext(ctx, query,
		reward.ID, reward.ReferralID, reward.UserID, reward.Type,
		reward.Amount, reward.Currency, reward.CreditedAt,
	)
	return err
}

func (r *postgresReferralRepo) GetStats(ctx context.Context, userID string) (total, completed, pending int, earnings float64, err error) {
	query := `SELECT 
		COUNT(*) as total,
		COUNT(*) FILTER (WHERE status = 'completed' OR status = 'rewarded') as completed,
		COUNT(*) FILTER (WHERE status = 'pending') as pending
		FROM referrals WHERE referrer_id = $1`
	err = r.db.QueryRowContext(ctx, query, userID).Scan(&total, &completed, &pending)
	if err != nil {
		return
	}

	earningsQuery := `SELECT COALESCE(SUM(amount), 0) FROM referral_rewards WHERE user_id = $1`
	err = r.db.QueryRowContext(ctx, earningsQuery, userID).Scan(&earnings)
	return
}
