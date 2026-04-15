package service

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/analytics/audit-log-service/internal/ports"
)

type auditLogService struct {
	repo      ports.AuditLogRepository
	publisher ports.EventPublisher
}

func NewAuditLogService(repo ports.AuditLogRepository, publisher ports.EventPublisher) ports.AuditLogService {
	return &auditLogService{repo: repo, publisher: publisher}
}

func (s *auditLogService) Create(ctx context.Context, req dto.CreateAuditEntryRequest) (*dto.AuditEntryResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *auditLogService) GetByID(ctx context.Context, id string) (*dto.AuditEntryResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *auditLogService) Update(ctx context.Context, req dto.UpdateAuditEntryRequest) (*dto.AuditEntryResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *auditLogService) Delete(ctx context.Context, id string) error {
	// TODO: implement
	return nil
}
