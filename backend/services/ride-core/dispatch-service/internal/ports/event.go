package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/model"
)

type DispatchEventPublisher interface {
	PublishOfferSent(ctx context.Context, offer *model.DispatchOffer) error
	PublishOfferAccepted(ctx context.Context, offer *model.DispatchOffer) error
	PublishOfferRejected(ctx context.Context, offer *model.DispatchOffer) error
	PublishOfferExpired(ctx context.Context, offer *model.DispatchOffer) error
	PublishDispatchFailed(ctx context.Context, bookingID string, reason string) error
}
