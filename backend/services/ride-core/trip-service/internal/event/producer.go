package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/trip-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, topic string, key string, value []byte) error
}

type Producer struct {
	writer KafkaWriter
}

func NewProducer(writer KafkaWriter) ports.EventPublisher {
	return &Producer{writer: writer}
}

type tripEvent struct {
	TripID      string  `json:"trip_id"`
	BookingID   string  `json:"booking_id"`
	UserID      string  `json:"user_id"`
	DriverID    string  `json:"driver_id"`
	Status      string  `json:"status"`
	FareAmount  float64 `json:"fare_amount"`
	DistanceKM  float64 `json:"distance_km"`
	DurationMin float64 `json:"duration_min"`
	Timestamp   string  `json:"timestamp"`
}

func (p *Producer) PublishTripStarted(ctx context.Context, trip *model.Trip) error {
	return p.publish(ctx, constants.TopicTripStarted, trip)
}

func (p *Producer) PublishTripCompleted(ctx context.Context, trip *model.Trip) error {
	return p.publish(ctx, constants.TopicTripCompleted, trip)
}

func (p *Producer) PublishTripCancelled(ctx context.Context, trip *model.Trip) error {
	return p.publish(ctx, constants.TopicTripCancelled, trip)
}

func (p *Producer) publish(ctx context.Context, topic string, trip *model.Trip) error {
	evt := tripEvent{
		TripID:      trip.ID,
		BookingID:   trip.BookingID,
		UserID:      trip.UserID,
		DriverID:    trip.DriverID,
		Status:      string(trip.Status),
		FareAmount:  trip.FareAmount,
		DistanceKM:  trip.DistanceKM,
		DurationMin: trip.DurationMin,
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	}

	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return p.writer.WriteMessage(ctx, topic, trip.ID, data)
}
