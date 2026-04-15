package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/ports"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresOTPRepository(db *sql.DB) ports.OTPRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(ctx context.Context, otp *model.OTP) error {
	query := `
		INSERT INTO otps (id, phone, code, purpose, verified, attempts, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(ctx, query,
		otp.ID, otp.Phone, otp.Code, otp.Purpose,
		otp.Verified, otp.Attempts, otp.ExpiresAt, otp.CreatedAt,
	)
	return err
}

func (r *postgresRepo) FindLatest(ctx context.Context, phone string, purpose model.Purpose) (*model.OTP, error) {
	query := `
		SELECT id, phone, code, purpose, verified, attempts, expires_at, created_at
		FROM otps
		WHERE phone = $1 AND purpose = $2
		ORDER BY created_at DESC
		LIMIT 1`

	var otp model.OTP
	err := r.db.QueryRowContext(ctx, query, phone, purpose).Scan(
		&otp.ID, &otp.Phone, &otp.Code, &otp.Purpose,
		&otp.Verified, &otp.Attempts, &otp.ExpiresAt, &otp.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *postgresRepo) MarkVerified(ctx context.Context, id string) error {
	query := `UPDATE otps SET verified = true WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *postgresRepo) IncrementAttempts(ctx context.Context, id string) error {
	query := `UPDATE otps SET attempts = attempts + 1 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
