package model

import "time"

type OfferStatus string

const (
	OfferStatusPending  OfferStatus = "pending"
	OfferStatusAccepted OfferStatus = "accepted"
	OfferStatusRejected OfferStatus = "rejected"
	OfferStatusExpired  OfferStatus = "expired"
)

type RoundStatus string

const (
	RoundStatusActive    RoundStatus = "active"
	RoundStatusCompleted RoundStatus = "completed"
	RoundStatusFailed    RoundStatus = "failed"
)

type DispatchOffer struct {
	ID             string      `json:"id" db:"id"`
	BookingID      string      `json:"booking_id" db:"booking_id"`
	DriverID       string      `json:"driver_id" db:"driver_id"`
	CityID         string      `json:"city_id" db:"city_id"`
	Status         OfferStatus `json:"status" db:"status"`
	OfferExpiresAt time.Time   `json:"offer_expires_at" db:"offer_expires_at"`
	RespondedAt    *time.Time  `json:"responded_at,omitempty" db:"responded_at"`
	CreatedAt      time.Time   `json:"created_at" db:"created_at"`
}

type DispatchRound struct {
	ID                 string      `json:"id" db:"id"`
	BookingID          string      `json:"booking_id" db:"booking_id"`
	RoundNumber        int         `json:"round_number" db:"round_number"`
	CandidateDriverIDs []string    `json:"candidate_driver_ids" db:"candidate_driver_ids"`
	Status             RoundStatus `json:"status" db:"status"`
	CreatedAt          time.Time   `json:"created_at" db:"created_at"`
}
