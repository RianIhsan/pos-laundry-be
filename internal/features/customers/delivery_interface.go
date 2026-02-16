package customers

import "github.com/gin-gonic/gin"

type CustomerDeliveryInterface interface {
	AddCustomer() gin.HandlerFunc
	GetList() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
