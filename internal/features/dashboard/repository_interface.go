package dashboard

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type DashboardRepositoryInterface interface {
	GetTotalOmzet(ctx context.Context) (float64, error)
	GetNewMembersCount(ctx context.Context) (int64, error)
	GetActiveProductionCount(ctx context.Context) (int64, error)
	GetReadyToPickupCount(ctx context.Context) (int64, error)
	GetUnpaidRevenue(ctx context.Context) (float64, error)
	GetProductionQueues(ctx context.Context, limit int) ([]entities.Transaction, error)
	GetPopularServices(ctx context.Context, limit int) ([]map[string]interface{}, error)
	GetActivityLogs(ctx context.Context, limit int) ([]entities.ActivityLog, error)
	GetActivityLogsPaginated(ctx context.Context, limit int, offset int) ([]entities.ActivityLog, error)
	CountActivityLogs(ctx context.Context) (int64, error)
}
