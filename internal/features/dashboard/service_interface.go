package dashboard

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/features/dashboard/dto"
)

type DashboardServiceInterface interface {
	GetStats(ctx context.Context) (dto.DashboardStatsResponse, error)
	GetActivities(ctx context.Context) ([]dto.ActivityLogDTO, error)
	GetActivityLogsPage(ctx context.Context, page int) (dto.ActivityLogListResponse, error)
}
