package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/dto"
)

type AuditLogService interface {
	Create(ctx context.Context, req dto.CreateAuditEntryRequest) (*dto.AuditEntryResponse, error)
	GetByID(ctx context.Context, id string) (*dto.AuditEntryResponse, error)
	Update(ctx context.Context, req dto.UpdateAuditEntryRequest) (*dto.AuditEntryResponse, error)
	Delete(ctx context.Context, id string) error
}
