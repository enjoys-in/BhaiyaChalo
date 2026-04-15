package ports

import (
	"context"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/model"
)

type EarningsRepository interface {
	Record(ctx context.Context, earning *model.Earning) error
	FindByDriverID(ctx context.Context, driverID string, from, to time.Time) ([]*model.Earning, error)
	GetDailySummary(ctx context.Context, driverID string, date time.Time) (*model.DailySummary, error)
	GetWeeklySummary(ctx context.Context, driverID string, weekStart time.Time) (*model.WeeklySummary, error)
	GetTotalByDateRange(ctx context.Context, driverID string, from, to time.Time) (float64, error)
}
