package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/stop-planning-service/internal/ports"
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

type stopEvent struct {
	StopID    string  `json:"stop_id"`
	TripID    string  `json:"trip_id"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Address   string  `json:"address"`
	StopOrder int     `json:"stop_order"`
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
}

func (p *Producer) PublishStopAdded(ctx context.Context, stop *model.Stop) error {
	return p.publish(ctx, constants.TopicStopAdded, stop)
}

func (p *Producer) PublishStopArrived(ctx context.Context, stop *model.Stop) error {
	return p.publish(ctx, constants.TopicStopArrived, stop)
}

func (p *Producer) PublishStopCompleted(ctx context.Context, stop *model.Stop) error {
	return p.publish(ctx, constants.TopicStopCompleted, stop)
}

func (p *Producer) publish(ctx context.Context, topic string, stop *model.Stop) error {
	evt := stopEvent{
		StopID:    stop.ID,
		TripID:    stop.TripID,
		Lat:       stop.Lat,
		Lng:       stop.Lng,
		Address:   stop.Address,
		StopOrder: stop.StopOrder,
		Status:    string(stop.Status),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return p.writer.WriteMessage(ctx, topic, stop.TripID, data)
}
