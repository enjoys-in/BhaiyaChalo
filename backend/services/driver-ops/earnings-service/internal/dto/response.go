package dto

import "time"

type EarningResponse struct {
	ID             string    `json:"id"`
	DriverID       string    `json:"driver_id"`
	TripID         string    `json:"trip_id"`
	FareAmount     float64   `json:"fare_amount"`
	Commission     float64   `json:"commission"`
	IncentiveBonus float64   `json:"incentive_bonus"`
	TipAmount      float64   `json:"tip_amount"`
	NetEarning     float64   `json:"net_earning"`
	Currency       string    `json:"currency"`
	EarnedAt       time.Time `json:"earned_at"`
}

type DailySummaryResponse struct {
	DriverID        string    `json:"driver_id"`
	Date            time.Time `json:"date"`
	TotalTrips      int       `json:"total_trips"`
	TotalFare       float64   `json:"total_fare"`
	TotalCommission float64   `json:"total_commission"`
	TotalIncentive  float64   `json:"total_incentive"`
	TotalTips       float64   `json:"total_tips"`
	NetEarning      float64   `json:"net_earning"`
}

type WeeklySummaryResponse struct {
	DriverID        string    `json:"driver_id"`
	WeekStart       time.Time `json:"week_start"`
	WeekEnd         time.Time `json:"week_end"`
	TotalTrips      int       `json:"total_trips"`
	TotalFare       float64   `json:"total_fare"`
	TotalCommission float64   `json:"total_commission"`
	TotalIncentive  float64   `json:"total_incentive"`
	TotalTips       float64   `json:"total_tips"`
	NetEarning      float64   `json:"net_earning"`
}
