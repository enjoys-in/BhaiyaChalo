package ports

import (
	"context"
	"time"

	"github.com/enjoys-in/BhaiyaChalo/services/driver-ops/earnings-service/internal/dto"
)

type EarningsService interface {
	RecordEarning(ctx context.Context, req dto.RecordEarningRequest) (*dto.EarningResponse, error)
	GetEarnings(ctx context.Context, driverID string, from, to time.Time) ([]dto.EarningResponse, error)
	GetDailySummary(ctx context.Context, driverID string, date time.Time) (*dto.DailySummaryResponse, error)
	GetWeeklySummary(ctx context.Context, driverID string, weekStart time.Time) (*dto.WeeklySummaryResponse, error)
}
