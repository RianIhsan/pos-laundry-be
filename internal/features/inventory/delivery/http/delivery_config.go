package http

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/inventory"
	"github.com/sirupsen/logrus"
)

type DeliveryConfig struct {
	Config                    *config.Config
	Logger                    *logrus.Logger
	InventoryServiceInterface inventory.InventoryServiceInterface
}
