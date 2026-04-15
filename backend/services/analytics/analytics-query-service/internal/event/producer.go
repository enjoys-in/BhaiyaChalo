package event

import (
	"log/slog"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/analytics-query-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
}

func NewKafkaPublisher(logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{logger: logger}
}
