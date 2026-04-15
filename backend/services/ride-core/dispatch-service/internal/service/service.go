package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/config"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/ports"
)

type dispatchService struct {
	repo      ports.DispatchRepository
	publisher ports.DispatchEventPublisher
	cfg       *config.Config
}

func NewDispatchService(
	repo ports.DispatchRepository,
	publisher ports.DispatchEventPublisher,
	cfg *config.Config,
) ports.DispatchService {
	return &dispatchService{
		repo:      repo,
		publisher: publisher,
		cfg:       cfg,
	}
}

func (s *dispatchService) Dispatch(ctx context.Context, req *dto.CreateDispatchRequest) (*model.DispatchOffer, error) {
	rounds, err := s.repo.FindRoundsByBooking(ctx, req.BookingID)
	if err != nil {
		return nil, fmt.Errorf("find rounds: %w", err)
	}

	roundNumber := len(rounds) + 1
	if roundNumber > s.cfg.MaxRetries {
		if pubErr := s.publisher.PublishDispatchFailed(ctx, req.BookingID, "max retries exceeded"); pubErr != nil {
			return nil, fmt.Errorf("publish dispatch failed: %w", pubErr)
		}
		return nil, fmt.Errorf("max dispatch retries exceeded for booking %s", req.BookingID)
	}

	candidates := req.DriverIDs
	if len(candidates) > constants.MaxCandidatesPerRound {
		candidates = candidates[:constants.MaxCandidatesPerRound]
	}

	round := &model.DispatchRound{
		ID:                 uuid.New().String(),
		BookingID:          req.BookingID,
		RoundNumber:        roundNumber,
		CandidateDriverIDs: candidates,
		Status:             model.RoundStatusActive,
		CreatedAt:          time.Now().UTC(),
	}
	if err := s.repo.CreateRound(ctx, round); err != nil {
		return nil, fmt.Errorf("create round: %w", err)
	}

	var lastOffer *model.DispatchOffer
	for _, driverID := range candidates {
		offer := &model.DispatchOffer{
			ID:             uuid.New().String(),
			BookingID:      req.BookingID,
			DriverID:       driverID,
			CityID:         req.CityID,
			Status:         model.OfferStatusPending,
			OfferExpiresAt: time.Now().UTC().Add(time.Duration(s.cfg.OfferTimeoutSeconds) * time.Second),
			CreatedAt:      time.Now().UTC(),
		}
		if err := s.repo.CreateOffer(ctx, offer); err != nil {
			return nil, fmt.Errorf("create offer for driver %s: %w", driverID, err)
		}
		if err := s.publisher.PublishOfferSent(ctx, offer); err != nil {
			return nil, fmt.Errorf("publish offer sent: %w", err)
		}
		lastOffer = offer
	}

	return lastOffer, nil
}

func (s *dispatchService) HandleDriverResponse(ctx context.Context, req *dto.DriverResponseRequest) (*model.DispatchOffer, error) {
	offer, err := s.repo.FindOfferByID(ctx, req.OfferID)
	if err != nil {
		return nil, fmt.Errorf("find offer: %w", err)
	}

	if offer.Status != model.OfferStatusPending {
		return nil, fmt.Errorf("offer %s is no longer pending (status: %s)", offer.ID, offer.Status)
	}

	if time.Now().UTC().After(offer.OfferExpiresAt) {
		if err := s.ExpireOffer(ctx, offer.ID); err != nil {
			return nil, fmt.Errorf("expire offer: %w", err)
		}
		return nil, fmt.Errorf("offer %s has expired", offer.ID)
	}

	if req.Accepted {
		return s.acceptOffer(ctx, offer)
	}
	return s.rejectOffer(ctx, offer)
}

func (s *dispatchService) acceptOffer(ctx context.Context, offer *model.DispatchOffer) (*model.DispatchOffer, error) {
	if err := s.repo.UpdateOfferStatus(ctx, offer.ID, model.OfferStatusAccepted); err != nil {
		return nil, fmt.Errorf("update offer status: %w", err)
	}
	offer.Status = model.OfferStatusAccepted
	now := time.Now().UTC()
	offer.RespondedAt = &now

	pending, err := s.repo.FindPendingByBooking(ctx, offer.BookingID)
	if err != nil {
		return offer, nil
	}
	for _, p := range pending {
		if p.ID != offer.ID {
			_ = s.repo.UpdateOfferStatus(ctx, p.ID, model.OfferStatusExpired)
		}
	}

	if err := s.publisher.PublishOfferAccepted(ctx, offer); err != nil {
		return offer, fmt.Errorf("publish accepted: %w", err)
	}
	return offer, nil
}

func (s *dispatchService) rejectOffer(ctx context.Context, offer *model.DispatchOffer) (*model.DispatchOffer, error) {
	if err := s.repo.UpdateOfferStatus(ctx, offer.ID, model.OfferStatusRejected); err != nil {
		return nil, fmt.Errorf("update offer status: %w", err)
	}
	offer.Status = model.OfferStatusRejected
	now := time.Now().UTC()
	offer.RespondedAt = &now

	if err := s.publisher.PublishOfferRejected(ctx, offer); err != nil {
		return offer, fmt.Errorf("publish rejected: %w", err)
	}
	return offer, nil
}

func (s *dispatchService) ExpireOffer(ctx context.Context, offerID string) error {
	offer, err := s.repo.FindOfferByID(ctx, offerID)
	if err != nil {
		return fmt.Errorf("find offer: %w", err)
	}

	if offer.Status != model.OfferStatusPending {
		return nil
	}

	if err := s.repo.UpdateOfferStatus(ctx, offerID, model.OfferStatusExpired); err != nil {
		return fmt.Errorf("update offer status: %w", err)
	}

	offer.Status = model.OfferStatusExpired
	return s.publisher.PublishOfferExpired(ctx, offer)
}
