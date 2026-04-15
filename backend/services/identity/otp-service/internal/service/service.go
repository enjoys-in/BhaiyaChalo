package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/config"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/ports"

	"github.com/google/uuid"
)

var (
	ErrRateLimited       = errors.New("otp: rate limited, try again later")
	ErrOTPExpired        = errors.New("otp: code has expired")
	ErrOTPInvalid        = errors.New("otp: invalid code")
	ErrMaxAttemptsExceed = errors.New("otp: max verification attempts exceeded")
	ErrOTPNotFound       = errors.New("otp: not found")
)

type otpService struct {
	repo      ports.OTPRepository
	publisher ports.EventPublisher
	redis     *redis.Client
	cfg       *config.Config
}

func NewOTPService(
	repo ports.OTPRepository,
	publisher ports.EventPublisher,
	rdb *redis.Client,
	cfg *config.Config,
) ports.OTPService {
	return &otpService{
		repo:      repo,
		publisher: publisher,
		redis:     rdb,
		cfg:       cfg,
	}
}

func (s *otpService) Send(ctx context.Context, req dto.SendOTPRequest) (*dto.SendOTPResponse, error) {
	rateLimitKey := fmt.Sprintf("otp:rate:%s:%s", req.Phone, req.Purpose)

	ttl, err := s.redis.TTL(ctx, rateLimitKey).Result()
	if err != nil {
		return nil, fmt.Errorf("otp: redis ttl check: %w", err)
	}
	if ttl > 0 {
		return nil, ErrRateLimited
	}

	code, err := generateCode(s.cfg.OTPLength)
	if err != nil {
		return nil, fmt.Errorf("otp: generate code: %w", err)
	}

	now := time.Now().UTC()
	expiry := now.Add(time.Duration(s.cfg.OTPExpirySeconds) * time.Second)

	otp := &model.OTP{
		ID:        uuid.New().String(),
		Phone:     req.Phone,
		Code:      code,
		Purpose:   req.Purpose,
		Verified:  false,
		Attempts:  0,
		ExpiresAt: expiry,
		CreatedAt: now,
	}

	if err := s.repo.Create(ctx, otp); err != nil {
		return nil, fmt.Errorf("otp: create: %w", err)
	}

	cooldown := 60 * time.Second
	if err := s.redis.Set(ctx, rateLimitKey, "1", cooldown).Err(); err != nil {
		return nil, fmt.Errorf("otp: redis set rate limit: %w", err)
	}

	if err := s.publisher.PublishOTPSent(ctx, otp); err != nil {
		return nil, fmt.Errorf("otp: publish sent event: %w", err)
	}

	return &dto.SendOTPResponse{
		ExpiresAt:  expiry.Unix(),
		RetryAfter: int(cooldown.Seconds()),
	}, nil
}

func (s *otpService) Verify(ctx context.Context, req dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, error) {
	otp, err := s.repo.FindLatest(ctx, req.Phone, req.Purpose)
	if err != nil {
		return nil, fmt.Errorf("otp: find latest: %w", err)
	}
	if otp == nil {
		return nil, ErrOTPNotFound
	}

	if otp.Verified {
		return &dto.VerifyOTPResponse{Verified: false}, ErrOTPInvalid
	}

	if time.Now().UTC().After(otp.ExpiresAt) {
		return &dto.VerifyOTPResponse{Verified: false}, ErrOTPExpired
	}

	if otp.Attempts >= constants.MaxOTPAttempts {
		return &dto.VerifyOTPResponse{Verified: false}, ErrMaxAttemptsExceed
	}

	if err := s.repo.IncrementAttempts(ctx, otp.ID); err != nil {
		return nil, fmt.Errorf("otp: increment attempts: %w", err)
	}

	if otp.Code != req.Code {
		return &dto.VerifyOTPResponse{Verified: false}, ErrOTPInvalid
	}

	if err := s.repo.MarkVerified(ctx, otp.ID); err != nil {
		return nil, fmt.Errorf("otp: mark verified: %w", err)
	}

	if err := s.publisher.PublishOTPVerified(ctx, otp); err != nil {
		return nil, fmt.Errorf("otp: publish verified event: %w", err)
	}

	return &dto.VerifyOTPResponse{Verified: true}, nil
}

func generateCode(length int) (string, error) {
	max := new(big.Int)
	max.Exp(big.NewInt(10), big.NewInt(int64(length)), nil)

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0*d", length, n), nil
}
