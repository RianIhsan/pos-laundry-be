package http

import (
	"github.com/RianIhsan/pos-laundry-be/internal/features/services"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapServiceRoutes(
	serviceGroup *gin.RouterGroup,
	delivery services.ServiceDeliveryInterface,
	mw *middleware.MiddlewareManager) {

	protected := serviceGroup.Group("")
	protected.Use(mw.AuthMiddleware())

	protected.POST("/services", delivery.AddService())
	protected.PUT("/services/:id", delivery.Update())
	protected.DELETE("/services/:id", delivery.Delete())

	protected.GET("/services", delivery.GetList())
	protected.GET("/services/:id", delivery.GetById())
}
