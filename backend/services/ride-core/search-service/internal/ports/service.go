package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/search-service/internal/dto"
)

type SearchService interface {
	Search(ctx context.Context, req *dto.SearchRequest) (*dto.SearchResponse, error)
}
