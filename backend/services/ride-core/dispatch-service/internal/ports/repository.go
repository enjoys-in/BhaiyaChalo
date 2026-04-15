package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/model"
)

type DispatchRepository interface {
	CreateOffer(ctx context.Context, offer *model.DispatchOffer) error
	FindOfferByID(ctx context.Context, offerID string) (*model.DispatchOffer, error)
	UpdateOfferStatus(ctx context.Context, offerID string, status model.OfferStatus) error
	FindPendingByBooking(ctx context.Context, bookingID string) ([]*model.DispatchOffer, error)
	CreateRound(ctx context.Context, round *model.DispatchRound) error
	FindRoundsByBooking(ctx context.Context, bookingID string) ([]*model.DispatchRound, error)
}
