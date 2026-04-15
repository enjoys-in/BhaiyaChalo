package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/auth-service/internal/ports"
)

type kafkaPublisher struct {
	logger *slog.Logger
	topic  string
}

func NewKafkaPublisher(topic string, logger *slog.Logger) ports.EventPublisher {
	return &kafkaPublisher{topic: topic, logger: logger}
}

type authEvent struct {
	Type      string    `json:"type"`
	UserID    string    `json:"user_id"`
	Role      string    `json:"role,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

func (p *kafkaPublisher) PublishUserLoggedIn(ctx context.Context, userID string, role string) error {
	evt := authEvent{Type: "user.logged_in", UserID: userID, Role: role, Timestamp: time.Now()}
	data, _ := json.Marshal(evt)
	p.logger.Info("event published", "type", evt.Type, "user_id", userID, "topic", p.topic)
	_ = data // TODO: push to actual Kafka producer
	return nil
}

func (p *kafkaPublisher) PublishUserLoggedOut(ctx context.Context, userID string) error {
	evt := authEvent{Type: "user.logged_out", UserID: userID, Timestamp: time.Now()}
	data, _ := json.Marshal(evt)
	p.logger.Info("event published", "type", evt.Type, "user_id", userID, "topic", p.topic)
	_ = data
	return nil
}
