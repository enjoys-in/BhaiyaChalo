package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/support/escalation-service/internal/model"
)

type EventPublisher interface {
	PublishEscalationCreated(ctx context.Context, entity *model.Escalation) error
	PublishEscalationResolved(ctx context.Context, entity *model.Escalation) error
}
