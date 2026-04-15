package dto

import "time"

type DriverResponse struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	AvatarURL      string    `json:"avatar_url"`
	LicenseNumber  string    `json:"license_number"`
	CityID         string    `json:"city_id"`
	Rating         float64   `json:"rating"`
	TotalTrips     int       `json:"total_trips"`
	Status         string    `json:"status"`
	OnboardingStep int       `json:"onboarding_step"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type DriverLocationResponse struct {
	DriverID  string    `json:"driver_id"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Heading   float64   `json:"heading"`
	Speed     float64   `json:"speed"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DriverPreferenceResponse struct {
	DriverID       string   `json:"driver_id"`
	AutoAccept     bool     `json:"auto_accept"`
	MaxDistance    float64  `json:"max_distance"`
	PreferredZones []string `json:"preferred_zones"`
}
