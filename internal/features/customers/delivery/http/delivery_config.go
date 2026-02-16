package http

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/customers"
	"github.com/sirupsen/logrus"
)

type DeliveryConfig struct {
	CustomerServiceInterface customers.CustomerServiceInterface
	Config                   *config.Config
	Logger                   *logrus.Logger
}
