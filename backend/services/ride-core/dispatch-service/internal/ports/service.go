package ports

import (
	"context"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/dispatch-service/internal/model"
)

type DispatchService interface {
	Dispatch(ctx context.Context, req *dto.CreateDispatchRequest) (*model.DispatchOffer, error)
	HandleDriverResponse(ctx context.Context, req *dto.DriverResponseRequest) (*model.DispatchOffer, error)
	ExpireOffer(ctx context.Context, offerID string) error
}
