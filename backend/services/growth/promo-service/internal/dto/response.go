package dto

import "time"

type PromoResponse struct {
	ID            string  `json:"id"`
	Code          string  `json:"code"`
	CityID        string  `json:"city_id"`
	Type          string  `json:"type"`
	DiscountValue float64 `json:"discount_value"`
	MaxDiscount   float64 `json:"max_discount"`
	MinOrderValue float64 `json:"min_order_value"`
	UsageLimit    int     `json:"usage_limit"`
	UsedCount     int     `json:"used_count"`
	ValidFrom     string  `json:"valid_from"`
	ValidUntil    string  `json:"valid_until"`
	Active        bool    `json:"active"`
	CreatedAt     string  `json:"created_at"`
}

type ApplyPromoResponse struct {
	Valid          bool    `json:"valid"`
	DiscountAmount float64 `json:"discount_amount,omitempty"`
	FinalAmount    float64 `json:"final_amount,omitempty"`
	Message        string  `json:"message,omitempty"`
}

type PromoListResponse struct {
	Promos []PromoResponse `json:"promos"`
	Total  int             `json:"total"`
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
