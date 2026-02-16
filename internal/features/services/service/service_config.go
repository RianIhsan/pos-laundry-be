package service

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/services"
	"github.com/sirupsen/logrus"
)

type ServiceConfig struct {
	ServiceRepoInterface services.ServiceRepositoryInterface
	Logger               *logrus.Logger
	Config               *config.Config
}
