package service

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/transactions"
	"github.com/RianIhsan/pos-laundry-be/pkg/activitylogger"
	"github.com/sirupsen/logrus"
)

type ServiceConfig struct {
	TransactionRepoInterface transactions.TransactionRepositoryInterface
	Logger                   *logrus.Logger
	Config                   *config.Config
	ActivityLogger           *activitylogger.ActivityLogger
}
