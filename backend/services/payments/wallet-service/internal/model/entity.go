package model

import "time"

type WalletStatus string

const (
	WalletStatusActive WalletStatus = "active"
	WalletStatusFrozen WalletStatus = "frozen"
)

type TransactionType string

const (
	TransactionTypeCredit TransactionType = "credit"
	TransactionTypeDebit  TransactionType = "debit"
)

type Wallet struct {
	ID        string       `json:"id" db:"id"`
	UserID    string       `json:"user_id" db:"user_id"`
	Balance   float64      `json:"balance" db:"balance"`
	Currency  string       `json:"currency" db:"currency"`
	Status    WalletStatus `json:"status" db:"status"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
}

type WalletTransaction struct {
	ID            string          `json:"id" db:"id"`
	WalletID      string          `json:"wallet_id" db:"wallet_id"`
	Type          TransactionType `json:"type" db:"type"`
	Amount        float64         `json:"amount" db:"amount"`
	Reference     string          `json:"reference" db:"reference"`
	Description   string          `json:"description" db:"description"`
	BalanceBefore float64         `json:"balance_before" db:"balance_before"`
	BalanceAfter  float64         `json:"balance_after" db:"balance_after"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
}
