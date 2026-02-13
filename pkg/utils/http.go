package utils

import (
	"github.com/RianIhsan/pos-laundry-be/pkg/contextutils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogErrorResponse(ctx *gin.Context, logger *logrus.Logger, err error) {
	logger.WithError(err).WithFields(logrus.Fields{
		"requestId": contextutils.GetRequestId(ctx),
		"IPAddress": contextutils.GetIPAddress(ctx),
	}).Error("Logger error response")
}
