package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/model"
)

type AuthRepository interface {
	FindUserByPhone(ctx context.Context, phone string, role model.Role) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	StoreRefreshToken(ctx context.Context, token *model.Token) error
	FindRefreshToken(ctx context.Context, refreshToken string) (*model.Token, error)
	RevokeRefreshToken(ctx context.Context, userID string) error
}
