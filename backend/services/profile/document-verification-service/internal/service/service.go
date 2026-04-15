package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/model"
	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/ports"
)

type documentService struct {
	repo      ports.DocumentRepository
	publisher ports.EventPublisher
}

func NewDocumentService(repo ports.DocumentRepository, publisher ports.EventPublisher) ports.DocumentService {
	return &documentService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *documentService) Upload(ctx context.Context, req dto.UploadDocumentRequest) (*dto.DocumentResponse, error) {
	now := time.Now().UTC()

	var expiresAt *time.Time
	if req.ExpiresAt != "" {
		t, err := time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			return nil, fmt.Errorf("invalid expires_at format: %w", err)
		}
		expiresAt = &t
	}

	doc := &model.Document{
		ID:        uuid.NewString(),
		OwnerID:   req.OwnerID,
		OwnerType: model.OwnerType(req.OwnerType),
		DocType:   model.DocType(req.DocType),
		FileURL:   req.FileURL,
		Status:    model.DocumentStatusPending,
		ExpiresAt: expiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(ctx, doc); err != nil {
		return nil, fmt.Errorf("create document: %w", err)
	}

	_ = s.publisher.PublishDocumentUploaded(ctx, doc)

	return toDocumentResponse(doc), nil
}

func (s *documentService) GetByID(ctx context.Context, id string) (*dto.DocumentResponse, error) {
	doc, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find document: %w", err)
	}
	if doc == nil {
		return nil, fmt.Errorf("document not found")
	}
	return toDocumentResponse(doc), nil
}

func (s *documentService) GetByOwner(ctx context.Context, ownerID, ownerType string) ([]dto.DocumentResponse, error) {
	docs, err := s.repo.FindByOwner(ctx, ownerID, model.OwnerType(ownerType))
	if err != nil {
		return nil, fmt.Errorf("find documents by owner: %w", err)
	}
	return toDocumentResponseList(docs), nil
}

func (s *documentService) Review(ctx context.Context, req dto.ReviewDocumentRequest) (*dto.VerificationResponse, error) {
	doc, err := s.repo.FindByID(ctx, req.DocumentID)
	if err != nil {
		return nil, fmt.Errorf("find document: %w", err)
	}
	if doc == nil {
		return nil, fmt.Errorf("document not found")
	}

	status := model.DocumentStatus(req.Status)
	if err := s.repo.UpdateStatus(ctx, req.DocumentID, status, req.ReviewerID, req.ReviewNote); err != nil {
		return nil, fmt.Errorf("update document status: %w", err)
	}

	doc.Status = status
	doc.ReviewerID = req.ReviewerID
	doc.ReviewNote = req.ReviewNote

	switch status {
	case model.DocumentStatusApproved:
		_ = s.publisher.PublishDocumentApproved(ctx, doc)
	case model.DocumentStatusRejected:
		_ = s.publisher.PublishDocumentRejected(ctx, doc)
	}

	now := time.Now().UTC()
	return &dto.VerificationResponse{
		DocumentID: doc.ID,
		Status:     string(status),
		ReviewerID: req.ReviewerID,
		ReviewNote: req.ReviewNote,
		VerifiedAt: now,
	}, nil
}

func (s *documentService) ListPending(ctx context.Context, page, perPage int) ([]dto.DocumentResponse, error) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 20
	}
	offset := (page - 1) * perPage

	docs, err := s.repo.ListPending(ctx, perPage, offset)
	if err != nil {
		return nil, fmt.Errorf("list pending documents: %w", err)
	}
	return toDocumentResponseList(docs), nil
}

func (s *documentService) ListExpiring(ctx context.Context, daysAhead int) ([]dto.DocumentResponse, error) {
	if daysAhead <= 0 {
		daysAhead = 30
	}
	before := time.Now().UTC().AddDate(0, 0, daysAhead)

	docs, err := s.repo.ListExpiring(ctx, before)
	if err != nil {
		return nil, fmt.Errorf("list expiring documents: %w", err)
	}
	return toDocumentResponseList(docs), nil
}

func toDocumentResponse(d *model.Document) *dto.DocumentResponse {
	return &dto.DocumentResponse{
		ID:         d.ID,
		OwnerID:    d.OwnerID,
		OwnerType:  string(d.OwnerType),
		DocType:    string(d.DocType),
		FileURL:    d.FileURL,
		Status:     string(d.Status),
		ReviewerID: d.ReviewerID,
		ReviewNote: d.ReviewNote,
		ExpiresAt:  d.ExpiresAt,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
	}
}

func toDocumentResponseList(docs []*model.Document) []dto.DocumentResponse {
	resp := make([]dto.DocumentResponse, len(docs))
	for i, d := range docs {
		resp[i] = *toDocumentResponse(d)
	}
	return resp
}
