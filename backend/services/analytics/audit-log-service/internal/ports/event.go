package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/model"
)

type EventPublisher interface {
	PublishAuditLogged(ctx context.Context, entity *model.AuditEntry) error
}
