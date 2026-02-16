package dashboard

import "github.com/gin-gonic/gin"

type DashboardDeliveryInterface interface {
	GetStats() gin.HandlerFunc
	GetActivities() gin.HandlerFunc
	GetActivityLogs() gin.HandlerFunc
}
