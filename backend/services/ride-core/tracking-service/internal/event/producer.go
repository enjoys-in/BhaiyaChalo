package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/tracking-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessages(ctx context.Context, topic string, key string, value []byte) error
}

type producer struct {
	writer KafkaWriter
}

func NewProducer(writer KafkaWriter) ports.TrackingEventPublisher {
	return &producer{writer: writer}
}

func (p *producer) PublishLocationUpdated(ctx context.Context, loc *model.LocationUpdate) error {
	return p.publish(ctx, constants.TopicLocationUpdated, loc.DriverID, loc)
}

func (p *producer) PublishTrackingStarted(ctx context.Context, session *model.TrackingSession) error {
	return p.publish(ctx, constants.TopicTrackingStarted, session.TripID, session)
}

func (p *producer) PublishTrackingStopped(ctx context.Context, session *model.TrackingSession) error {
	return p.publish(ctx, constants.TopicTrackingStopped, session.TripID, session)
}

func (p *producer) publish(ctx context.Context, topic, key string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal event payload: %w", err)
	}
	return p.writer.WriteMessages(ctx, topic, key, data)
}
