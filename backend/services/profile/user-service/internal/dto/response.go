package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID          uuid.UUID  `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	AvatarURL   string     `json:"avatar_url"`
	Gender      string     `json:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	CityID      string     `json:"city_id"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type AddressResponse struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"user_id"`
	Label            string    `json:"label"`
	Lat              float64   `json:"lat"`
	Lng              float64   `json:"lng"`
	FormattedAddress string    `json:"formatted_address"`
	IsDefault        bool      `json:"is_default"`
	CreatedAt        time.Time `json:"created_at"`
}

type EmergencyContactResponse struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Relation string    `json:"relation"`
}

type UserProfileResponse struct {
	User              UserResponse               `json:"user"`
	Addresses         []AddressResponse          `json:"addresses"`
	EmergencyContacts []EmergencyContactResponse `json:"emergency_contacts"`
}
