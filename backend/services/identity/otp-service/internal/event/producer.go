package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"

	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/identity/otp-service/internal/ports"
)

const (
	topicOTPSent     = "otp.sent"
	topicOTPVerified = "otp.verified"
)

type otpEvent struct {
	EventID   string        `json:"event_id"`
	Phone     string        `json:"phone"`
	Purpose   model.Purpose `json:"purpose"`
	Timestamp time.Time     `json:"timestamp"`
}

type kafkaPublisher struct {
	producer sarama.SyncProducer
}

func NewKafkaPublisher(producer sarama.SyncProducer) ports.EventPublisher {
	return &kafkaPublisher{producer: producer}
}

func (p *kafkaPublisher) PublishOTPSent(ctx context.Context, otp *model.OTP) error {
	return p.publish(topicOTPSent, otp)
}

func (p *kafkaPublisher) PublishOTPVerified(ctx context.Context, otp *model.OTP) error {
	return p.publish(topicOTPVerified, otp)
}

func (p *kafkaPublisher) publish(topic string, otp *model.OTP) error {
	evt := otpEvent{
		EventID:   otp.ID,
		Phone:     otp.Phone,
		Purpose:   otp.Purpose,
		Timestamp: time.Now().UTC(),
	}

	payload, err := json.Marshal(evt)
	if err != nil {
		return fmt.Errorf("event: marshal: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(otp.Phone),
		Value: sarama.ByteEncoder(payload),
	}

	_, _, err = p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("event: send %s: %w", topic, err)
	}
	return nil
}
