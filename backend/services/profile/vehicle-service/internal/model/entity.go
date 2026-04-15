package model

import "time"

type VehicleType string

const (
	VehicleTypeAuto    VehicleType = "auto"
	VehicleTypeMini    VehicleType = "mini"
	VehicleTypeSedan   VehicleType = "sedan"
	VehicleTypeSUV     VehicleType = "suv"
	VehicleTypePremium VehicleType = "premium"
)

type VehicleStatus string

const (
	VehicleStatusPending  VehicleStatus = "pending"
	VehicleStatusApproved VehicleStatus = "approved"
	VehicleStatusRejected VehicleStatus = "rejected"
	VehicleStatusExpired  VehicleStatus = "expired"
)

type Vehicle struct {
	ID              string        `json:"id" db:"id"`
	DriverID        string        `json:"driver_id" db:"driver_id"`
	Make            string        `json:"make" db:"make"`
	Model           string        `json:"model" db:"model"`
	Year            int           `json:"year" db:"year"`
	Color           string        `json:"color" db:"color"`
	PlateNumber     string        `json:"plate_number" db:"plate_number"`
	VehicleType     VehicleType   `json:"vehicle_type" db:"vehicle_type"`
	InsuranceExpiry time.Time     `json:"insurance_expiry" db:"insurance_expiry"`
	FitnessExpiry   time.Time     `json:"fitness_expiry" db:"fitness_expiry"`
	Status          VehicleStatus `json:"status" db:"status"`
	CreatedAt       time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" db:"updated_at"`
}
