package http

import (
	"github.com/RianIhsan/pos-laundry-be/internal/features/inventory"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapInventoryRoutes(
	inventoryGroup *gin.RouterGroup,
	delivery inventory.InventoryDeliveryInterface,
	mw *middleware.MiddlewareManager) {

	protected := inventoryGroup.Group("")
	protected.Use(mw.AuthMiddleware())
	// Inventory CRUD
	protected.GET("/inventory", delivery.List())
	protected.POST("/inventory", delivery.Create())
	protected.GET("/inventory/:id", delivery.GetByID())
	protected.PUT("/inventory/:id", delivery.Update())
	protected.DELETE("/inventory/:id", delivery.Delete())

	// Alerts
	protected.GET("/inventory/alerts", delivery.GetAlerts())
}
