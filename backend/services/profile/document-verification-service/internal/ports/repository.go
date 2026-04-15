package ports

import (
	"context"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/model"
)

type DocumentRepository interface {
	Create(ctx context.Context, doc *model.Document) error
	FindByID(ctx context.Context, id string) (*model.Document, error)
	FindByOwner(ctx context.Context, ownerID string, ownerType model.OwnerType) ([]*model.Document, error)
	UpdateStatus(ctx context.Context, id string, status model.DocumentStatus, reviewerID, reviewNote string) error
	ListPending(ctx context.Context, limit, offset int) ([]*model.Document, error)
	ListExpiring(ctx context.Context, before time.Time) ([]*model.Document, error)
}
