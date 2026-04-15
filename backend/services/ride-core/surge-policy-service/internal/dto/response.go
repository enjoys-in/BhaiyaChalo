package dto

import "time"

type SurgeResponse struct {
	ZoneID            string    `json:"zone_id"`
	CityID            string    `json:"city_id"`
	CurrentMultiplier float64   `json:"current_multiplier"`
	DemandCount       int       `json:"demand_count"`
	SupplyCount       int       `json:"supply_count"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type SurgePolicyResponse struct {
	ID                   string    `json:"id"`
	CityID               string    `json:"city_id"`
	MinDemandSupplyRatio float64   `json:"min_demand_supply_ratio"`
	MaxMultiplier        float64   `json:"max_multiplier"`
	StepSize             float64   `json:"step_size"`
	CooldownMinutes      int       `json:"cooldown_minutes"`
	Active               bool      `json:"active"`
	CreatedAt            time.Time `json:"created_at"`
}

type SurgeHistoryResponse struct {
	ID           string    `json:"id"`
	ZoneID       string    `json:"zone_id"`
	Multiplier   float64   `json:"multiplier"`
	DemandCount  int       `json:"demand_count"`
	SupplyCount  int       `json:"supply_count"`
	CalculatedAt time.Time `json:"calculated_at"`
}
