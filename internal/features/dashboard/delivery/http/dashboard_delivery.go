package http

import (
	"net/http"
	"strconv"

	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/dashboard"
	"github.com/RianIhsan/pos-laundry-be/pkg/httpErrors/response"
	"github.com/RianIhsan/pos-laundry-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type dashboardDelivery struct {
	cfg     *config.Config
	logger  *logrus.Logger
	service dashboard.DashboardServiceInterface
}

func NewDashboardDelivery(cfg *DeliveryConfig) dashboard.DashboardDeliveryInterface {
	return &dashboardDelivery{
		cfg:     cfg.Config,
		logger:  cfg.Logger,
		service: cfg.DashboardServiceInterface,
	}
}

func (d *dashboardDelivery) GetStats() gin.HandlerFunc {
	return func(c *gin.Context) {
		stats, err := d.service.GetStats(c)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", stats)
	}
}

func (d *dashboardDelivery) GetActivities() gin.HandlerFunc {
	return func(c *gin.Context) {
		activities, err := d.service.GetActivities(c)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", activities)
	}
}

func (d *dashboardDelivery) GetActivityLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		if p := c.Query("page"); p != "" {
			if v, err := strconv.Atoi(p); err == nil {
				if v > 0 {
					page = v
				}
			}
		}

		res, err := d.service.GetActivityLogsPage(c, page)
		if err != nil {
			utils.LogErrorResponse(c, d.logger, err)
			response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.SendSuccesResponse(c, http.StatusOK, "success", res)
	}
}
