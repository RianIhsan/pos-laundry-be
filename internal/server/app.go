package server

import (
	usersDelivery "github.com/RianIhsan/pos-laundry-be/internal/features/users/delivery/http"
	usersRepo "github.com/RianIhsan/pos-laundry-be/internal/features/users/repository"
	usersService "github.com/RianIhsan/pos-laundry-be/internal/features/users/service"
	"github.com/RianIhsan/pos-laundry-be/pkg/activitylogger"

	customersDelivery "github.com/RianIhsan/pos-laundry-be/internal/features/customers/delivery/http"
	customersRepo "github.com/RianIhsan/pos-laundry-be/internal/features/customers/repository"
	customersService "github.com/RianIhsan/pos-laundry-be/internal/features/customers/service"

	servicesDelivery "github.com/RianIhsan/pos-laundry-be/internal/features/services/delivery/http"
	servicesRepo "github.com/RianIhsan/pos-laundry-be/internal/features/services/repository"
	servicesService "github.com/RianIhsan/pos-laundry-be/internal/features/services/service"

	transactionsDelivery "github.com/RianIhsan/pos-laundry-be/internal/features/transactions/delivery/http"
	transactionsRepo "github.com/RianIhsan/pos-laundry-be/internal/features/transactions/repository"
	transactionsService "github.com/RianIhsan/pos-laundry-be/internal/features/transactions/service"

	dashboardDelivery "github.com/RianIhsan/pos-laundry-be/internal/features/dashboard/delivery/http"
	dashboardRepo "github.com/RianIhsan/pos-laundry-be/internal/features/dashboard/repository"
	dashboardService "github.com/RianIhsan/pos-laundry-be/internal/features/dashboard/service"

	inventoryDelivery "github.com/RianIhsan/pos-laundry-be/internal/features/inventory/delivery/http"
	inventoryRepo "github.com/RianIhsan/pos-laundry-be/internal/features/inventory/repository"
	inventoryService "github.com/RianIhsan/pos-laundry-be/internal/features/inventory/service"

	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) Bootstrap() error {

	// Initialize Activity Logger
	activityLog := activitylogger.NewActivityLogger(s.db)

	// Initialize Repositories
	usersPostgres := usersRepo.NewUserPostgresRepository(s.db)
	customersPostgres := customersRepo.NewCustomerPostgresRepository(s.db)
	servicesPostgres := servicesRepo.NewServicePostgresRepository(s.db)
	transactionsPostgres := transactionsRepo.NewTransactionPostgresRepository(s.db)
	dashboardPostgres := dashboardRepo.NewDashboardPostgresRepository(s.db)
	inventoryPostgres := inventoryRepo.NewInventoryPostgresRepository(s.db)

	// Initialize Services
	usersSvc := usersService.NewUserService(&usersService.ServiceConfig{
		UserRepoInterface: usersPostgres,
		Logger:            s.logger,
		Config:            s.cfg,
	})

	customersSvc := customersService.NewCustomerService(&customersService.ServiceConfig{
		CustomerRepoInterface: customersPostgres,
		Logger:                s.logger,
		Config:                s.cfg,
		ActivityLogger:        activityLog,
	})

	servicesSvc := servicesService.NewServiceService(&servicesService.ServiceConfig{
		ServiceRepoInterface: servicesPostgres,
		Logger:               s.logger,
		Config:               s.cfg,
	})

	transactionsSvc := transactionsService.NewTransactionService(&transactionsService.ServiceConfig{
		TransactionRepoInterface: transactionsPostgres,
		Logger:                   s.logger,
		Config:                   s.cfg,
		ActivityLogger:           activityLog,
	})

	dashboardSvc := dashboardService.NewDashboardService(&dashboardService.ServiceConfig{
		DashboardRepoInterface: dashboardPostgres,
		Logger:                 s.logger,
		Config:                 s.cfg,
	})

	inventorySvc := inventoryService.NewInventoryService(&inventoryService.ServiceConfig{
		InventoryRepoInterface: inventoryPostgres,
		Logger:                 s.logger,
		Config:                 s.cfg,
		ActivityLogger:         activityLog,
	})

	// intialize delivery
	usersDel := usersDelivery.NewUserDelivery(&usersDelivery.DeliveryConfig{
		UserServiceInterface: usersSvc,
		Config:               s.cfg,
		Logger:               s.logger,
	})
	customersDel := customersDelivery.NewCustomerDelivery(&customersDelivery.DeliveryConfig{
		CustomerServiceInterface: customersSvc,
		Config:                   s.cfg,
		Logger:                   s.logger,
	})

	servicesDel := servicesDelivery.NewServiceDelivery(&servicesDelivery.DeliveryConfig{
		ServiceServiceInterface: servicesSvc,
		Config:                  s.cfg,
		Logger:                  s.logger,
	})

	transactionsDel := transactionsDelivery.NewTransactionDelivery(&transactionsDelivery.DeliveryConfig{
		TransactionServiceInterface: transactionsSvc,
		Config:                      s.cfg,
		Logger:                      s.logger,
	})

	dashboardDel := dashboardDelivery.NewDashboardDelivery(&dashboardDelivery.DeliveryConfig{
		DashboardServiceInterface: dashboardSvc,
		Config:                    s.cfg,
		Logger:                    s.logger,
	})

	inventoryDel := inventoryDelivery.NewInventoryDelivery(&inventoryDelivery.DeliveryConfig{
		InventoryServiceInterface: inventorySvc,
		Config:                    s.cfg,
		Logger:                    s.logger,
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
		customersGroupV1 := api.Group("/v1")
		{
			customersDelivery.MapCustomerRoutes(customersGroupV1, customersDel, middlewareManager)
		}
		servicesGroupV1 := api.Group("/v1")
		{
			servicesDelivery.MapServiceRoutes(servicesGroupV1, servicesDel, middlewareManager)
		}
		transactionsGroupV1 := api.Group("/v1")
		{
			transactionsDelivery.MapTransactionRoutes(transactionsGroupV1, transactionsDel, middlewareManager)
		}
		dashboardGroupV1 := api.Group("/v1")
		{
			dashboardDelivery.MapDashboardRoutes(dashboardGroupV1, dashboardDel, middlewareManager)
		}
		inventoryGroupV1 := api.Group("/v1")
		{
			inventoryDelivery.MapInventoryRoutes(inventoryGroupV1, inventoryDel, middlewareManager)
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
