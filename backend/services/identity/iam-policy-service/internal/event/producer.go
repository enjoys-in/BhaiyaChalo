package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/iam-policy-service/internal/ports"
	"github.com/segmentio/kafka-go"
)

type kafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(brokers []string) ports.EventPublisher {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        constants.TopicIAMEvents,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	}
	return &kafkaPublisher{writer: w}
}

type eventMessage struct {
	Type      string    `json:"type"`
	Payload   any       `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func (p *kafkaPublisher) PublishRoleCreated(ctx context.Context, roleID, roleName string) error {
	return p.publish(ctx, constants.EventRoleCreated, map[string]string{
		"role_id":   roleID,
		"role_name": roleName,
	})
}

func (p *kafkaPublisher) PublishPolicyChanged(ctx context.Context, policyID, subjectType, subjectID string) error {
	return p.publish(ctx, constants.EventPolicyChanged, map[string]string{
		"policy_id":    policyID,
		"subject_type": subjectType,
		"subject_id":   subjectID,
	})
}

func (p *kafkaPublisher) publish(ctx context.Context, eventType string, payload any) error {
	msg := eventMessage{
		Type:      eventType,
		Payload:   payload,
		Timestamp: time.Now().UTC(),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(eventType),
		Value: data,
	})
}
