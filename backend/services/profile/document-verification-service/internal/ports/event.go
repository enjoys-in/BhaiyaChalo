package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/profile/document-verification-service/internal/model"
)

type EventPublisher interface {
	PublishDocumentUploaded(ctx context.Context, doc *model.Document) error
	PublishDocumentApproved(ctx context.Context, doc *model.Document) error
	PublishDocumentRejected(ctx context.Context, doc *model.Document) error
}
