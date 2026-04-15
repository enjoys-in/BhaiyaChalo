package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/ports"
)

type walletService struct {
	repo      ports.WalletRepository
	publisher ports.EventPublisher
}

func NewWalletService(repo ports.WalletRepository, publisher ports.EventPublisher) ports.WalletService {
	return &walletService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *walletService) CreateWallet(ctx context.Context, req dto.CreateWalletRequest) (*model.Wallet, error) {
	existing, _ := s.repo.FindByUserID(ctx, req.UserID)
	if existing != nil {
		return nil, fmt.Errorf("wallet already exists for user")
	}

	now := time.Now().UTC()
	wallet := &model.Wallet{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		Balance:   0,
		Currency:  req.Currency,
		Status:    model.WalletStatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(ctx, wallet); err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *walletService) Credit(ctx context.Context, req dto.CreditRequest) (*model.WalletTransaction, error) {
	now := time.Now().UTC()
	txn := &model.WalletTransaction{
		ID:          uuid.New().String(),
		WalletID:    req.WalletID,
		Type:        model.TransactionTypeCredit,
		Amount:      req.Amount,
		Reference:   req.Reference,
		Description: req.Description,
		CreatedAt:   now,
	}

	if err := s.repo.Credit(ctx, req.WalletID, req.Amount, txn); err != nil {
		return nil, err
	}

	_ = s.publisher.PublishWalletCredited(ctx, txn)
	return txn, nil
}

func (s *walletService) Debit(ctx context.Context, req dto.DebitRequest) (*model.WalletTransaction, error) {
	now := time.Now().UTC()
	txn := &model.WalletTransaction{
		ID:          uuid.New().String(),
		WalletID:    req.WalletID,
		Type:        model.TransactionTypeDebit,
		Amount:      req.Amount,
		Reference:   req.Reference,
		Description: req.Description,
		CreatedAt:   now,
	}

	if err := s.repo.Debit(ctx, req.WalletID, req.Amount, txn); err != nil {
		return nil, err
	}

	_ = s.publisher.PublishWalletDebited(ctx, txn)
	return txn, nil
}

func (s *walletService) GetBalance(ctx context.Context, userID string) (*model.Wallet, error) {
	return s.repo.GetBalance(ctx, userID)
}

func (s *walletService) GetTransactions(ctx context.Context, walletID string, limit, offset int) ([]model.WalletTransaction, error) {
	return s.repo.GetTransactions(ctx, walletID, limit, offset)
}
