package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/ports"
)

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ports.WalletRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(ctx context.Context, wallet *model.Wallet) error {
	query := `INSERT INTO wallets (id, user_id, balance, currency, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query,
		wallet.ID, wallet.UserID, wallet.Balance, wallet.Currency,
		wallet.Status, wallet.CreatedAt, wallet.UpdatedAt,
	)
	return err
}

func (r *postgresRepo) FindByUserID(ctx context.Context, userID string) (*model.Wallet, error) {
	query := `SELECT id, user_id, balance, currency, status, created_at, updated_at FROM wallets WHERE user_id = $1`
	var w model.Wallet
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&w.ID, &w.UserID, &w.Balance, &w.Currency, &w.Status, &w.CreatedAt, &w.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("wallet not found")
	}
	return &w, err
}

func (r *postgresRepo) Credit(ctx context.Context, walletID string, amount float64, txn *model.WalletTransaction) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var balance float64
	err = tx.QueryRowContext(ctx, `SELECT balance FROM wallets WHERE id = $1 FOR UPDATE`, walletID).Scan(&balance)
	if err != nil {
		return fmt.Errorf("wallet not found")
	}

	newBalance := balance + amount
	txn.BalanceBefore = balance
	txn.BalanceAfter = newBalance

	_, err = tx.ExecContext(ctx,
		`UPDATE wallets SET balance = $2, updated_at = $3 WHERE id = $1`,
		walletID, newBalance, time.Now().UTC(),
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO wallet_transactions (id, wallet_id, type, amount, reference, description, balance_before, balance_after, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		txn.ID, txn.WalletID, txn.Type, txn.Amount, txn.Reference,
		txn.Description, txn.BalanceBefore, txn.BalanceAfter, txn.CreatedAt,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *postgresRepo) Debit(ctx context.Context, walletID string, amount float64, txn *model.WalletTransaction) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var balance float64
	err = tx.QueryRowContext(ctx, `SELECT balance FROM wallets WHERE id = $1 FOR UPDATE`, walletID).Scan(&balance)
	if err != nil {
		return fmt.Errorf("wallet not found")
	}

	if balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	newBalance := balance - amount
	txn.BalanceBefore = balance
	txn.BalanceAfter = newBalance

	_, err = tx.ExecContext(ctx,
		`UPDATE wallets SET balance = $2, updated_at = $3 WHERE id = $1`,
		walletID, newBalance, time.Now().UTC(),
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO wallet_transactions (id, wallet_id, type, amount, reference, description, balance_before, balance_after, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		txn.ID, txn.WalletID, txn.Type, txn.Amount, txn.Reference,
		txn.Description, txn.BalanceBefore, txn.BalanceAfter, txn.CreatedAt,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *postgresRepo) GetTransactions(ctx context.Context, walletID string, limit, offset int) ([]model.WalletTransaction, error) {
	query := `SELECT id, wallet_id, type, amount, reference, description, balance_before, balance_after, created_at
		FROM wallet_transactions WHERE wallet_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, walletID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txns []model.WalletTransaction
	for rows.Next() {
		var t model.WalletTransaction
		if err := rows.Scan(&t.ID, &t.WalletID, &t.Type, &t.Amount, &t.Reference,
			&t.Description, &t.BalanceBefore, &t.BalanceAfter, &t.CreatedAt); err != nil {
			return nil, err
		}
		txns = append(txns, t)
	}
	return txns, rows.Err()
}

func (r *postgresRepo) GetBalance(ctx context.Context, userID string) (*model.Wallet, error) {
	return r.FindByUserID(ctx, userID)
}
