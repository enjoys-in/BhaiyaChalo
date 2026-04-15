package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/constants"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/realtime-core/socket-session-service/internal/ports"
)

type KafkaWriter interface {
	WriteMessage(ctx context.Context, topic string, key string, value []byte) error
}

type Producer struct {
	writer KafkaWriter
}

func NewProducer(writer KafkaWriter) ports.EventPublisher {
	return &Producer{writer: writer}
}

type sessionEvent struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	DeviceID  string `json:"device_id"`
	ServerID  string `json:"server_id"`
	Timestamp string `json:"timestamp"`
}

func (p *Producer) PublishSessionConnected(ctx context.Context, session *model.SocketSession) error {
	return p.publish(ctx, constants.TopicSessionConnected, session)
}

func (p *Producer) PublishSessionDisconnected(ctx context.Context, session *model.SocketSession) error {
	return p.publish(ctx, constants.TopicSessionDisconnected, session)
}

func (p *Producer) publish(ctx context.Context, topic string, session *model.SocketSession) error {
	evt := sessionEvent{
		SessionID: session.ID,
		UserID:    session.UserID,
		Role:      session.Role,
		DeviceID:  session.DeviceID,
		ServerID:  session.ServerID,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	return p.writer.WriteMessage(ctx, topic, session.UserID, data)
}
