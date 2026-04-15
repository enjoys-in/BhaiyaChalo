package dto

import "time"

type CreateUserRequest struct {
	FirstName   string     `json:"first_name" validate:"required,min=1,max=100"`
	LastName    string     `json:"last_name" validate:"required,min=1,max=100"`
	Phone       string     `json:"phone" validate:"required,e164"`
	Email       string     `json:"email" validate:"omitempty,email"`
	AvatarURL   string     `json:"avatar_url" validate:"omitempty,url"`
	Gender      string     `json:"gender" validate:"omitempty,oneof=male female other"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	CityID      string     `json:"city_id" validate:"omitempty"`
}

type UpdateUserRequest struct {
	FirstName   *string    `json:"first_name" validate:"omitempty,min=1,max=100"`
	LastName    *string    `json:"last_name" validate:"omitempty,min=1,max=100"`
	Email       *string    `json:"email" validate:"omitempty,email"`
	AvatarURL   *string    `json:"avatar_url" validate:"omitempty,url"`
	Gender      *string    `json:"gender" validate:"omitempty,oneof=male female other"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	CityID      *string    `json:"city_id"`
}

type AddAddressRequest struct {
	Label            string  `json:"label" validate:"required,oneof=home work other"`
	Lat              float64 `json:"lat" validate:"required,latitude"`
	Lng              float64 `json:"lng" validate:"required,longitude"`
	FormattedAddress string  `json:"formatted_address" validate:"required,min=1,max=500"`
	IsDefault        bool    `json:"is_default"`
}

type UpdateAddressRequest struct {
	Label            *string  `json:"label" validate:"omitempty,oneof=home work other"`
	Lat              *float64 `json:"lat" validate:"omitempty,latitude"`
	Lng              *float64 `json:"lng" validate:"omitempty,longitude"`
	FormattedAddress *string  `json:"formatted_address" validate:"omitempty,min=1,max=500"`
	IsDefault        *bool    `json:"is_default"`
}

type AddEmergencyContactRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Phone    string `json:"phone" validate:"required,e164"`
	Relation string `json:"relation" validate:"required,min=1,max=50"`
}
