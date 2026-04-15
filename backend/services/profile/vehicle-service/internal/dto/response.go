package dto

import "time"

type VehicleResponse struct {
	ID              string    `json:"id"`
	DriverID        string    `json:"driver_id"`
	Make            string    `json:"make"`
	Model           string    `json:"model"`
	Year            int       `json:"year"`
	Color           string    `json:"color"`
	PlateNumber     string    `json:"plate_number"`
	VehicleType     string    `json:"vehicle_type"`
	InsuranceExpiry time.Time `json:"insurance_expiry"`
	FitnessExpiry   time.Time `json:"fitness_expiry"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
