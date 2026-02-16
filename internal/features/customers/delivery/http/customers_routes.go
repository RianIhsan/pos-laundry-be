package http

import (
	"github.com/RianIhsan/pos-laundry-be/internal/features/customers"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapCustomerRoutes(
	customerGroup *gin.RouterGroup,
	delivery customers.CustomerDeliveryInterface,
	mw *middleware.MiddlewareManager) {

	
	protected := customerGroup.Group("")
	protected.Use(mw.AuthMiddleware())

	protected.POST("/customers", delivery.AddCustomer())
	protected.GET("/customers", delivery.GetList())
	protected.GET("/customers/:id", delivery.GetById())

	protected.PUT("/customers/:id", delivery.Update())
	protected.DELETE("/customers/:id", delivery.Delete())
}
