package http

import (
	"github.com/RianIhsan/pos-laundry-be/internal/features/transactions"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapTransactionRoutes(
	transactionGroup *gin.RouterGroup,
	delivery transactions.TransactionDeliveryInterface,
	mw *middleware.MiddlewareManager) {

	protected := transactionGroup.Group("")
	protected.Use(mw.AuthMiddleware())

	protected.POST("/transactions", delivery.CreateTransaction())
	protected.GET("/transactions", delivery.GetList())
	protected.GET("/transactions/:id", delivery.GetById())
	protected.PUT("/transactions/:id/status", delivery.UpdateStatus())
}
