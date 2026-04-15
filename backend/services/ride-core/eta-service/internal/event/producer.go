package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/eta-service/internal/model"
)

type KafkaWriter interface {
	WriteMessages(ctx context.Context, topic string, key string, value []byte) error
}

type producer struct {
	writer KafkaWriter
}

func NewProducer(writer KafkaWriter) *producer {
	return &producer{writer: writer}
}

func (p *producer) PublishETACalculated(ctx context.Context, req *model.ETARequest, result *model.ETAResult) error {
	payload := map[string]any{
		"request": req,
		"result":  result,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal event payload: %w", err)
	}
	return p.writer.WriteMessages(ctx, constants.TopicETACalculated, req.CityID, data)
}
