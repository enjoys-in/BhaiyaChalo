package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/model"
)

type EventPublisher interface {
	PublishWalletCredited(ctx context.Context, txn *model.WalletTransaction) error
	PublishWalletDebited(ctx context.Context, txn *model.WalletTransaction) error
}
