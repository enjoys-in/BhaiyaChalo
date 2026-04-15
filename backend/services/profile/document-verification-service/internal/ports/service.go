package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/dto"
)

type DocumentService interface {
	Upload(ctx context.Context, req dto.UploadDocumentRequest) (*dto.DocumentResponse, error)
	GetByID(ctx context.Context, id string) (*dto.DocumentResponse, error)
	GetByOwner(ctx context.Context, ownerID, ownerType string) ([]dto.DocumentResponse, error)
	Review(ctx context.Context, req dto.ReviewDocumentRequest) (*dto.VerificationResponse, error)
	ListPending(ctx context.Context, page, perPage int) ([]dto.DocumentResponse, error)
	ListExpiring(ctx context.Context, daysAhead int) ([]dto.DocumentResponse, error)
}
