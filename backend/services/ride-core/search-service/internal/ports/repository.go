package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/model"
)

type SearchRepository interface {
	SaveQuery(ctx context.Context, query *model.SearchQuery) error
	FindByID(ctx context.Context, id string) (*model.SearchQuery, error)
}
