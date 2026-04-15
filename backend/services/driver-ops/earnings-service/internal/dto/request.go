package dto

import "time"

type RecordEarningRequest struct {
	DriverID       string  `json:"driver_id" validate:"required"`
	TripID         string  `json:"trip_id" validate:"required"`
	FareAmount     float64 `json:"fare_amount" validate:"required,gt=0"`
	IncentiveBonus float64 `json:"incentive_bonus"`
	TipAmount      float64 `json:"tip_amount"`
	Currency       string  `json:"currency" validate:"required"`
}

type GetEarningsRequest struct {
	DriverID string    `json:"driver_id" validate:"required"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

type GetSummaryRequest struct {
	DriverID string    `json:"driver_id" validate:"required"`
	Date     time.Time `json:"date"`
}
