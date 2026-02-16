package service

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/dashboard"
	"github.com/sirupsen/logrus"
)

type ServiceConfig struct {
	DashboardRepoInterface dashboard.DashboardRepositoryInterface
	Logger                 *logrus.Logger
	Config                 *config.Config
}
