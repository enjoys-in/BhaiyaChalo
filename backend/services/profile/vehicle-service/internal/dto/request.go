package dto

import "time"

type CreateVehicleRequest struct {
	DriverID        string `json:"driver_id" validate:"required"`
	Make            string `json:"make" validate:"required"`
	Model           string `json:"model" validate:"required"`
	Year            int    `json:"year" validate:"required"`
	Color           string `json:"color" validate:"required"`
	PlateNumber     string `json:"plate_number" validate:"required"`
	VehicleType     string `json:"vehicle_type" validate:"required,oneof=auto mini sedan suv premium"`
	InsuranceExpiry time.Time `json:"insurance_expiry" validate:"required"`
	FitnessExpiry   time.Time `json:"fitness_expiry" validate:"required"`
}

type UpdateVehicleRequest struct {
	Make            string    `json:"make,omitempty"`
	Model           string    `json:"model,omitempty"`
	Year            int       `json:"year,omitempty"`
	Color           string    `json:"color,omitempty"`
	PlateNumber     string    `json:"plate_number,omitempty"`
	VehicleType     string    `json:"vehicle_type,omitempty"`
	InsuranceExpiry *time.Time `json:"insurance_expiry,omitempty"`
	FitnessExpiry   *time.Time `json:"fitness_expiry,omitempty"`
	Status          string    `json:"status,omitempty"`
}
