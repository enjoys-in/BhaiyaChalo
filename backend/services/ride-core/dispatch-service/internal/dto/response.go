package dto

import "time"

type DispatchResponse struct {
	OfferID   string `json:"offer_id"`
	BookingID string `json:"booking_id"`
	DriverID  string `json:"driver_id"`
	Status    string `json:"status"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
}

type OfferStatusResponse struct {
	OfferID     string     `json:"offer_id"`
	Status      string     `json:"status"`
	RespondedAt *time.Time `json:"responded_at,omitempty"`
}

type DispatchRoundResponse struct {
	RoundID     string   `json:"round_id"`
	BookingID   string   `json:"booking_id"`
	RoundNumber int      `json:"round_number"`
	DriverIDs   []string `json:"driver_ids"`
	Status      string   `json:"status"`
}
