package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/model"
)

type EventPublisher interface {
	PublishSearchPerformed(ctx context.Context, query *model.SearchQuery) error
}
