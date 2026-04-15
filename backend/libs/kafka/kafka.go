package kafka

import (
	"context"
	"log/slog"
)

// ProducerConfig holds Kafka producer settings.
type ProducerConfig struct {
	Brokers []string
	Topic   string
}

// Producer wraps a Kafka producer.
type Producer struct {
	cfg    ProducerConfig
	logger *slog.Logger
}

// NewProducer creates a Kafka producer.
func NewProducer(cfg ProducerConfig, logger *slog.Logger) (*Producer, error) {
	if len(cfg.Brokers) == 0 {
		cfg.Brokers = []string{"localhost:9092"}
	}
	return &Producer{cfg: cfg, logger: logger}, nil
}

// Publish sends a message to the configured topic.
func (p *Producer) Publish(ctx context.Context, key string, value []byte) error {
	p.logger.Info("kafka publish", "topic", p.cfg.Topic, "key", key)
	// TODO: actual kafka produce when dependency is added
	return nil
}

// Close shuts down the producer.
func (p *Producer) Close() error { return nil }

// ConsumerConfig holds Kafka consumer settings.
type ConsumerConfig struct {
	Brokers []string
	Topic   string
	GroupID string
}

// Handler is the function signature for processing a Kafka message.
type Handler func(ctx context.Context, key string, value []byte) error

// Consumer wraps a Kafka consumer group.
type Consumer struct {
	cfg     ConsumerConfig
	handler Handler
	logger  *slog.Logger
}

// NewConsumer creates a Kafka consumer.
func NewConsumer(cfg ConsumerConfig, handler Handler, logger *slog.Logger) (*Consumer, error) {
	return &Consumer{cfg: cfg, handler: handler, logger: logger}, nil
}

// Start begins consuming messages. Blocks until context is cancelled.
func (c *Consumer) Start(ctx context.Context) error {
	c.logger.Info("kafka consumer started", "topic", c.cfg.Topic, "group", c.cfg.GroupID)
	<-ctx.Done()
	return ctx.Err()
}

// Close shuts down the consumer.
func (c *Consumer) Close() error { return nil }
