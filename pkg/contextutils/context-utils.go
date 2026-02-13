package contextutils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	KeyRequestID = "requestId"
)

func GetRequestId(ctx *gin.Context) string {
	requestId := ctx.GetString(KeyRequestID)
	return requestId
}

func GetIPAddress(ctx *gin.Context) string {
	return ctx.ClientIP()
}

func AssignRequestId(ctx *gin.Context) string {
	id := uuid.New().String()

	ctx.Set(KeyRequestID, id)

	return id
}
