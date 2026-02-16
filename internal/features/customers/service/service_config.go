package service

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/customers"
	"github.com/RianIhsan/pos-laundry-be/pkg/activitylogger"
	"github.com/sirupsen/logrus"
)

type ServiceConfig struct {
	CustomerRepoInterface customers.CustomerRepositoryInterface
	Logger                *logrus.Logger
	Config                *config.Config
	ActivityLogger        *activitylogger.ActivityLogger
}
