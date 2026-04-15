package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessages(ctx context.Context, topic string, key string, value []byte) error
}

type producer struct {
	writer KafkaWriter
}

func NewProducer(writer KafkaWriter) ports.DispatchEventPublisher {
	return &producer{writer: writer}
}

func (p *producer) PublishOfferSent(ctx context.Context, offer *model.DispatchOffer) error {
	return p.publish(ctx, constants.TopicOfferSent, offer.BookingID, offer)
}

func (p *producer) PublishOfferAccepted(ctx context.Context, offer *model.DispatchOffer) error {
	return p.publish(ctx, constants.TopicOfferAccepted, offer.BookingID, offer)
}

func (p *producer) PublishOfferRejected(ctx context.Context, offer *model.DispatchOffer) error {
	return p.publish(ctx, constants.TopicOfferRejected, offer.BookingID, offer)
}

func (p *producer) PublishOfferExpired(ctx context.Context, offer *model.DispatchOffer) error {
	return p.publish(ctx, constants.TopicOfferExpired, offer.BookingID, offer)
}

func (p *producer) PublishDispatchFailed(ctx context.Context, bookingID string, reason string) error {
	payload := map[string]string{
		"booking_id": bookingID,
		"reason":     reason,
	}
	return p.publish(ctx, constants.TopicDispatchFailed, bookingID, payload)
}

func (p *producer) publish(ctx context.Context, topic, key string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal event payload: %w", err)
	}
	return p.writer.WriteMessages(ctx, topic, key, data)
}
