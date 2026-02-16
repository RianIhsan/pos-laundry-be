package http

import (
	"net/http"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/inventory"
	"github.com/RianIhsan/pos-laundry-be/internal/features/inventory/dto"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/RianIhsan/pos-laundry-be/pkg/httpErrors/response"
	"github.com/RianIhsan/pos-laundry-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type inventoryDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service inventory.InventoryServiceInterface
}

func NewInventoryDelivery(cfg *DeliveryConfig) inventory.InventoryDeliveryInterface {
	return &inventoryDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.InventoryServiceInterface,
	}
}

func (d *inventoryDelivery) GetAlerts() gin.HandlerFunc {
	return func(c *gin.Context) {
		alerts, err := d.service.GetAlerts(c)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", alerts)
	}
}

func (d *inventoryDelivery) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(dto.CreateInventoryRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		authInterface, exists := c.Get("auth")
		if !exists {
			utils.LogErrorResponse(c, d.logger, nil)
			response.SendErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
			return
		}
		auth, ok := authInterface.(*middleware.Auth)
		if !ok {
			utils.LogErrorResponse(c, d.logger, nil)
			response.SendErrorResponse(c, http.StatusUnauthorized, "Invalid auth data")
			return
		}
		userID := uint(auth.Id)

		created, err := d.service.Create(c, userID, *req)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusCreated, "created", created)
	}
}

func (d *inventoryDelivery) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		list, err := d.service.List(c)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", list)
	}
}

func (d *inventoryDelivery) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		found, err := d.service.GetByID(c, uint(idUint))
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", found)
	}
}

func (d *inventoryDelivery) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		req := new(dto.UpdateInventoryRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		authInterface, exists := c.Get("auth")
		if !exists {
			utils.LogErrorResponse(c, d.logger, nil)
			response.SendErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
			return
		}
		auth, ok := authInterface.(*middleware.Auth)
		if !ok {
			utils.LogErrorResponse(c, d.logger, nil)
			response.SendErrorResponse(c, http.StatusUnauthorized, "Invalid auth data")
			return
		}
		userID := uint(auth.Id)

		updated, err := d.service.Update(c, userID, uint(idUint), *req)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", updated)
	}
}

func (d *inventoryDelivery) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)

		authInterface, exists := c.Get("auth")
		if !exists {
			utils.LogErrorResponse(c, d.logger, nil)
			response.SendErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
			return
		}
		auth, ok := authInterface.(*middleware.Auth)
		if !ok {
			utils.LogErrorResponse(c, d.logger, nil)
			response.SendErrorResponse(c, http.StatusUnauthorized, "Invalid auth data")
			return
		}
		userID := uint(auth.Id)

		if err := d.service.Delete(c, userID, uint(idUint)); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}
