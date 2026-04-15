package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/matching-service/internal/model"
)

type MatchRepository interface {
	SaveRequest(ctx context.Context, req *model.MatchRequest) error
	UpdateStatus(ctx context.Context, id string, status model.MatchStatus) error
	FindByBookingID(ctx context.Context, bookingID string) (*model.MatchRequest, error)
}
