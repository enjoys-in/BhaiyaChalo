package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/ports"
)

type authService struct {
	repo      ports.AuthRepository
	publisher ports.EventPublisher
	jwtSecret string
}

func NewAuthService(repo ports.AuthRepository, publisher ports.EventPublisher, jwtSecret string) ports.AuthService {
	return &authService{repo: repo, publisher: publisher, jwtSecret: jwtSecret}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	role := model.Role(req.Role)
	user, err := s.repo.FindUserByPhone(ctx, req.Phone, role)
	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}

	if user == nil {
		user = &model.User{
			ID:        generateID(),
			Phone:     req.Phone,
			Role:      role,
			Status:    model.StatusActive,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := s.repo.CreateUser(ctx, user); err != nil {
			return nil, fmt.Errorf("create user: %w", err)
		}
	}

	accessToken := generateToken()
	refreshToken := generateToken()
	expiresAt := time.Now().Add(24 * time.Hour)

	token := &model.Token{
		ID:           generateID(),
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
	}
	if err := s.repo.StoreRefreshToken(ctx, token); err != nil {
		return nil, fmt.Errorf("store refresh token: %w", err)
	}

	_ = s.publisher.PublishUserLoggedIn(ctx, user.ID, string(user.Role))

	return &dto.LoginResponse{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}

func (s *authService) Refresh(ctx context.Context, req dto.RefreshRequest) (*dto.TokenResponse, error) {
	token, err := s.repo.FindRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("find refresh token: %w", err)
	}
	if token == nil {
		return nil, fmt.Errorf("invalid or expired refresh token")
	}

	newAccess := generateToken()
	newRefresh := generateToken()
	expiresAt := time.Now().Add(24 * time.Hour)

	_ = s.repo.RevokeRefreshToken(ctx, token.UserID)

	newToken := &model.Token{
		ID:           generateID(),
		UserID:       token.UserID,
		RefreshToken: newRefresh,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
	}
	if err := s.repo.StoreRefreshToken(ctx, newToken); err != nil {
		return nil, fmt.Errorf("store new refresh token: %w", err)
	}

	return &dto.TokenResponse{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
		ExpiresAt:    expiresAt,
	}, nil
}

func (s *authService) Verify(ctx context.Context, req dto.VerifyRequest) (*dto.VerifyResponse, error) {
	// TODO: JWT validation with s.jwtSecret
	return &dto.VerifyResponse{Valid: true}, nil
}

func (s *authService) Logout(ctx context.Context, req dto.LogoutRequest) error {
	if err := s.repo.RevokeRefreshToken(ctx, req.UserID); err != nil {
		return fmt.Errorf("revoke tokens: %w", err)
	}
	_ = s.publisher.PublishUserLoggedOut(ctx, req.UserID)
	return nil
}

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
