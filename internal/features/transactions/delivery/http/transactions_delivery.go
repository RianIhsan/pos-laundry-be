package http

import (
	"net/http"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/transactions"
	"github.com/RianIhsan/pos-laundry-be/internal/features/transactions/dto"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/RianIhsan/pos-laundry-be/pkg/httpErrors/response"
	"github.com/RianIhsan/pos-laundry-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type transactionDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service transactions.TransactionServiceInterface
}

func NewTransactionDelivery(cfg *DeliveryConfig) transactions.TransactionDeliveryInterface {
	return &transactionDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.TransactionServiceInterface,
	}
}

func (d *transactionDelivery) CreateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(dto.CreateTransactionRequest)
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

		// Convert uint64 to uint for userID
		userId := uint(auth.Id)

		created, err := d.service.CreateTransaction(c, *req, userId)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(c, http.StatusCreated, "Transaction created successfully", created)
	}
}

func (d *transactionDelivery) GetList() gin.HandlerFunc {
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

func (d *transactionDelivery) GetById() gin.HandlerFunc {
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

func (d *transactionDelivery) UpdateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)

		req := new(dto.UpdateTransactionStatusRequest)
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

		updated, err := d.service.UpdateStatus(c, uint(idUint), userID, *req)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(c, http.StatusOK, "Transaction status updated successfully", updated)
	}
}
