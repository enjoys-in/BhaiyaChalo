package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, topic string, key string, value []byte) error
}

type Producer struct {
	writer KafkaWriter
}

func NewProducer(writer KafkaWriter) ports.EventPublisher {
	return &Producer{writer: writer}
}

type walletEvent struct {
	TransactionID string  `json:"transaction_id"`
	WalletID      string  `json:"wallet_id"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	Reference     string  `json:"reference"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	Timestamp     string  `json:"timestamp"`
}

func (p *Producer) PublishWalletCredited(ctx context.Context, txn *model.WalletTransaction) error {
	return p.publish(ctx, constants.TopicWalletCredited, txn)
}

func (p *Producer) PublishWalletDebited(ctx context.Context, txn *model.WalletTransaction) error {
	return p.publish(ctx, constants.TopicWalletDebited, txn)
}

func (p *Producer) publish(ctx context.Context, topic string, txn *model.WalletTransaction) error {
	evt := walletEvent{
		TransactionID: txn.ID,
		WalletID:      txn.WalletID,
		Type:          string(txn.Type),
		Amount:        txn.Amount,
		Reference:     txn.Reference,
		BalanceBefore: txn.BalanceBefore,
		BalanceAfter:  txn.BalanceAfter,
		Timestamp:     time.Now().UTC().Format(time.RFC3339),
	}
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	return p.writer.WriteMessage(ctx, topic, txn.WalletID, data)
}
