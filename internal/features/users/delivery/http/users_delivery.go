package http

import (
	"net/http"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/users"
	"github.com/RianIhsan/pos-laundry-be/internal/features/users/dto"
	"github.com/RianIhsan/pos-laundry-be/internal/middleware"
	"github.com/RianIhsan/pos-laundry-be/pkg/httpErrors/response"
	"github.com/RianIhsan/pos-laundry-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type userDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service users.UserServiceInterface
}

func NewUserDelivery(cfg *DeliveryConfig) users.UserDeliveryInterface {
	return &userDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.UserServiceInterface,
	}
}

func (u *userDelivery) RegisterUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := new(dto.RegisterUserRequest)
		if err := context.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "Invalid request body")
			return
		}

		createdUser, err := u.service.AddUser(context, *request)
		if err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(context, http.StatusCreated, "User created successfully", createdUser)

	}
}

func (u *userDelivery) LoginUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := new(dto.LoginUserRequest)
		if err := context.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusBadRequest, "Invalid request body")
			return
		}

		token, err := u.service.LoginUser(context, request)
		if err != nil {
			utils.LogErrorResponse(context, u.logger, err)
			response.SendErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(context, http.StatusOK, "Login successful", token)
	}
}

func (u *userDelivery) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		authData, exists := c.Get("auth")
		if !exists {
			utils.LogErrorResponse(c, u.logger, errors.New("no auth found"))
			response.SendErrorResponse(c, http.StatusUnauthorized, "no auth found")
			return
		}

		auth := authData.(*middleware.Auth)
		getUser, err := u.service.GetById(c, auth.Id)
		if err != nil {
			utils.LogErrorResponse(c, u.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(c, http.StatusOK, "success", getUser)
	}
}

func (u *userDelivery) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		request := new(dto.UpdateUserRequest)
		if err := c.ShouldBindJSON(request); err != nil {
			utils.LogErrorResponse(c, u.logger, err)
			response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request body")
			return
		}

		err := u.service.Update(c, idUint, *request)
		if err != nil {
			utils.LogErrorResponse(c, u.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}

func (u *userDelivery) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idUint, _ := utils.ConvertStringToUint(id)
		err := u.service.Delete(c, idUint)
		if err != nil {
			utils.LogErrorResponse(c, u.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", nil)
	}
}
