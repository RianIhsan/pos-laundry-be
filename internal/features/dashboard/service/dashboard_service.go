package service

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/dashboard"
	"github.com/RianIhsan/pos-laundry-be/internal/features/dashboard/dto"
	"github.com/pkg/errors"
)

type dashboardService struct {
	cfg           *ServiceConfig
	dashboardRepo dashboard.DashboardRepositoryInterface
}

func NewDashboardService(cfg *ServiceConfig) dashboard.DashboardServiceInterface {
	return &dashboardService{
		cfg:           cfg,
		dashboardRepo: cfg.DashboardRepoInterface,
	}
}

func (s *dashboardService) GetStats(ctx context.Context) (dto.DashboardStatsResponse, error) {
	// Get summary data
	totalOmzet, err := s.dashboardRepo.GetTotalOmzet(ctx)
	if err != nil {
		return dto.DashboardStatsResponse{}, errors.Wrap(err, "failed to get total omzet")
	}

	newMembers, err := s.dashboardRepo.GetNewMembersCount(ctx)
	if err != nil {
		return dto.DashboardStatsResponse{}, errors.Wrap(err, "failed to get new members count")
	}

	activeProduction, err := s.dashboardRepo.GetActiveProductionCount(ctx)
	if err != nil {
		return dto.DashboardStatsResponse{}, errors.Wrap(err, "failed to get active production")
	}

	readyToPickup, err := s.dashboardRepo.GetReadyToPickupCount(ctx)
	if err != nil {
		return dto.DashboardStatsResponse{}, errors.Wrap(err, "failed to get ready to pickup")
	}

	unpaidRevenue, err := s.dashboardRepo.GetUnpaidRevenue(ctx)
	if err != nil {
		return dto.DashboardStatsResponse{}, errors.Wrap(err, "failed to get unpaid revenue")
	}

	// Get production queues
	transactions, err := s.dashboardRepo.GetProductionQueues(ctx, 10)
	if err != nil {
		return dto.DashboardStatsResponse{}, errors.Wrap(err, "failed to get production queues")
	}

	productionQueues := s.convertToProductionQueuesDTO(transactions)

	// Get popular services
	servicesData, err := s.dashboardRepo.GetPopularServices(ctx, 5)
	if err != nil {
		return dto.DashboardStatsResponse{}, errors.Wrap(err, "failed to get popular services")
	}

	popularServices := s.convertToPopularServicesDTO(servicesData)

	// Calculate total revenue for percentage calculation
	totalRevenue := 0.0
	for _, svc := range popularServices {
		totalRevenue += svc.TotalRevenue
	}

	// Calculate percentage
	for i := range popularServices {
		if totalRevenue > 0 {
			popularServices[i].Percentage = int32((popularServices[i].TotalRevenue / totalRevenue) * 100)
		}
	}

	summary := dto.DashboardSummary{
		TotalOmzet:       totalOmzet,
		OmzetTrend:       "+12.5%", // TODO: Calculate from previous day
		NewMembers:       newMembers,
		MemberTrend:      "+8%", // TODO: Calculate from previous day
		ActiveProduction: activeProduction,
		ReadyToPickupCnt: readyToPickup,
		UnpaidRevenue:    unpaidRevenue,
	}

	return dto.DashboardStatsResponse{
		Summary:          summary,
		ProductionQueues: productionQueues,
		PopularServices:  popularServices,
	}, nil
}

func (s *dashboardService) GetActivities(ctx context.Context) ([]dto.ActivityLogDTO, error) {
	logs, err := s.dashboardRepo.GetActivityLogs(ctx, 10)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get activity logs")
	}

	activities := make([]dto.ActivityLogDTO, 0)
	for _, log := range logs {
		activities = append(activities, dto.ActivityLogDTO{
			UserName: log.User.Name,
			Action:   log.Action,
			Target:   log.TargetID,
			TimeAgo:  s.getTimeAgo(log.CreatedAt),
		})
	}

	return activities, nil
}

func (s *dashboardService) GetActivityLogsPage(ctx context.Context, page int) (dto.ActivityLogListResponse, error) {
	const perPage = 10
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * perPage

	logs, err := s.dashboardRepo.GetActivityLogsPaginated(ctx, perPage, offset)
	if err != nil {
		return dto.ActivityLogListResponse{}, errors.Wrap(err, "failed to get activity logs paginated")
	}

	total, err := s.dashboardRepo.CountActivityLogs(ctx)
	if err != nil {
		return dto.ActivityLogListResponse{}, errors.Wrap(err, "failed to count activity logs")
	}

	items := make([]dto.ActivityLogItemDTO, 0)
	for _, log := range logs {
		items = append(items, dto.ActivityLogItemDTO{
			ID:          log.ID,
			UserName:    log.User.Name,
			Action:      log.Action,
			Target:      log.TargetID,
			TargetType:  log.TargetType,
			Description: log.Description,
			CreatedAt:   log.CreatedAt.Format(time.RFC3339),
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return dto.ActivityLogListResponse{
		Items:      items,
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *dashboardService) convertToProductionQueuesDTO(transactions []entities.Transaction) []dto.ProductionQueueDTO {
	queues := make([]dto.ProductionQueueDTO, 0)
	for _, tx := range transactions {
		// Calculate total qty from items
		totalQty := 0.0
		for _, item := range tx.Items {
			totalQty += item.Qty
		}

		queueDTO := dto.ProductionQueueDTO{
			ID:           tx.InvoiceNo,
			CustomerName: tx.Customer.Name,
			OrderStatus:  tx.OrderStatus,
			CreatedAt:    tx.CreatedAt,
		}

		// Get service name from first item
		if len(tx.Items) > 0 {
			queueDTO.ServiceName = tx.Items[0].Service.Name
			// Format qty display
			if tx.Items[0].Service.Unit != "" {
				queueDTO.QtyDisplay = fmt.Sprintf("%.1f %s", totalQty, tx.Items[0].Service.Unit)
			} else {
				queueDTO.QtyDisplay = fmt.Sprintf("%.1f", totalQty)
			}
		}

		queues = append(queues, queueDTO)
	}
	return queues
}

func (s *dashboardService) convertToPopularServicesDTO(servicesData []map[string]interface{}) []dto.PopularServiceDTO {
	services := make([]dto.PopularServiceDTO, 0)
	for _, data := range servicesData {
		service := dto.PopularServiceDTO{
			Name:         fmt.Sprintf("%v", data["name"]),
			TotalRevenue: s.getFloat64(data["total_revenue"]),
		}
		if totalOrders, ok := data["total_orders"].(int64); ok {
			service.TotalOrders = totalOrders
		}
		services = append(services, service)
	}
	return services
}

func (s *dashboardService) getTimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	if diff.Minutes() < 1 {
		return "Baru saja"
	} else if diff.Minutes() < 60 {
		return fmt.Sprintf("%.0f menit yang lalu", diff.Minutes())
	} else if diff.Hours() < 24 {
		return fmt.Sprintf("%.0f jam yang lalu", diff.Hours())
	} else if diff.Hours() < 168 {
		return fmt.Sprintf("%.0f hari yang lalu", diff.Hours()/24)
	}
	return t.Format("2006-01-02")
}

func (s *dashboardService) getFloat64(val interface{}) float64 {
	switch v := val.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	default:
		return 0
	}
}
