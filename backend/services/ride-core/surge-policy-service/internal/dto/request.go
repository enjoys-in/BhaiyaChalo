package dto

type CalculateSurgeRequest struct {
	CityID      string `json:"city_id" validate:"required"`
	ZoneID      string `json:"zone_id" validate:"required"`
	DemandCount int    `json:"demand_count" validate:"required,gte=0"`
	SupplyCount int    `json:"supply_count" validate:"required,gte=0"`
}

type UpdatePolicyRequest struct {
	MinDemandSupplyRatio *float64 `json:"min_demand_supply_ratio,omitempty"`
	MaxMultiplier        *float64 `json:"max_multiplier,omitempty"`
	StepSize             *float64 `json:"step_size,omitempty"`
	CooldownMinutes      *int     `json:"cooldown_minutes,omitempty"`
	Active               *bool    `json:"active,omitempty"`
}
