package model

import "time"

type Driver struct {
	ID             string    `json:"id" db:"id"`
	FirstName      string    `json:"first_name" db:"first_name"`
	LastName       string    `json:"last_name" db:"last_name"`
	Phone          string    `json:"phone" db:"phone"`
	Email          string    `json:"email" db:"email"`
	AvatarURL      string    `json:"avatar_url" db:"avatar_url"`
	LicenseNumber  string    `json:"license_number" db:"license_number"`
	CityID         string    `json:"city_id" db:"city_id"`
	Rating         float64   `json:"rating" db:"rating"`
	TotalTrips     int       `json:"total_trips" db:"total_trips"`
	Status         string    `json:"status" db:"status"`
	OnboardingStep int       `json:"onboarding_step" db:"onboarding_step"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type DriverLocation struct {
	DriverID  string    `json:"driver_id" db:"driver_id"`
	Lat       float64   `json:"lat" db:"lat"`
	Lng       float64   `json:"lng" db:"lng"`
	Heading   float64   `json:"heading" db:"heading"`
	Speed     float64   `json:"speed" db:"speed"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type DriverPreference struct {
	DriverID       string   `json:"driver_id" db:"driver_id"`
	AutoAccept     bool     `json:"auto_accept" db:"auto_accept"`
	MaxDistance    float64  `json:"max_distance" db:"max_distance"`
	PreferredZones []string `json:"preferred_zones" db:"preferred_zones"`
}
