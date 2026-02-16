package transactions

import "github.com/gin-gonic/gin"

type TransactionDeliveryInterface interface {
	CreateTransaction() gin.HandlerFunc
	GetList() gin.HandlerFunc
	GetById() gin.HandlerFunc
	UpdateStatus() gin.HandlerFunc
}
