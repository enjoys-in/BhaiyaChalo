package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/model"
)

type WalletService interface {
	CreateWallet(ctx context.Context, req dto.CreateWalletRequest) (*model.Wallet, error)
	Credit(ctx context.Context, req dto.CreditRequest) (*model.WalletTransaction, error)
	Debit(ctx context.Context, req dto.DebitRequest) (*model.WalletTransaction, error)
	GetBalance(ctx context.Context, userID string) (*model.Wallet, error)
	GetTransactions(ctx context.Context, walletID string, limit, offset int) ([]model.WalletTransaction, error)
}
