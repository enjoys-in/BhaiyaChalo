package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/ports"
)

type kafkaPublisher struct {
	brokers string
}

func NewKafkaPublisher(brokers string) ports.EventPublisher {
	return &kafkaPublisher{brokers: brokers}
}

type eventEnvelope struct {
	Type      string      `json:"type"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func (p *kafkaPublisher) PublishSearchPerformed(ctx context.Context, query *model.SearchQuery) error {
	envelope := eventEnvelope{
		Type:      constants.EventSearchPerformed,
		Payload:   query,
		Timestamp: time.Now().UTC(),
	}

	data, err := json.Marshal(envelope)
	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}

	// TODO: integrate with actual Kafka producer
	_ = data
	_ = ctx

	return nil
}
