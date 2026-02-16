package services

import "github.com/gin-gonic/gin"

type ServiceDeliveryInterface interface {
	AddService() gin.HandlerFunc
	GetList() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
