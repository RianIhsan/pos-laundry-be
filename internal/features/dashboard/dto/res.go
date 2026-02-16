package dto

import "time"

type DashboardSummary struct {
	TotalOmzet       float64 `json:"total_omzet"`
	OmzetTrend       string  `json:"omzet_trend"`
	NewMembers       int64   `json:"new_members"`
	MemberTrend      string  `json:"member_trend"`
	ActiveProduction int64   `json:"active_production"`
	ReadyToPickupCnt int64   `json:"ready_to_pickup_count"`
	UnpaidRevenue    float64 `json:"unpaid_revenue"`
}

type ProductionQueueDTO struct {
	ID           string    `json:"id"`
	CustomerName string    `json:"customer_name"`
	ServiceName  string    `json:"service_name"`
	QtyDisplay   string    `json:"qty_display"`
	OrderStatus  string    `json:"order_status"`
	CreatedAt    time.Time `json:"created_at"`
}

type PopularServiceDTO struct {
	Name         string  `json:"name"`
	TotalOrders  int64   `json:"total_orders"`
	TotalRevenue float64 `json:"total_revenue"`
	Percentage   int32   `json:"percentage"`
}

type DashboardStatsResponse struct {
	Summary          DashboardSummary     `json:"summary"`
	ProductionQueues []ProductionQueueDTO `json:"production_queues"`
	PopularServices  []PopularServiceDTO  `json:"popular_services"`
}

type ActivityLogDTO struct {
	UserName string `json:"user_name"`
	Action   string `json:"action"`
	Target   string `json:"target"`
	TimeAgo  string `json:"time_ago"`
}

type ActivityLogItemDTO struct {
	ID          uint   `json:"id"`
	UserName    string `json:"user_name"`
	Action      string `json:"action"`
	Target      string `json:"target"`
	TargetType  string `json:"target_type"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type ActivityLogListResponse struct {
	Items      []ActivityLogItemDTO `json:"items"`
	Page       int                  `json:"page"`
	PerPage    int                  `json:"per_page"`
	Total      int64                `json:"total"`
	TotalPages int                  `json:"total_pages"`
}
