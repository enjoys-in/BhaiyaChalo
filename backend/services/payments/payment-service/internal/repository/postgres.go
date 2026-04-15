package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/payment-service/internal/ports"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.PaymentRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(ctx context.Context, payment *model.Payment) error {
	query := `INSERT INTO payments (id, booking_id, user_id, amount, currency, method, gateway_id, gateway_status, status, failure_reason, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := r.db.ExecContext(ctx, query,
		payment.ID, payment.BookingID, payment.UserID, payment.Amount, payment.Currency,
		payment.Method, payment.GatewayID, payment.GatewayStatus, payment.Status,
		payment.FailureReason, payment.CreatedAt, payment.UpdatedAt,
	)
	return err
}

func (r *postgresRepo) FindByID(ctx context.Context, id string) (*model.Payment, error) {
	query := `SELECT id, booking_id, user_id, amount, currency, method, gateway_id, gateway_status, status, failure_reason, created_at, updated_at
		FROM payments WHERE id = $1`
	var p model.Payment
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.BookingID, &p.UserID, &p.Amount, &p.Currency, &p.Method,
		&p.GatewayID, &p.GatewayStatus, &p.Status, &p.FailureReason,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("payment not found")
	}
	return &p, err
}

func (r *postgresRepo) FindByBookingID(ctx context.Context, bookingID string) (*model.Payment, error) {
	query := `SELECT id, booking_id, user_id, amount, currency, method, gateway_id, gateway_status, status, failure_reason, created_at, updated_at
		FROM payments WHERE booking_id = $1`
	var p model.Payment
	err := r.db.QueryRowContext(ctx, query, bookingID).Scan(
		&p.ID, &p.BookingID, &p.UserID, &p.Amount, &p.Currency, &p.Method,
		&p.GatewayID, &p.GatewayStatus, &p.Status, &p.FailureReason,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("payment not found")
	}
	return &p, err
}

func (r *postgresRepo) UpdateStatus(ctx context.Context, id string, status model.PaymentStatus, gatewayID, gatewayStatus, failureReason string) error {
	query := `UPDATE payments SET status = $2, gateway_id = $3, gateway_status = $4, failure_reason = $5, updated_at = $6 WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id, status, gatewayID, gatewayStatus, failureReason, time.Now().UTC())
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("payment not found")
	}
	return nil
}

func (r *postgresRepo) CreateRefund(ctx context.Context, refund *model.Refund) error {
	query := `INSERT INTO refunds (id, payment_id, amount, reason, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query,
		refund.ID, refund.PaymentID, refund.Amount, refund.Reason, refund.Status, refund.CreatedAt,
	)
	return err
}

func (r *postgresRepo) FindRefundByPaymentID(ctx context.Context, paymentID string) (*model.Refund, error) {
	query := `SELECT id, payment_id, amount, reason, status, created_at FROM refunds WHERE payment_id = $1`
	var ref model.Refund
	err := r.db.QueryRowContext(ctx, query, paymentID).Scan(
		&ref.ID, &ref.PaymentID, &ref.Amount, &ref.Reason, &ref.Status, &ref.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("refund not found")
	}
	return &ref, err
}
