package http

import (
	"github.com/RianIhsan/pos-laundry-be/internal/features/dashboard"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapDashboardRoutes(
	dashboardGroup *gin.RouterGroup,
	delivery dashboard.DashboardDeliveryInterface,
	mw *middleware.MiddlewareManager) {

	protected := dashboardGroup.Group("")
	protected.Use(mw.AuthMiddleware())

	protected.GET("/dashboard/stats", delivery.GetStats())
	protected.GET("/dashboard/activities", delivery.GetActivities())
	protected.GET("/dashboard/activity-logs", delivery.GetActivityLogs())
}
