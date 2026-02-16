package http

import (
	"net/http"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/services"
	"github.com/RianIhsan/pos-laundry-be/internal/features/services/dto"
	"github.com/RianIhsan/pos-laundry-be/pkg/httpErrors/response"
	"github.com/RianIhsan/pos-laundry-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type serviceDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service services.ServiceServiceInterface
}

func NewServiceDelivery(cfg *DeliveryConfig) services.ServiceDeliveryInterface {
	return &serviceDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.ServiceServiceInterface,
	}
}

func (d *serviceDelivery) AddService() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(dto.CreateServiceRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		created, err := d.service.AddService(c, *req)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(c, http.StatusCreated, "Service created", created)
	}
}

func (d *serviceDelivery) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		list, err := d.service.GetList(c)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", list)
	}
}

func (d *serviceDelivery) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)

		found, err := d.service.GetById(c, uint(idUint))
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", found)
	}
}

func (d *serviceDelivery) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		req := new(dto.UpdateServiceRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		if err := d.service.Update(c, uint(idUint), *req); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}

func (d *serviceDelivery) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)

		if err := d.service.Delete(c, uint(idUint)); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}
