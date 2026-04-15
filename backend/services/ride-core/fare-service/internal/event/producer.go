package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/fare-service/internal/model"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, topic string, key string, value []byte) error
}

type Producer struct {
	writer KafkaWriter
}

func NewProducer(writer KafkaWriter) *Producer {
	return &Producer{writer: writer}
}

type fareCalculatedEvent struct {
	FareID          string  `json:"fare_id"`
	BookingID       string  `json:"booking_id"`
	TotalFare       float64 `json:"total_fare"`
	Currency        string  `json:"currency"`
	SurgeMultiplier float64 `json:"surge_multiplier"`
	CityID          string  `json:"city_id"`
	VehicleType     string  `json:"vehicle_type"`
	Timestamp       string  `json:"timestamp"`
}

func (p *Producer) PublishFareCalculated(ctx context.Context, calc *model.FareCalculation) error {
	evt := fareCalculatedEvent{
		FareID:          calc.ID,
		BookingID:       calc.BookingID,
		TotalFare:       calc.TotalFare,
		Currency:        calc.Currency,
		SurgeMultiplier: calc.SurgeMultiplier,
		CityID:          calc.CityID,
		VehicleType:     calc.VehicleType,
		Timestamp:       time.Now().UTC().Format(time.RFC3339),
	}

	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return p.writer.WriteMessage(ctx, constants.TopicFareCalculated, calc.BookingID, data)
}
