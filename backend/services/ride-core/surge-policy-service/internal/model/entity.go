package model

import "time"

type SurgeZone struct {
	ID                string    `json:"id" db:"id"`
	CityID            string    `json:"city_id" db:"city_id"`
	GeofenceID        string    `json:"geofence_id" db:"geofence_id"`
	CurrentMultiplier float64   `json:"current_multiplier" db:"current_multiplier"`
	DemandCount       int       `json:"demand_count" db:"demand_count"`
	SupplyCount       int       `json:"supply_count" db:"supply_count"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

type SurgePolicy struct {
	ID                   string    `json:"id" db:"id"`
	CityID               string    `json:"city_id" db:"city_id"`
	MinDemandSupplyRatio float64   `json:"min_demand_supply_ratio" db:"min_demand_supply_ratio"`
	MaxMultiplier        float64   `json:"max_multiplier" db:"max_multiplier"`
	StepSize             float64   `json:"step_size" db:"step_size"`
	CooldownMinutes      int       `json:"cooldown_minutes" db:"cooldown_minutes"`
	Active               bool      `json:"active" db:"active"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
}

type SurgeHistory struct {
	ID           string    `json:"id" db:"id"`
	ZoneID       string    `json:"zone_id" db:"zone_id"`
	Multiplier   float64   `json:"multiplier" db:"multiplier"`
	DemandCount  int       `json:"demand_count" db:"demand_count"`
	SupplyCount  int       `json:"supply_count" db:"supply_count"`
	CalculatedAt time.Time `json:"calculated_at" db:"calculated_at"`
}
