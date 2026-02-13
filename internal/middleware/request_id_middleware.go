package middleware

import (
	"github.com/RianIhsan/pos-laundry-be/pkg/contextutils"
	"github.com/gin-gonic/gin"
)

func (mw *MiddlewareManager) RequestIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.GetHeader("X-Request-Id")
		if requestId == "" {
			requestId = contextutils.AssignRequestId(ctx)
		}

		ctx.Writer.Header().Set("X-Request-Id", requestId)

		ctx.Next()
	}
}
