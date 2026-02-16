package inventory

import "github.com/gin-gonic/gin"

type InventoryDeliveryInterface interface {
	GetAlerts() gin.HandlerFunc
	Create() gin.HandlerFunc
	List() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
