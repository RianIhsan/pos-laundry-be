package http

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/services"
	"github.com/sirupsen/logrus"
)

type DeliveryConfig struct {
	ServiceServiceInterface services.ServiceServiceInterface
	Config                  *config.Config
	Logger                  *logrus.Logger
}
