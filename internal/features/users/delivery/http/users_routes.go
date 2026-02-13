package http

import (
	"github.com/RianIhsan/pos-laundry-be/internal/features/users"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapUserRoutes(
	userGroup *gin.RouterGroup,
	delivery users.UserDeliveryInterface,
	mw *middleware.MiddlewareManager) {

	userGroup.POST("/register", delivery.RegisterUser())
	userGroup.POST("/login", delivery.LoginUser())

	protected := userGroup.Group("")
	protected.Use(mw.AuthMiddleware())

	//protected.GET("/users", delivery.GetList())
	//protected.GET("/users/:id", delivery.GetById())
	//protected.PUT("/users/protected", delivery.SelfUpdate())
	//protected.PUT("/users/avatar", delivery.UpdateAvatar())
	protected.GET("/users/me", delivery.GetMe())
	protected.PUT("/users/:id", delivery.Update())
	protected.DELETE("users/:id", delivery.DeleteUser())
}
