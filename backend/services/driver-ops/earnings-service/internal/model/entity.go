package model

import "time"

type Earning struct {
	ID             string    `json:"id" db:"id"`
	DriverID       string    `json:"driver_id" db:"driver_id"`
	TripID         string    `json:"trip_id" db:"trip_id"`
	FareAmount     float64   `json:"fare_amount" db:"fare_amount"`
	Commission     float64   `json:"commission" db:"commission"`
	IncentiveBonus float64   `json:"incentive_bonus" db:"incentive_bonus"`
	TipAmount      float64   `json:"tip_amount" db:"tip_amount"`
	NetEarning     float64   `json:"net_earning" db:"net_earning"`
	Currency       string    `json:"currency" db:"currency"`
	EarnedAt       time.Time `json:"earned_at" db:"earned_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type DailySummary struct {
	DriverID        string    `json:"driver_id" db:"driver_id"`
	Date            time.Time `json:"date" db:"date"`
	TotalTrips      int       `json:"total_trips" db:"total_trips"`
	TotalFare       float64   `json:"total_fare" db:"total_fare"`
	TotalCommission float64   `json:"total_commission" db:"total_commission"`
	TotalIncentive  float64   `json:"total_incentive" db:"total_incentive"`
	TotalTips       float64   `json:"total_tips" db:"total_tips"`
	NetEarning      float64   `json:"net_earning" db:"net_earning"`
}

type WeeklySummary struct {
	DriverID        string    `json:"driver_id" db:"driver_id"`
	WeekStart       time.Time `json:"week_start" db:"week_start"`
	WeekEnd         time.Time `json:"week_end" db:"week_end"`
	TotalTrips      int       `json:"total_trips" db:"total_trips"`
	TotalFare       float64   `json:"total_fare" db:"total_fare"`
	TotalCommission float64   `json:"total_commission" db:"total_commission"`
	TotalIncentive  float64   `json:"total_incentive" db:"total_incentive"`
	TotalTips       float64   `json:"total_tips" db:"total_tips"`
	NetEarning      float64   `json:"net_earning" db:"net_earning"`
}
