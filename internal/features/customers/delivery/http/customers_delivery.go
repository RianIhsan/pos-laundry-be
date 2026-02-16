package http

import (
	"net/http"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/customers"
	"github.com/RianIhsan/pos-laundry-be/internal/features/customers/dto"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/RianIhsan/pos-laundry-be/pkg/httpErrors/response"
	"github.com/RianIhsan/pos-laundry-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type customerDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service customers.CustomerServiceInterface
}

func NewCustomerDelivery(cfg *DeliveryConfig) customers.CustomerDeliveryInterface {
	return &customerDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.CustomerServiceInterface,
	}
}

func (d *customerDelivery) AddCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(dto.CreateCustomerRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		// Get auth from middleware
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

		created, err := d.service.AddCustomer(c, userID, *req)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(c, http.StatusCreated, "Customer created", created)
	}
}

func (d *customerDelivery) GetList() gin.HandlerFunc {
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

func (d *customerDelivery) GetById() gin.HandlerFunc {
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

func (d *customerDelivery) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		req := new(dto.UpdateCustomerRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		// Get auth from middleware
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

		if err := d.service.Update(c, userID, uint(idUint), *req); err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}

func (d *customerDelivery) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)

		// Get auth from middleware
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
