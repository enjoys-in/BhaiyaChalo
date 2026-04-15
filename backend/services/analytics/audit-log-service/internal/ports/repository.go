package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/model"
)

type AuditLogRepository interface {
	Create(ctx context.Context, entity *model.AuditEntry) error
	FindByID(ctx context.Context, id string) (*model.AuditEntry, error)
	Update(ctx context.Context, entity *model.AuditEntry) error
	Delete(ctx context.Context, id string) error
}
