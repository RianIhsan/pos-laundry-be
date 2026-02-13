package service

import (
	"github.com/RianIhsan/pos-laundry-be/config"
	"github.com/RianIhsan/pos-laundry-be/internal/features/users"

	"github.com/sirupsen/logrus"
)

type ServiceConfig struct {
	UserRepoInterface users.UserRepositoryInterface
	Logger            *logrus.Logger
	Config            *config.Config
}
