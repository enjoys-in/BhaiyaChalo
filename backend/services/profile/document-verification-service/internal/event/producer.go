package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/ports"
)

type message struct {
	EventType string      `json:"event_type"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

type KafkaWriter interface {
	WriteMessages(ctx context.Context, topic string, key string, value []byte) error
}

type kafkaPublisher struct {
	writer KafkaWriter
	topic  string
}

func NewKafkaPublisher(writer KafkaWriter, topic string) ports.EventPublisher {
	return &kafkaPublisher{
		writer: writer,
		topic:  topic,
	}
}

func (p *kafkaPublisher) PublishDocumentUploaded(ctx context.Context, doc *model.Document) error {
	return p.publish(ctx, constants.EventDocumentUploaded, doc.OwnerID, doc)
}

func (p *kafkaPublisher) PublishDocumentApproved(ctx context.Context, doc *model.Document) error {
	return p.publish(ctx, constants.EventDocumentApproved, doc.OwnerID, doc)
}

func (p *kafkaPublisher) PublishDocumentRejected(ctx context.Context, doc *model.Document) error {
	return p.publish(ctx, constants.EventDocumentRejected, doc.OwnerID, doc)
}

func (p *kafkaPublisher) publish(ctx context.Context, eventType string, key string, payload interface{}) error {
	msg := message{
		EventType: eventType,
		Payload:   payload,
		Timestamp: time.Now().UTC(),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, p.topic, key, data)
}
