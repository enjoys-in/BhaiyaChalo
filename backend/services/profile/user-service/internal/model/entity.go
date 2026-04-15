package model

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusInactive  UserStatus = "inactive"
	UserStatusSuspended UserStatus = "suspended"
)

type User struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    string     `json:"last_name" db:"last_name"`
	Phone       string     `json:"phone" db:"phone"`
	Email       string     `json:"email" db:"email"`
	AvatarURL   string     `json:"avatar_url" db:"avatar_url"`
	Gender      string     `json:"gender" db:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth" db:"date_of_birth"`
	CityID      string     `json:"city_id" db:"city_id"`
	Status      UserStatus `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type AddressLabel string

const (
	AddressLabelHome  AddressLabel = "home"
	AddressLabelWork  AddressLabel = "work"
	AddressLabelOther AddressLabel = "other"
)

type Address struct {
	ID               uuid.UUID    `json:"id" db:"id"`
	UserID           uuid.UUID    `json:"user_id" db:"user_id"`
	Label            AddressLabel `json:"label" db:"label"`
	Lat              float64      `json:"lat" db:"lat"`
	Lng              float64      `json:"lng" db:"lng"`
	FormattedAddress string       `json:"formatted_address" db:"formatted_address"`
	IsDefault        bool         `json:"is_default" db:"is_default"`
	CreatedAt        time.Time    `json:"created_at" db:"created_at"`
}

type EmergencyContact struct {
	ID       uuid.UUID `json:"id" db:"id"`
	UserID   uuid.UUID `json:"user_id" db:"user_id"`
	Name     string    `json:"name" db:"name"`
	Phone    string    `json:"phone" db:"phone"`
	Relation string    `json:"relation" db:"relation"`
}
