package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/route-planning-service/internal/model"
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

type routePlannedEvent struct {
	RouteID     string  `json:"route_id"`
	BookingID   string  `json:"booking_id"`
	DistanceKM  float64 `json:"distance_km"`
	DurationMin float64 `json:"duration_min"`
	Polyline    string  `json:"polyline"`
	Timestamp   string  `json:"timestamp"`
}

func (p *Producer) PublishRoutePlanned(ctx context.Context, route *model.Route) error {
	evt := routePlannedEvent{
		RouteID:     route.ID,
		BookingID:   route.BookingID,
		DistanceKM:  route.DistanceKM,
		DurationMin: route.DurationMin,
		Polyline:    route.Polyline,
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	}

	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return p.writer.WriteMessage(ctx, constants.TopicRoutePlanned, route.BookingID, data)
}
