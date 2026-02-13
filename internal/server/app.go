package server

import (
	usersDelivery "github.com/RianIhsan/pos-laundry-be/internal/features/users/delivery/http"
	usersRepo "github.com/RianIhsan/pos-laundry-be/internal/features/users/repository"
	usersService "github.com/RianIhsan/pos-laundry-be/internal/features/users/service"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) Bootstrap() error {

	// Initialize Repositories
	usersPostgres := usersRepo.NewUserPostgresRepository(s.db)

	// Initialize Services
	usersSvc := usersService.NewUserService(&usersService.ServiceConfig{
		UserRepoInterface: usersPostgres,
		Logger:            s.logger,
		Config:            s.cfg,
	})

	// intialize delivery
	usersDel := usersDelivery.NewUserDelivery(&usersDelivery.DeliveryConfig{
		UserServiceInterface: usersSvc,
		Config:               s.cfg,
		Logger:               s.logger,
	})

	// initialize Middleware
	middlewareManager := middleware.NewMiddlewareManager(&middleware.MiddlewareConfig{
		Logger: s.logger,
		Config: s.cfg,
	})

	s.app.Use(middlewareManager.RequestIdMiddleware())
	s.app.Use(middlewareManager.RequestLoggerMiddleware())

	// intialize routes
	api := s.app.Group("/api")
	{
		userGroupV1 := api.Group("/v1")
		{
			usersDelivery.MapUserRoutes(userGroupV1, usersDel, middlewareManager)
		}
	}

	// Ping route
	s.app.GET("/ping", func(ctxCtx *gin.Context) {
		ctxCtx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return nil
}
