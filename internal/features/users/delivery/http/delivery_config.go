package http

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/users"
	"github.com/sirupsen/logrus"
)

type DeliveryConfig struct {
	UserServiceInterface users.UserServiceInterface
	Config               *config.Config
	Logger               *logrus.Logger
}
