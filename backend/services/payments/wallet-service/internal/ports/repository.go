package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/model"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *model.Wallet) error
	FindByUserID(ctx context.Context, userID string) (*model.Wallet, error)
	Credit(ctx context.Context, walletID string, amount float64, txn *model.WalletTransaction) error
	Debit(ctx context.Context, walletID string, amount float64, txn *model.WalletTransaction) error
	GetTransactions(ctx context.Context, walletID string, limit, offset int) ([]model.WalletTransaction, error)
	GetBalance(ctx context.Context, userID string) (*model.Wallet, error)
}
